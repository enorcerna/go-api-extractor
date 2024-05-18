package controllers

import (
	"net/http"

	"github.com/enorcerna/go-api-invest/src/services"
	"github.com/labstack/echo/v4"
)

func GetHtml(c echo.Context) error {
	link_web := c.QueryParam("url")
	return c.HTML(http.StatusOK, services.ExtractHtml(link_web))
}
