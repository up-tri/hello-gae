package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", indexHandler)

	// e.HTTPErrorHandler = httpErrorHandler

	e.Logger.Fatal(e.Start(":1323"))
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

// func httpErrorHandler(err error, c echo.Context) {
// 	if header, ok := err.(*echo.HTTPError); ok {
// 	} else {
// 		if header.Code == 404 {
// 			c.Render(http.StatusNotFound, "404", nil)
// 		} else {
// 			c.Render(http.StatusInternalServerError, "500", nil)
// 		}
// 	}
// }
