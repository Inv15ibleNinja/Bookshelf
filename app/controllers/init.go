package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB) // invoke InitDB function before
	revel.InterceptMethod((*GormController).SetDB, revel.BEFORE)
}
