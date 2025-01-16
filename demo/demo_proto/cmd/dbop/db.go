package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/lgxyc/gomall/demo/demo_proto/biz/dal"
	"github.com/lgxyc/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/lgxyc/gomall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// test mysql crud
	email1 := "124413@qq.com"
	// single create
	mysql.DB.Create(&model.User{Email: email1, Password: "28998"})
	// update
	mysql.DB.Model(&model.User{}).Where("email = ?", email1).Update("password", "555")
	// read
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email = ?", email1).First(&row)
	fmt.Printf("%v", row)
	// soft delete
	mysql.DB.Where("email = ? ", email1).Delete(&model.User{})
	// acutal delete
	mysql.DB.Unscoped().Where("email = ? ", email1).Delete(&model.User{})
}
