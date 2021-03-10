package controllers

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // my example for postgres

	// short name for revel

	r "github.com/revel/revel"
)

// type: revel controller with `*gorm.DB`
type GormController struct {
	*r.Controller
	DB *gorm.DB
}

// it can be used for jobs
var Gdb *gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	Gdb, err = gorm.Open("postgres", "user=postgres dbname=Book sslmode=disable password =postgres")
	Gdb.LogMode(true) // Print SQL statements
	if err != nil {
		//r.ERROR.Println("FATAL", err)
		log.Println(err)
		panic(err)
	}
}

func (c *GormController) SetDB() r.Result {
	c.DB = Gdb
	return nil
}
