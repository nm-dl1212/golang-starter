package main

import (
	"app/handlers"
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"

var e = createMux()

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/css", "src/css")
	return e
}

func main() {
	e.GET("/", articleIndex)
	e.GET("/tasks", handlers.GetTasks)

	e.Logger.Fatal(e.Start(":8080"))
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}
	return render(c, "article/index.html", data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	b, err := pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
