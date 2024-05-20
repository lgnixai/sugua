package orm

import (
	`log`

	`github.com/pocketbase/dbx`
	`github.com/pocketbase/pocketbase/daos`
	"github.com/pocketbase/pocketbase/models"
)

type Feeds struct {
	models.BaseModel
	FolderId    *int64 `db:"folder_id" json:"folder_id"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	FeedLink    string ` db:"feed_link" json:"feed_link"`
	Icon        string ` db:"icon" json:"icon,omitempty"`
	HasIcon     bool   ` db:"has_icon" json:"has_icon"`
}

var _ models.Model = (*Feeds)(nil)

func (m *Feeds) TableName() string {
	return "feeds" // the name of your collection
}
func FeedsQuery(dao *daos.Dao) *dbx.SelectQuery {
	return dao.ModelQuery(&Feeds{})
}

func GetFeeds(dao *daos.Dao) ([]*Feeds, error) {

	var c []*Feeds
	err := FeedsQuery(dao).OrderBy("score asc").All(&c)
	return c, err
}
func ListFeedsMissingIcons(dao *daos.Dao) []*Feeds {

	//query := dao.RecordQuery("feeds")
	//
	//records := []*Feeds{}
	//if err := query.All(&records); err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//spew.Dump(records)
	//return records
	var c []*Feeds
	err := FeedsQuery(dao).OrderBy("score asc").All(&c)
	if err != nil {
		log.Print(err)
	}
	return c
}
func UpdateFeedIcon(dao *daos.Dao, feedId string, icon string) bool {

	_, err := dao.DB().NewQuery("update feeds set icon = {:icon} where id = {:feed_id} ").
		Bind(dbx.Params{
			"feed_id": feedId,
			"icon":    icon,
		}).Execute()

	if err != nil {
		log.Print(err)
	}
	return true
	//_, err := s.db.Exec(`update feeds set icon = ? where id = ?`, icon, feedId)
	//return err == nil
}

func ResetFeedErrors(dao *daos.Dao) {
	//if _, err := s.db.Exec(`delete from feed_errors`); err != nil {
	//	log.Print(err)
	//}
}

func SetFeedSize(dao *daos.Dao, feedId string, size int) {

	_, err := dao.DB().NewQuery("insert into feed_size (id, size) values ({:id}, {:size}) on conflict (id) do update set size = excluded.size").
		Bind(dbx.Params{
			"id":   feedId,
			"size": size,
		}).Execute()

	if err != nil {
		log.Print(err)
	}
	//_, err := s.db.Exec(`
	//	insert into feed_sizes (feed_id, size)
	//	values (?, ?)
	//	on conflict (feed_id) do update set size = excluded.size`,
	//	feedId, size,
	//)
	//if err != nil {
	//	log.Print(err)
	//}
}
func SyncSearch(dao *daos.Dao) {

}
func SetFeedError(dao *daos.Dao, feedID string, lastError error) {
	_, err := dao.DB().NewQuery("insert into feed_errors (id, error)values ({:id}, {:error}) " +
		"on conflict (id) do update set error = excluded.error ").
		Bind(dbx.Params{
			"id":    feedID,
			"error": lastError.Error(),
		}).Execute()

	if err != nil {
		log.Print(err)
	}
}
