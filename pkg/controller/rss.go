package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
)

func Hello(c echo.Context, app *pocketbase.PocketBase) error {

	return c.JSON(200, "blueprint")
}

//func (s *Server) FeedList(c *gin.Context) {
//	result, err := s.db.Db.Query("select * from feed", nil)
//	if err != nil {
//		panic(err)
//	}
//	rsMap := result.([]interface{})[0].(map[string]interface{})
//	rsData := rsMap["result"].([]interface{})
//
//	c.JSON(200, rsData)
//}

//
//func (s *Server) FeedAdd(ctx *gin.Context) {
//
//	request := FeedCreateForm{}
//	err := ctx.BindJSON(&request)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
//		return
//	}
//
//	result, err := worker.DiscoverFeed(request.Url)
//	switch {
//	case err != nil:
//		log.Printf("Faild to discover feed for %s: %s", request.Url, err)
//
//		ctx.JSON(http.StatusOK, gin.H{"status": "notfound"})
//
//	case len(result.Sources) > 0:
//		ctx.JSON(http.StatusOK, map[string]interface{}{"status": "multiple", "choice": result.Sources})
//	case result.Feed != nil:
//
//		feed, err := s.db.Db.Create("feed", &model.Feed{
//			Title:       result.Feed.Title,
//			Description: "",
//			Link:        result.Feed.SiteURL,
//			FeedLink:    result.FeedLink,
//			FolderId:    request.FolderID,
//		})
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
//			return
//		}
//
//		feedMap := feed.([]interface{})[0].(map[string]interface{})
//		//把feedMap转化为model.Feed
//		//ff := model.Feed{
//		//	Id:       feedMap["id"].(string),
//		//	Title:    feedMap["title"].(string),
//		//	FeedLink: feedMap["feed_link"].(string),
//		//	Link:     feedMap["link"].(string),
//		//}
//		id := feedMap["id"].(string)
//
//		items := worker.ConvertItems(result.Feed.Items, id)
//		fmt.Println(len(items))
//		if len(items) > 0 {
//			feedMap["Size"] = int64(len(items))
//			for _, item := range items {
//				s.db.Db.Create("feed_item", item)
//			}
//
//			s.db.Db.Update(id, feedMap)
//			//s.db.SyncSearch()
//		}
//		//s.worker.FindFeedFavicon(*feed)
//
//		ctx.JSON(http.StatusOK, map[string]interface{}{
//			"status": "success",
//			"feed":   feed,
//		})
//	default:
//		ctx.JSON(http.StatusOK, map[string]string{"status": "notfound"})
//	}
//
//}
//
////频道列表
//func (s *Server) FeedItemList(c *gin.Context) {
//	result, err := s.db.Db.Query("select * from feed_item", nil)
//	if err != nil {
//		panic(err)
//	}
//	rsMap := result.([]interface{})[0].(map[string]interface{})
//	rsData := rsMap["result"].([]interface{})
//
//	c.JSON(200, rsData)
//}

//func (s *Server) OPMLImport(ctx *gin.Context) {
//	file, err := ctx.FormFile("File")
//	if err != nil {
//		log.Print(err)
//		return
//	}
//	//如何把file 转成  io.Reader
//	fileReader, _ := file.Open()
//
//	doc, err := opml.Parse(fileReader)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
//		return
//	}
//
//	for _, f := range doc.Feeds {
//		s.db.Db.Create("feed", &model.Feed{
//			Title:       f.Title,
//			Description: "",
//			Link:        f.SiteUrl,
//			FeedLink:    f.FeedUrl,
//			Icon:        "",
//		})
//
//	}
//	//for _, f := range doc.Folders {
//	//	folder := s.db.CreateFolder(f.Title)
//	//	for _, ff := range f.AllFeeds() {
//	//		s.db.CreateFeed(ff.Title, "", ff.SiteUrl, ff.FeedUrl, &folder.Id)
//	//	}
//	//}
//
//	s.worker.FindFavicons()
//	s.worker.RefreshFeeds()
//
//	ctx.JSON(http.StatusOK, map[string]interface{}{
//		"status": "success",
//	})
//
//}
//func HandleBlueprintAdd(c echo.Context, app *pocketbase.PocketBase, blueprintId string) error {
//	blueprint, err := app.Dao().FindRecordById("blueprints", blueprintId)
//	if err != nil {
//		return err
//	}
//
//	userRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
//	if userRecord == nil {
//		return c.JSON(401, "Unauthorized")
//	}
//
//	// Add the user to the blueprint under the "users" field
//	currentUsers := blueprint.GetStringSlice("users")
//
//	// Check if the user is already "owner" of the blueprint
//	if blueprint.GetString("owner") == userRecord.Id {
//		return c.JSON(200, blueprint)
//	}
//
//	// Check if the user is already in the list
//	for _, u := range currentUsers {
//		if u == userRecord.Id {
//			return c.JSON(200, blueprint)
//		}
//	}
//
//	currentUsers = append(currentUsers, userRecord.Id)
//	blueprint.Set("users", currentUsers)
//
//	if err := app.Dao().SaveRecord(blueprint); err != nil {
//		return err
//	}
//
//	return c.JSON(200, blueprint)
//}
