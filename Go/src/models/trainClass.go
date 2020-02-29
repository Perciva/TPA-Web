package models

import (
	"Connection"
	"log"
	"time"
)

type TrainClass struct {
	id        int    `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(100); NOT NULL"`
	TrainId   int    `gorm:"INTEGER; NOT NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&TrainClass{})

	log.Println("Initialize Train Class Success")
}

func GetAllTrainClass() []TrainClass {
	db := Connection.Connect()
	defer db.Close()

	var trainClass []TrainClass
	db.Find(&trainClass)

	return trainClass
}

func SearchTrainClassByName(name string) TrainClass {
	db := Connection.Connect()
	defer db.Close()

	var trainClass TrainClass
	db.
		Where("name = ?", name).
		First(&trainClass)

	return trainClass
}

func InsertTrainClass(name string, id int) *TrainClass {
	db := Connection.Connect()
	defer db.Close()

	newTrainClass := &TrainClass{
		Name:      name,
		TrainId:   id,
	}
	db.Save(newTrainClass)

	log.Println("Insert New Train Class Success")
	return newTrainClass
}
