package models

import (
	"Connection"
	"log"
	"time"
)

type FlightFacility struct {
	Id        int    `gorm:"PRIMARY_KEY"`
	FlightId  int    `gorm:"INTEGER; NOT NULL"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	Image     string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&FlightFacility{})

	log.Println("Initialize Flight Facility Success")
}

func GetAllFlightFacility() []FlightFacility {
	db := Connection.Connect()
	defer db.Close()

	var facility []FlightFacility
	if ValidateKey() == false {
		return facility
	}
	db.Find(&facility)

	return facility
}

func InsertFlightFacility(flightId int, name string) *FlightFacility {
	db := Connection.Connect()
	defer db.Close()

	imageLink := name + ".jpg"

	newFacility := &FlightFacility{
		FlightId:  flightId,
		Name:      name,
		Image:     imageLink,
	}
	db.Save(newFacility)

	log.Println("Insert New Flight Facility Success")
	return newFacility
}

