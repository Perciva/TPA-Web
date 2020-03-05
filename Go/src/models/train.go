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
	Class         string  `gorm:"VARCHAR(100)"`
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
	}

	return train
}
func GetFilteredTrain(arrival string, dest string, date time.Time, classes []string, names []string)[]Train{
	db := Connection.Connect()
	defer db.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(dest)

	var res []Train
	db.Find(&res,"(class IN (?) OR ?) AND (name In (?) OR ?) AND arrival_Id = ? AND departure_Id = ? AND DATE_PART('day',arrival_time) = ?", classes, len(classes) == 0, names, len(names)==0,arrivalStation.Id, departureStation.Id, date.Day())

	for i, _ := range res {
		db.Model(&res[i]).Related(&res[i].Departure, "departure_Id")
		db.Model(&res[i]).Related(&res[i].Transit, "transit_Id")
		db.Model(&res[i]).Related(&res[i].Arrival, "arrival_Id")
	}

	return res
}

func GetTrainByLoc(arrival string, destination string, date time.Time) []Train {
	db := Connection.Connect()
	defer db.Close()

	arrivalStation := SearchTrainStationByName(arrival)
	departureStation := SearchTrainStationByName(destination)


	var train []Train
	db.Where("arrival_Id = ? AND departure_Id = ? AND DATE_PART('day',arrival_time) = ?",
		arrivalStation.Id, departureStation.Id, date.Day()).Find(&train)

	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_Id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_Id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_Id")
	}

	return train
}

func InsertTrain(name string, code string, class string, arrival string, arrivalTime time.Time, transit string, departure string, departureTime time.Time, seat int, price int) *Train {
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
		Class:         class,
		ArrivalId:     arrivalStation.Id,
		ArrivalTime:   arrivalTime.Add(-7 * time.Hour),
		TransitId:     transitId,
		DepartureId:   departureStation.Id,
		DepartureTime: departureTime.Add(-7 * time.Hour),
		Seat:          seat,
		Price:         price,
	}
	db.Save(newTrain)

	db.Model(newTrain).Related(&newTrain.Departure, "departure_Id")
	db.Model(newTrain).Related(&newTrain.Transit, "transit_Id")
	db.Model(newTrain).Related(&newTrain.Arrival, "arrival_Id")

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

	return newTrain
}

func DeleteTrain(Id int) *Train {
	db := Connection.Connect()
	defer db.Close()

	var train Train
	train = GetTrainById(Id)

	if train.Id == 0 {
		log.Println("Delete User Failed")
		return &train
	}

	err := db.Delete(train).Error

	if err != nil {
		panic("Error Delete Train !" + err.Error())
	}
	return &train
}
func GetTrainName()[]Train{
	db := Connection.Connect()
	defer db.Close()

	var train []Train
	db.Group("Name").Find(&train)
	db.Raw("Select distinct Name from trains").Scan(&train)
	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_Id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_Id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_Id")
		db.Model(train[i]).Related(&train[i].Class, "trainId")
	}
	fmt.Println("train names", train)

	return train

}
func GetTrainClass()[]Train{
	db := Connection.Connect()
	defer db.Close()

	var train []Train
	db.Raw("Select distinct Class from trains").Scan(&train)
	for i, _ := range train {
		db.Model(train[i]).Related(&train[i].Departure, "departure_Id")
		db.Model(train[i]).Related(&train[i].Transit, "transit_Id")
		db.Model(train[i]).Related(&train[i].Arrival, "arrival_Id")
		db.Model(train[i]).Related(&train[i].Class, "trainId")
	}
	fmt.Println("train class", train)

	return train

}