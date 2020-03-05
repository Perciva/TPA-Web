package models

import (
	"Connection"
	"log"
	"time"
)

type HotelImage struct{
	Id int `gorm:"PRIMARY_KEY"`
	HotelId int `gorm:"INTEGER"`
	Path string `gorm:"VARCHAR(MAX)"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time	`sql : "index"`
}
func init(){
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&HotelImage{})
}

func GetAllHotelImages()[]HotelImage{
	db:= Connection.Connect()
	defer db.Close()

	var res []HotelImage
	if ValidateKey() == false {
		return res
	}
	db.Find(&res)

	return res
}
func InsertHotelImage(HotelId int, path string)*HotelImage{
	db:= Connection.Connect()
	defer db.Close()

	newImage:= &HotelImage{
		HotelId:   HotelId,
		Path:      path,
	}
	db.Save(newImage)
	log.Println(newImage,"Inserted")

	return newImage
}
