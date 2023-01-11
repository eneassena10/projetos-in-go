package main

import (
	"net/http"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(echojwt.JWT([]byte("123456")))
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("123456"),
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
