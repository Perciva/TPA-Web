package models

import (
	"Connection"
	"log"
	"time"
)

type FlightAirport struct {
	Id         int           `gorm:"PRIMARY_KEY"`
	Name       string        `gorm:"VARCHAR(100); NOT NULL"`
	Code       string        `gorm:"VARCHAR(5); NOT NULL"`
	Location   Location `gorm:"FOREIGNKEY:LocationId"`
	LocationId int           `gorm:"INTEGER"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&FlightAirport{}).
		AddForeignKey("location_Id", "locations(Id)", "CASCADE", "CASCADE")

	log.Println("Initialize Flight Airport Success")
}

func GetAllFlightAirport() []FlightAirport {
	db := Connection.Connect()
	defer db.Close()

	var airport []FlightAirport
	db.Find(&airport)

	for i, _ := range airport {
		db.Model(&airport[i]).Related(&airport[i].Location, "location_Id")
	}

	return airport
}

func InsertFlightAirport(name string, code string, city string) *FlightAirport {
	db := Connection.Connect()
	defer db.Close()

	location := GetLocByCity(city)

	newAirport := &FlightAirport{
		Name:       name,
		Code:       code,
		LocationId: location.Id,
	}
	db.Save(newAirport)

	log.Println("Insert New Flight Airport Success")
	return newAirport
}

func GetFlightAirportByCode(code string) FlightAirport {
	db := Connection.Connect()
	defer db.Close()

	var airport FlightAirport
	db.Where("code = ?", code).First(&airport)

	db.Model(&airport).Related(&airport.Location, "location_Id")

	return airport
}