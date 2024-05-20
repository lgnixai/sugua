package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func HandleBlueprint(c echo.Context, app *pocketbase.PocketBase, blueprintId string) error {

	blueprint, err := app.Dao().FindRecordById("blueprints", blueprintId)
	if err != nil {
		return err
	}

	userRecord, _ := c.Get(apis.ContextAuthRecordKey).(*models.Record)
	if userRecord == nil {
		return c.JSON(401, "Unauthorized")
	}

	return c.JSON(200, blueprint)
}
