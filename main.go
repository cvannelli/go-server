package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	r := e.Group("/v1")

	r.GET("/", handler.HelloWorld)
	r.GET("/books", handler.GetBooks)
	// r.GET("/test", handler.CollectionSize)
	// r.GET("/contents", handler.TestContents)
	e.Logger.Fatal(e.Start(":1323"))
}