package orm

import (
	`fmt`
	`time`

	`github.com/pocketbase/pocketbase/daos`
	"github.com/pocketbase/pocketbase/models"
)

type ItemStatus int

const (
	UNREAD  ItemStatus = 0
	READ    ItemStatus = 1
	STARRED ItemStatus = 2
)

// WIP - this is a work in progress
type Item struct {
	models.BaseModel

	GUID     string     `db:"guid" json:"guid"`
	FeedId   string     `db:"feed_id" json:"feed_id"`
	Title    string     `db:"title" json:"title"`
	Link     string     `db:"link" json:"link"`
	Content  string     `db:"content" json:"content,omitempty"`
	Date     time.Time  `db:"date" json:"date"`
	Status   ItemStatus `db:"status" json:"status"`
	ImageURL *string    `db:"image_url" json:"image"`
	AudioURL *string    `db:"audio_url" json:"podcast_url"`

	//		Title    string     `json:"title"`
	//Link     string     `json:"link"`
	//Content  string     `json:"content,omitempty"`
	//Date     time.Time  `json:"date"`
	//Status   ItemStatus `json:"status"`
	//ImageURL *string    `json:"image"`
	//AudioURL *string    `json:"podcast_url"`

}
type ItemList []Item

func (list ItemList) Len() int {
	return len(list)
}

func (list ItemList) SortKey(i int) string {
	return list[i].Date.Format(time.RFC3339) + "::" + list[i].GUID
}

func (list ItemList) Less(i, j int) bool {
	return list.SortKey(i) < list.SortKey(j)
}

func (list ItemList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

var _ models.Model = (*Item)(nil)

func (m *Item) TableName() string {
	return "items" // the name of your collection
}
func CreateItems(dao *daos.Dao, items []Item) bool {
	collection, err := dao.FindCollectionByNameOrId("items")

	if err != nil {
		return false
	}
	err = dao.RunInTransaction(func(txDao *daos.Dao) error {
		now := time.Now().UTC()

		itemsSorted := ItemList(items)

		for _, f := range itemsSorted {

			record := models.NewRecord(collection)
			//spew.Dump(f)
			// set individual fields
			// or bulk load with record.Load(map[string]any{...})
			record.Set("guid", f.GUID)
			record.Set("feed_id", f.FeedId)

			record.Set("title", f.Title)

			record.Set("link", f.Link)
			record.Set("content", f.Content)
			record.Set("date", f.Date)
			record.Set("status", UNREAD)
			record.Set("image", f.ImageURL)
			record.Set("podcast_url", f.AudioURL)
			record.Set("date_arrived", now)

			err = txDao.SaveRecord(record)
			if err != nil {

				fmt.Println(err)

			}
			//fmt.Println(f.Title)
			//if err := dao.SaveRecord(record); err != nil {
			//	fmt.Println(err)
			//	return err
			//
			//}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// ArtistQuery returns a new dbx.SelectQuery for the Artist model.
//func ArtistQuery(dao *daos.Dao) *dbx.SelectQuery {
//	return dao.ModelQuery(&Artist{})
//}
//
//// GetArtists retrieves all art forms from the database and returns them as a slice of Artist pointers.
//// It takes a dao object as a parameter and returns the slice of Artist pointers and an error (if any).
//func GetArtists(dao *daos.Dao) ([]*Artist, error) {
//	var c []*Artist
//	err := ArtistQuery(dao).OrderBy("name asc").All(&c)
//	return c, err
//}
//
//// GetArtistBySlug retrieves an art form from the database by its slug.
//// It takes a dao object and a slug string as arguments and returns a pointer to the retrieved Artist object and an error (if any).
//func GetArtistBySlug(dao *daos.Dao, slug string) (*Artist, error) {
//	var c Artist
//	err := ArtistQuery(dao).AndWhere(dbx.NewExp("LOWER(slug)={:slug}", dbx.Params{
//		"slug": slug,
//	})).
//		Limit(1).
//		One(&c)
//	return &c, err
//}
//
//func GetArtistByNameLike(dao *daos.Dao, name string) ([]*Artist, error) {
//	var c []*Artist
//	err := ArtistQuery(dao).AndWhere(dbx.NewExp("LOWER(name) LIKE {:name}", dbx.Params{
//		"name": "%" + name + "%",
//	})).All(&c)
//	return c, err
//}
//
//func GetArtistById(dao *daos.Dao, id string) (*Artist, error) {
//	var c Artist
//	err := ArtistQuery(dao).AndWhere(dbx.NewExp("id={:id}", dbx.Params{
//		"id": id,
//	})).
//		Limit(1).
//		One(&c)
//	return &c, err
//}
