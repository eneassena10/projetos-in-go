package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	words = []string{"kind", "warm", "cup", "coin", "blue"}
	users = []User{
		{
			Name:  "Carlos",
			Email: "carlos@hotmail.com",
		},
		{
			Name:  "Lucas Santos",
			Email: "lucas@hmail.com",
		},
	}
)

func getWords(e echo.Context) error {
	return e.JSON(http.StatusOK, words)
}

func getUsers(e echo.Context) error {
	return e.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))
	const path = "/api"

	e.GET(path+"/words", getWords)
	e.GET(path+"/users", getUsers)

	e.Logger.Fatal(e.Start(":8080"))

}
