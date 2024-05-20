package controller

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"

	"github.com/lgnixai/sugua/pkg/k8s"
)

func HandleClusterInfo(c echo.Context, app *pocketbase.PocketBase) error {
	clusterInfo, err := k8s.GetClusterInfo()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, clusterInfo)
}
