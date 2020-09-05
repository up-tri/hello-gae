package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", indexHandler)
	// e.HTTPErrorHandler = httpErrorHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	if err := e.Start(":" + port); err != nil {
		e.Logger.Fatal(err)
	} else {
		e.Logger.Info("Listening on port %s", port)
	}
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
