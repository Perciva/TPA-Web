package models

import (
	"Connection"
	"log"
	"time"
)

type TrainStation struct {
	Id         int      `gorm:"PRIMARY_KEY"`
	Name       string   `gorm:"VARCHAR(100); NOT NULL"`
	Code       string   `gorm:"VARCHAR(100); NOT NULL"`
	Location   Location `gorm:"FOREIGNKEY:LocationId"`
	LocationId int      `gorm:"INTEGER; NOT NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:index`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&TrainStation{}).
		AddForeignKey("location_Id", "locations(Id)", "CASCADE", "CASCADE")

	log.Println("Initialize Train Station Success")
}

func GetAllTrainStation() []TrainStation {
	db := Connection.Connect()
	defer db.Close()

	var trainStation []TrainStation
	if ValidateKey() == false {
		return trainStation
	}
	db.Find(&trainStation)

	for i, _ := range trainStation {
		db.Model(trainStation[i]).Related(&trainStation[i].Location, "location_Id")
	}

	return trainStation
}
func InTrainStation() {
	InsertTrainStation("Ancol", "AC", "Administrasi Jakarta Utara")
	InsertTrainStation("Angke", "AK", "Administrasi Jakarta Barat")
	InsertTrainStation("BNI City", "BNC", "Administrasi Jakarta Barat")
	InsertTrainStation("Bojong Indah", "BOI", "Administrasi Jakarta Barat")
	InsertTrainStation("Buaran", "BUA", "Administrasi Jakarta Barat")
	InsertTrainStation("Cakung", "CUK", "Administrasi Jakarta Barat")
	InsertTrainStation("Cawang", "CW", "Administrasi Jakarta Barat")
	InsertTrainStation("Cikini", "CKI", "Administrasi Jakarta Barat")
	InsertTrainStation("Cipinang", "CPN", "Administrasi Jakarta Barat")
	InsertTrainStation("Duren Kalibata", "DRN", "Administrasi Jakarta Barat")
	InsertTrainStation("Duri", "DU", "Administrasi Jakarta Barat")
	InsertTrainStation("Gambir", "GMR", "Administrasi Jakarta Barat")
	InsertTrainStation("Gang Sentiong", "GST", "Administrasi Jakarta Barat")
	InsertTrainStation("Gondangdia", "GDD", "Administrasi Jakarta Barat")
	InsertTrainStation("Grogol", "GGL", "Administrasi Jakarta Barat")
	InsertTrainStation("Jakarta Gudang", "JAKG", "Administrasi Jakarta Barat")
	InsertTrainStation("Jakarta Kota", "JAKK", "Administrasi Jakarta Barat")
	InsertTrainStation("Jatinegara", "JNG", "Administrasi Jakarta Barat")
	InsertTrainStation("Jayakarta", "JYK", "Administrasi Jakarta Barat")
	InsertTrainStation("Juanda", "JUA", "Administrasi Jakarta Barat")
	InsertTrainStation("Kalideres", "KDS", "Administrasi Jakarta Barat")
	InsertTrainStation("Kampung Bandan", "KPB", "Administrasi Jakarta Barat")
	InsertTrainStation("Karet", "KAT", "Administrasi Jakarta Barat")
	InsertTrainStation("Kebayoran", "KBY", "Administrasi Jakarta Barat")
	InsertTrainStation("Kemayoran", "KMO", "Administrasi Jakarta Barat")
	InsertTrainStation("Klender", "KLD", "Administrasi Jakarta Barat")
	InsertTrainStation("Klender Baru", "KLDB", "Administrasi Jakarta Barat")
	InsertTrainStation("Kramat", "KMT", "Administrasi Jakarta Barat")
	InsertTrainStation("Lenteng Agung", "LNA", "Administrasi Jakarta Barat")
	InsertTrainStation("Mampang", "MPG", "Administrasi Jakarta Barat")
	InsertTrainStation("Mangga Besar", "MGB", "Administrasi Jakarta Barat")
	InsertTrainStation("Manggarai", "MRI", "Administrasi Jakarta Barat")
	InsertTrainStation("Palmerah", "PLM", "Administrasi Jakarta Barat")
	InsertTrainStation("Pasar Minggu", "PSM", "Administrasi Jakarta Barat")
	InsertTrainStation("Pasar Minggu Baru", "PSMB", "Administrasi Jakarta Barat")
	InsertTrainStation("Pasar Senen", "PSE", "Administrasi Jakarta Barat")
	InsertTrainStation("Pasoso", "PSO", "Administrasi Jakarta Barat")
	InsertTrainStation("Pesing", "PSG", "Administrasi Jakarta Barat")
	InsertTrainStation("Pondok Jati", "POK", "Administrasi Jakarta Barat")
	InsertTrainStation("Rajawali", "RJW", "Administrasi Jakarta Barat")
	InsertTrainStation("Rawa Buaya", "RW", "Administrasi Jakarta Barat")
	InsertTrainStation("Sawah Besar", "SW", "Administrasi Jakarta Barat")
	InsertTrainStation("Sudirman", "SUD", "Administrasi Jakarta Barat")
	InsertTrainStation("Sungai Lagoa", "SAO", "Administrasi Jakarta Barat")
	InsertTrainStation("Taman Kota", "TKO", "Administrasi Jakarta Barat")
	InsertTrainStation("Tanah Abang", "THB", "Administrasi Jakarta Barat")
	InsertTrainStation("Tanjung Barat", "TNT", "Administrasi Jakarta Barat")
	InsertTrainStation("Tanjung Priok", "TPK", "Administrasi Jakarta Barat")
}
func InsertTrainStation(name string, code string, city string) *TrainStation {
	db := Connection.Connect()
	defer db.Close()

	location := GetLocByCity(city)

	newTrainStation := &TrainStation{
		Name:       name,
		Code:       code,
		LocationId: location.Id,
	}
	db.Save(newTrainStation)

	db.Model(newTrainStation).Related(&newTrainStation.Location, "location_Id")

	log.Println("Insert New Train Station Success")
	return newTrainStation
}

func SearchTrainStationByName(name string) TrainStation {
	db := Connection.Connect()
	defer db.Close()

	var trainStation TrainStation
	if ValidateKey() == false {
		return trainStation
	}
	db.
		Where("name = ?", name).
		First(&trainStation)

	db.Model(trainStation).Related(&trainStation.Location, "location_Id")

	return trainStation
}
