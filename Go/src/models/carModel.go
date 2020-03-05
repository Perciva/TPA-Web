package models

import (
	"Connection"
	"log"
	"time"
)

type CarModel struct {
	Id        int    `gorm:"PRIMARY_KEY"`
	Brand     string `gorm:"VARCHAR(100); NOT NULL"`
	Model     string `gorm:"VARCHAR(100); NOT NULL"`
	Passenger int    `gorm:"INTEGER; NOT NULL"`
	Baggage   int    `gorm:"INTEGER; NOT NULL"`
	Image 	  string `gorm:"VARCHAR(100); NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func DropTableCarModel() {
	db := Connection.Connect()
	defer db.Close()

	db.DropTableIfExists(&CarModel{})
	db.AutoMigrate(&CarModel{})

	log.Println("Drop db Success")
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&CarModel{})

	log.Println("Initialize Car Model Success")
}

func GetAllCarModel() []CarModel {
	db := Connection.Connect()
	defer db.Close()

	var cars []CarModel
	if ValidateKey() == false {
		return cars
	}
	db.Find(&cars)

	return cars
}

func GetCarModelByModel(model string) CarModel {
	db := Connection.Connect()
	defer db.Close()

	var cars CarModel
	if ValidateKey() == false {
		return cars
	}
	db.
		Where("model = ?", model).
		First(&cars)

	return cars
}

func InsertCarModel(brand string, model string, passenger int, baggage int, image string) *CarModel {
	db := Connection.Connect()
	defer db.Close()

	newCar := &CarModel{
		Brand:     brand,
		Model:     model,
		Passenger: passenger,
		Baggage:   baggage,
		Image:     image,
	}
	db.Save(newCar)

	log.Println("Insert New Car Success")
	return newCar
}


