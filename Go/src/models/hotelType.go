package models

import (
	"Connection"
	"fmt"
	"time"
)

type HotelType struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"VARCHAR(MAX)"`
	Path      string `gorm:"VARCHAR(MAX)"`
	HotelId   int    `gorm:"Integer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql : "index"`
}
func init(){
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&HotelType{})

}
func GetAllHotelType() []HotelType {
	db := Connection.Connect()
	defer db.Close()

	var res []HotelType
	db.Find(&res)
	return res
}

func InsertHotelType(HotelId int, name string, path string)*HotelType {
	db := Connection.Connect()
	defer db.Close()

	newHotel := &HotelType{
		Name:    name,
		Path:    path,
		HotelId: HotelId,
	}

	db.Save(newHotel)
	fmt.Println(newHotel, "inserted")

	return newHotel
}
