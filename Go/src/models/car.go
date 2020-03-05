package models

import (
	"Connection"
	"log"
	"time"
)

type Car struct {
	Id         int           `gorm:"PRIMARY_KEY"`
	CarModel   CarModel      `gorm:"FOREIGNKEY:CarModelID"`
	CarModelID int           `gorm:"INTEGER; NOT NULL"`
	Location   Location `gorm:"FOREIGNKEY:LocationID"`
	LocationID int           `gorm:"INTEGER; NOT NULL"`
	Price      int           `gorm:"INTEGER; NOT NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&Car{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE").
		AddForeignKey("car_model_id", "car_models(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Car Success")
}

func DropTableCar() {
	db := Connection.Connect()
	defer db.Close()

	db.DropTableIfExists(&Car{})
	db.AutoMigrate(&Car{})

	log.Println("Drop db Success")
}

func GetAllCar() []Car {
	db := Connection.Connect()
	defer db.Close()

	var cars []Car
	if ValidateKey() == false {
		return cars
	}
	db.Find(&cars)

	for i, _ := range cars {
		db.Model(&cars[i]).Related(&cars[i].Location, "location_id")
		db.Model(&cars[i]).Related(&cars[i].CarModel, "car_model_id")
	}

	return cars
}

func GetCarByLocation(city string) []Car {
	db := Connection.Connect()
	defer db.Close()

	location := GetLocByProvince(city)

	var cars []Car
	if ValidateKey() == false {
		return cars
	}

	if len(location) == 1 {
		db.Where("location_id = ?", location[0].Id).Find(&cars)
	} else {
		var listLocationID []int
		for i := 0; i < len(location); i++ {
			listLocationID = append(listLocationID, location[i].Id)
		}
		db.Where("location_id IN (?)", listLocationID).Find(&cars)
	}

	for i, _ := range cars {
		db.Model(&cars[i]).Related(&cars[i].Location, "location_id")
		db.Model(&cars[i]).Related(&cars[i].CarModel, "car_model_id")
	}

	return cars
}

func InsertCar(model string, city string, price int) *Car {
	db := Connection.Connect()
	defer db.Close()

	location := GetLocByCity(city)
	carModel := GetCarModelByModel(model)

	newCar := &Car{
		Price:      price,
		CarModelID: carModel.Id,
		LocationID: location.Id,
	}
	db.Save(newCar)

	log.Println("Insert New Car Success")
	return newCar
}
