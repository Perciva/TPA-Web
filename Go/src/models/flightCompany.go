package models

import (
	"Connection"
	"log"
	"time"
)

type FlightCompany struct {
	Id        int    `gorm:"PRIMARY_KEY"`
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
		AutoMigrate(&FlightCompany{})

	log.Println("Initialize Flight Company Success")
}

func GetAllFlightCompany() []FlightCompany {
	db := Connection.Connect()
	defer db.Close()

	var company []FlightCompany
	db.Find(&company)

	return company
}

func GetFlightCompanyByName(name string) FlightCompany {
	db := Connection.Connect()
	defer db.Close()

	var company FlightCompany
	db.Where("name = ?", name).First(&company)

	return company
}

func InsertFlightCompany(name string, image string) *FlightCompany {
	db := Connection.Connect()
	defer db.Close()

	newCompany := &FlightCompany{
		Name:  name,
		Image: image,
	}
	db.Save(newCompany)

	log.Println("Insert New Flight Company Success")
	return newCompany
}
