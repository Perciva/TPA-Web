package models

import (
	"Connection"
	"fmt"
	"time"
)
type HotelFacility struct {
	Id        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(MAX)"`
	HotelId   int    `gorm:"INTEGER"`
	Path      string `gorm:"VARCHAR(MAX)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql : "index"`
}
func init() {
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&HotelFacility{})
}
func GetAllHotelFacility() []HotelFacility {
	db := Connection.Connect()
	defer db.Close()

	var res []HotelFacility
	db.Find(&res)

	return res
}
func InsertHotelFacility(HotelId int, name string, path string) *HotelFacility {
	db := Connection.Connect()
	defer db.Close()

	newFac := &HotelFacility{
		Name:    name,
		HotelId: HotelId,
		Path:    path,
	}
	db.Save(newFac)
	fmt.Println(newFac, "Inserted Hotel Facility")

	return newFac
}

