package models

import (
	"Connection"
	"fmt"
	"log"
	"time"
)

type Train struct {
	Id            int          `gorm:"PRIMARY_KEY"`
	Name          string       `gorm:"VARCHAR(100); NOT NULL"`
	Code          string       `gorm:"VARCHAR(100); NOT NULL"`
	Class         TrainClass  `gorm:"FOREIGNKEY:TrainId"`
	Arrival       TrainStation `gorm:"FOREIGNKEY:ArrivalId"`
	ArrivalId     int          `gorm:"INTEGER; NOT NULL"`
	ArrivalTime   time.Time
	Transit       TrainStation `gorm:"FOREIGNKEY:TransitId"`
	TransitId     int          `gorm:"INTEGER; NOT NULL"`
	Departure     TrainStation `gorm:"FOREIGNKEY:DepartmentId"`
	DepartureId   int          `gorm:"INTEGER; NOT NULL"`
	DepartureTime time.Time
	Seat          int `gorm:"INTEGER; NOT NULL"`
	Price         int `gorm:"INTEGER; NOT NULL"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:index`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&Train{}).
		AddForeignKey("departure_Id", "train_stations(Id)", "CASCADE", "CASCADE").
		AddForeignKey("transit_Id", "train_stations(Id)", "CASCADE", "CASCADE").
		AddForeignKey("arrival_Id", "train_stations(Id)", "CASCADE", "CASCADE")

	log.Println("Initialize Train Success")
}

func GetAllTrain() []Train {
	db := Connection.Connect()
	defer db.Close()

	var train []Train
	db.Find(&train)

	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_Id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_Id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_Id")
		db.Model(train[i]).Related(&train[i].Class, "trainId")
	}

	return train
}

func InsertTrain(name string, code string, arrival string, arrivalTime time.Time, transit string, departure string, departureTime time.Time, seat int, price int) *Train {
	db := Connection.Connect()
	defer db.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(departure)

	var transitStation TrainStation
	var transitId = 1
	if len(transit) != 0 {
		transitStation = SearchTrainStationByName(transit)
		transitId = transitStation.Id
	}

	newTrain := &Train{
		Name:          name,
		Code:          code,
		ArrivalId:     arrivalStation.Id,
		ArrivalTime:   arrivalTime,
		TransitId:     transitId,
		DepartureId:   departureStation.Id,
		DepartureTime: departureTime,
		Seat:          seat,
		Price:         price,
	}
	db.Save(newTrain)

	db.Model(newTrain).Related(&newTrain.Departure, "departure_Id")
	db.Model(newTrain).Related(&newTrain.Transit, "transit_Id")
	db.Model(newTrain).Related(&newTrain.Arrival, "arrival_Id")
	db.Model(newTrain).Related(&newTrain.Class, "train_Id")

	log.Println(newTrain,"Insert New Train Success")
	return newTrain
}

func GetTrainById(Id int) Train {
	db := Connection.Connect()
	defer db.Close()

	var train Train
	db.
		Where("Id = ?", Id).
		First(&train)

	db.Model(train).Related(&train.Departure, "departure_Id")
	db.Model(train).Related(&train.Transit, "transit_Id")
	db.Model(train).Related(&train.Arrival, "arrival_Id")
	db.Model(train).Related(&train.Class, "train_Id")

	return train
}
func GetTrainByLoc(arrival string, destination string, date time.Time) []Train {
	db := Connection.Connect()
	defer db.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(destination)

	fmt.Println(arrival,destination,date.Day())

	var train []Train
	db.Where("arrival_id = ? AND departure_id = ? AND DATE_PART('day',arrival_time) = ?",
		arrivalStation.Id, departureStation.Id, date.Day()).Find(&train)

	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_id")
	}

	return train
}
func GetTrainName()[]Train{
	db := Connection.Connect()
	defer db.Close()

	var train []Train
	db.Select("Name").Group("Name").Find(&train)
	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_Id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_Id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_Id")
		db.Model(train[i]).Related(&train[i].Class, "trainId")
	}
	fmt.Println("train names", train)

	return train

}

func UpdateTrain(Id int, arrivalTime time.Time, departureTime time.Time, seat int, price int) Train {
	db := Connection.Connect()
	defer db.Close()

	var train Train
	db.
		Model(&train).
		Where("Id = ?", Id).
		Updates(map[string]interface{}{
			"arrival_time":   arrivalTime.Add(-7 * time.Hour),
			"departure_time": departureTime.Add(-7 * time.Hour),
			"seat":           seat,
			"price":          price,
		})
	var newTrain Train
	db.
		Where("Id = ?", Id).
		Find(&newTrain)

	db.Model(newTrain).Related(&newTrain.Departure, "departure_Id")
	db.Model(newTrain).Related(&newTrain.Transit, "transit_Id")
	db.Model(newTrain).Related(&newTrain.Arrival, "arrival_Id")
	db.Model(newTrain).Related(&newTrain.Class, "train_Id")

	log.Println(newTrain,"Update Train Success")
	return newTrain
}

func DeleteTrain(Id int) *Train {
	db := Connection.Connect()
	defer db.Close()

	var train Train
	train = GetTrainById(Id)
	fmt.Println("Deleteing -> Resolver",train)
	if train.Id == 0 {
		log.Println("Delete User Failed")
		return &train
	}

	err := db.Delete(train).Error

	if err != nil {
		panic("Error Delete Train !" + err.Error())
	}

	log.Println("Delete User Success")
	return &train
}
