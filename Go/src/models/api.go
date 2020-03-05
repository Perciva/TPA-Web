package models

import (
	"Connection"
	"fmt"
	"middleware"
	"time"

)

type API struct{
	Id int `gorm:"PRIMARY_KEY"`
	Key string `gorm:"VARCHAR(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init(){
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&API{})
}

func ValidateKey()bool{
	db := Connection.Connect()
	defer db.Close()

	var Api API
	db.Where("key=?", middleware.Apikey).First(&Api)
	fmt.Println("apo",Api,middleware.Apikey)

	if Api.Id == 0 {
		return false
	}
	return true
}
