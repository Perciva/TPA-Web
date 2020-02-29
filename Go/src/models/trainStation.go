package models

import (
	"Connection"
	"log"
	"time"
)

type TrainStation struct {
	ID         int           `gorm:"PRIMARY_KEY"`
	Name       string        `gorm:"VARCHAR(100); NOT NULL"`
	Code       string        `gorm:"VARCHAR(100); NOT NULL"`
	Location   Location `gorm:"FOREIGNKEY:LocationID"`
	LocationID int           `gorm:"INTEGER; NOT NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:index`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&TrainStation{}).
		AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")

	log.Println("Initialize Train Station Success")
}

func GetAllTrainStation() []TrainStation {
	db := Connection.Connect()
	defer db.Close()

	var trainStation []TrainStation
	db.Find(&trainStation)

	for i, _ := range trainStation {
		db.Model(trainStation[i]).Related(&trainStation[i].Location, "location_id")
	}

	return trainStation
}

func InsertTrainStation(name string, code string, city string) *TrainStation {
	db := Connection.Connect()
	defer db.Close()

	location := GetLocByCity(city)

	newTrainStation := &TrainStation{
		Name:       name,
		Code:       code,
		LocationID: location.Id,
	}
	db.Save(newTrainStation)

	db.Model(newTrainStation).Related(&newTrainStation.Location, "location_id")

	log.Println("Insert New Train Station Success")
	return newTrainStation
}

func SearchTrainStationByName(name string) TrainStation {
	db := Connection.Connect()
	defer db.Close()

	var trainStation TrainStation
	db.
		Where("name = ?", name).
		First(&trainStation)

	db.Model(trainStation).Related(&trainStation.Location, "location_id")

	return trainStation
}
