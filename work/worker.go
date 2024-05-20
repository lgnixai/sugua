package worker

import (
	`fmt`
	`log`
	"sync"
	`sync/atomic`
	"time"

	`github.com/pocketbase/pocketbase`
	`github.com/pocketbase/pocketbase/daos`
	`github.com/spf13/afero`

	`github.com/lgnixai/sugua/orm`
)

const NUM_WORKERS = 4

type Worker struct {
	dao *daos.Dao

	pending *int32
	refresh *time.Ticker
	reflock sync.Mutex
	stopper chan bool
}

func NewWorker(app *pocketbase.PocketBase) *Worker {
	pending := int32(0)
	return &Worker{dao: app.Dao(), pending: &pending}
}
func (w *Worker) FeedsPending() int32 {
	return *w.pending
}

//func (w *Worker) StartFeedCleaner() {
//	go w.db.DeleteOldItems()
//	ticker := time.NewTicker(time.Hour * 24)
//	go func() {
//		for {
//			<-ticker.C
//			w.db.DeleteOldItems()
//		}
//	}()
//}

func (w *Worker) FindFavicons() {
	go func() {
		for _, feed := range orm.ListFeedsMissingIcons(w.dao) {
			w.FindFeedFavicon(feed)
		}
	}()
}

//
func (w *Worker) FindFeedFavicon(feed *orm.Feeds) {
	icon, err := findFavicon(feed.Link, feed.FeedLink)
	if err != nil {
		log.Printf("Failed to find favicon for %s (%s): %s", feed.FeedLink, feed.Link, err)
	}

	appfs := afero.NewOsFs()
	appfs.MkdirAll("assets/public/icon", 0755)
	afero.WriteFile(appfs, fmt.Sprintf("assets/public/icon/%s.ico", feed.Id), *icon, 0644)
	//err = ioutil.WriteFile("data/"+feed.Id+".icon", *icon, 0644)
	//fmt.Println("err", err, "data/"+feed.Id+".icon")

	if err != nil {
		log.Printf("Failed to find favicon for %s (%s): %s", feed.FeedLink, feed.Link, err)
	}
	feed.Icon = fmt.Sprintf("assets/public/icon/%s.ico", feed.Id)
	if icon != nil {
		orm.UpdateFeedIcon(w.dao, feed.Id, feed.Icon)
	}

}

func (w *Worker) RefreshFeeds() {
	w.reflock.Lock()
	defer w.reflock.Unlock()

	if *w.pending > 0 {
		log.Print("Refreshing already in progress")
		return
	}

	feeds, _ := orm.GetFeeds(w.dao)
	if len(feeds) == 0 {
		log.Print("Nothing to refresh")
		return
	}

	log.Print("Refreshing feeds")
	atomic.StoreInt32(w.pending, int32(len(feeds)))
	go w.refresher(feeds)
}

//
func (w *Worker) refresher(feeds []*orm.Feeds) {
	orm.ResetFeedErrors(w.dao)

	srcqueue := make(chan orm.Feeds, len(feeds))
	dstqueue := make(chan []orm.Item)

	for i := 0; i < NUM_WORKERS; i++ {
		go w.worker(srcqueue, dstqueue)
	}

	for _, feed := range feeds {
		srcqueue <- *feed
	}
	for i := 0; i < len(feeds); i++ {
		items := <-dstqueue
		if len(items) > 0 {
			orm.CreateItems(w.dao, items)
			orm.SetFeedSize(w.dao, items[0].FeedId, len(items))
		}
		atomic.AddInt32(w.pending, -1)
		orm.SyncSearch(w.dao)
	}
	close(srcqueue)
	close(dstqueue)

	log.Printf("Finished refreshing %d feeds", len(feeds))
}

//
func (w *Worker) worker(srcqueue <-chan orm.Feeds, dstqueue chan<- []orm.Item) {
	for feed := range srcqueue {
		items, err := listItems(w.dao, feed)
		if err != nil {
			orm.SetFeedError(w.dao, feed.Id, err)
		}
		dstqueue <- items
	}
}
