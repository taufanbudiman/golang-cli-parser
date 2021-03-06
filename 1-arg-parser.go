package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("App", "Simple app")
	argAppName = app.Flag("name", "Application name").Required().String()
	argPort    = app.Flag("port", "Web server port").Short('p').Default("9000").Int()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// more code here ...
	appName := *argAppName
	port := fmt.Sprintf(":%d", *argPort)

	fmt.Printf("Starting %s at %s", appName, port)

	e := echo.New()
	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(port))
}
