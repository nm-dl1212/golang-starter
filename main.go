package main

import (
	"app/handlers"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/flosch/pongo2"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const tmplPath = "src/template/"

var db *pgx.Conn
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
	// PostgreSQLに接続
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432")
	if err != nil {
		e.Logger.Fatal("Unable to connect to database:", err)
	}
	defer db.Close(context.Background())

	// endpoint
	e.GET("/", articleIndex)
	e.GET("/tasks", handlers.GetTasks)

	e.Logger.Fatal(e.Start(":8080"))
}

func articleIndex(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Hello, World!",
		"Now":     time.Now(),
	}

	// PostgreSQL からのデータ取得
	users, err := getUsers()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	data["Users"] = users
	return render(c, "article/index.html", data)
}

func getUsers() ([]string, error) {
	rows, err := db.Query(context.Background(), "SELECT user_name FROM users")
	fmt.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var userName string
		if err := rows.Scan(&userName); err != nil {
			return nil, err
		}
		users = append(users, userName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func render(c echo.Context, file string, data map[string]interface{}) error {
	// Pongo2 テンプレートで使用するデータを適切な形式に変換
	tpl, err := pongo2.FromFile(tmplPath + file)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	b, err := tpl.ExecuteBytes(data)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.HTMLBlob(http.StatusOK, b)
}
