package orm

import (
	`fmt`
	"log"
	"time"

	`github.com/pocketbase/dbx`
	`github.com/pocketbase/pocketbase/daos`
)

type HTTPState struct {
	FeedID        int64
	LastRefreshed time.Time

	LastModified string
	Etag         string
}

func ListHTTPStates(dao *daos.Dao) map[int64]HTTPState {
	result := make(map[int64]HTTPState)
	query := dao.RecordQuery("http_states")
	records := []*HTTPState{}
	if err := query.All(&records); err != nil {
		return nil
	}
	for _, record := range records {
		result[record.FeedID] = *record
	}
	return result
}

func GetHTTPState(dao *daos.Dao, feedID string) *HTTPState {
	record, err := dao.FindRecordById("http_states", (feedID))
	if err != nil {
		log.Print(err)
		return nil
	}

	if record == nil {
		return nil
	}

	var state HTTPState
	state.FeedID = int64(record.GetInt("feed_id"))
	state.LastRefreshed = record.GetTime("last_refreshed")
	state.LastModified = record.GetString("last_modified")
	state.Etag = record.GetString("etag")

	return &state
}

func SetHTTPState(dao *daos.Dao, feedID string, lastModified, etag string) {
	res, err := dao.DB().NewQuery("insert into http_states (id, last_modified, etag, last_refreshed)values" +
		" ({:id}, {:lastModified}, {:etag}, datetime())on conflict (id) " +
		"do update set last_modified = {:lastModified}, etag = {:etag}, last_refreshed = datetime()",

	).Bind(dbx.Params{
		"id":           feedID,
		"lastModified": lastModified,
		"etag":         etag,
	}).Execute()

	fmt.Println(res)
	if err != nil {
		log.Print(err)
	}
}
