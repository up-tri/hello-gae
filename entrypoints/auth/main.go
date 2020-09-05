package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	myLogger "github.com/up-tri/hello-gae/infrastructure/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

var (
	panicStackSize int  = 4 << 10
	panicStackAll  bool = true
)

var (
	logger    *zap.Logger
	errLogger *zap.Logger
)

func main() {
	logger = myLogger.NewLogger()
	errLogger = myLogger.NewErrorLogger()

	// Catch Panic.
	defer func() {
		if x := recover(); x != nil {
			stack := make([]byte, panicStackSize)
			length := runtime.Stack(stack, panicStackAll)
			errLogger.Panic(fmt.Sprintf("%v", x), zap.ByteString("stack", stack[:length]))
		}
	}()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(myLogger.ZapLogger(logger))
	e.Use(myLogger.GCPRecoverWithConfig(errLogger,
		myLogger.RecoverConfig{
			StackSize:       panicStackSize,
			DisableStackAll: !panicStackAll,
		}))

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
	logger.Info("Hello World!")
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
