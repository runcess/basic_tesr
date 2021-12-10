package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		name := c.QueryParam("name")
		var helloMessage string = name + ", Hello!"
		return c.JSON(http.StatusOK, helloMessage)
	})
	e.Logger.Fatal(e.Start(":1331"))
}
