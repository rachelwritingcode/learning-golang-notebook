// main.go
package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Response struct {
	Message string `json:"mesage"`
}

func Hello(name string) string {
	return "Hello " + name
}

func AddCool(str string) string {
	return str + " is cool"
}

func GetHello(c echo.Context) error {
	r := &Response{
		Message: "Hello",
	}
	return c.JSON(http.StatusOK, r)
}

func main() {

	AddCool("Code")
	Hello("World")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/hello", GetHello)
	e.Logger.Fatal(e.Start(":8080"))

}
