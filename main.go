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
	r.GET("/test", handler.CollectionSize)
	r.GET("/contents", handler.TestContents)
	e.Logger.Fatal(e.Start(":1323"))
}

// func main() {

// 	client, err := mongo.NewClient("mongodb://foo:bar@localhost:27017")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// server := server{db = }
// 	// mux := httprouter.New()
// 	// mux.GET("/", index)
// 	// mux.GET("/about", about)
// 	// mux.GET("/contact", contact)
// 	// mux.GET("/apply", apply)
// 	// mux.POST("/apply", applyProcess)
// 	// mux.GET("/user/:name", user)
// 	// mux.GET("/blog/:category/:article", blogRead)
// 	// mux.POST("/blog/:category/:article", blogWrite)
// 	// http.ListenAndServe(":8080", mux)
// }

// func HandleError(w http.ResponseWriter, err error) {
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Fatalln(err)
// 	}
// }
