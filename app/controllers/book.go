package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"Bookshelf/app/models"

	"github.com/revel/revel"
)

// type ApiBook struct {
// 	ApiController
// 	GormController
// }

func BindJsonParams(i io.Reader, s interface{}) error {
	bytes, err := ioutil.ReadAll(i)
	if err != nil {
		return errors.New("can't read request data.")
	}

	if len(bytes) == 0 {
		return errors.New("data is nil")
	}

	return json.Unmarshal(bytes, s)
}

// func (c *GormController) Show1(id int64) revel.Result {
// 	var book models.Book
// 	//bookData, err := Dbm.Get(book, id)
// 	bookData := c.DB.First(&book, id)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	//return c.Response(&Response{OK, bookData})
// 	return c.Render(bookData)
// 	//return c.RenderJSON(bookData)
// }

func (c *GormController) Show1(id int64) revel.Result {
	//достаю из бд книжку с ID= id
	var book models.Book
	bookData := c.DB.First(&book, id)

	//в переменную ololo записываю данные книжки
	c.ViewArgs["data"] = bookData
	c.ViewArgs["book"] = book

	//рисую страницу, у которой есть параметр "ololo"
	return c.Render()
}

func (c *GormController) Show() revel.Result {

	//var Books interface{}
	var bk []models.Book
	_ = c.DB.Find(&bk)
	//Books = resp.Value
	// if resp != nil {

	// }
	c.ViewArgs["book"] = bk
	return c.Render()
}
