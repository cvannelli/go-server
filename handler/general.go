package handler

import (
	"net/http"

	"../controller"
	"github.com/labstack/echo"
)

// HelloWorld : Basic hello world function
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// CollectionSize : Returns collection size
func CollectionSize(c echo.Context) error {

	collection := "test"
	result, err := controller.GetCollectionSize(&collection)

	if err != nil {
		return err
	}
	return c.String(http.StatusOK, result)
}

// // TestContents : Returns contents of Test Collection
// func TestContents(c echo.Context) error {

// 	collection := "test"
// 	result, err := controller.GetTestContents(&collection)

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, result)
// }

// GetBooks : Returns contents of Book Collection
func GetBooks(c echo.Context) error {

	result, err := controller.Find("Books")

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
