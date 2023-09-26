package main

import (
	"test-server/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", api.HelloGolang)

	e.Logger.Fatal(e.Start(":8000"))
}
