# Hello Worldを表示する

```go:main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

```bash
go run main.go
```


# webサーバーを起動する

```go:main.go
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
```

```bash
# go modの初期化
go mod init app

# コードから必要パッケージを探索，インストールする
go mod tidy

# 実行
go run main.go
```