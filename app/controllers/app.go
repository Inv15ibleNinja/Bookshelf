package controllers

import (
	//"Bookshelf/app/controllers"
	"Bookshelf/app/models"

	"github.com/revel/revel"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
	// book := models.Book{Title: "Jinzhup"} // Note: In practice you should initialize all struct fields

	// if err := c.DB.Create(&book).Error; err != nil {
	// 	panic(err)
	// }
	var book models.Book
	ans := c.DB.First(&book, "id > 1")
	return c.RenderJSON(ans)
}
