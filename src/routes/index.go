package routes

import (
	"net/http"

	ctr "github.com/enorcerna/go-api-invest/src/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/api", func(c echo.Context) error {
		value := c.QueryParam("value")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"value": value,
		})
	})
	e.GET("/api/site", ctr.GetHtml)
	e.Logger.Fatal(e.Start(":1705"))
}
