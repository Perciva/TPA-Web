package models

import (
	"../../connection"
	"Connection"
	"log"
	"time"
)

type Flight struct {
	Id                int           `gorm:"PRIMARY_KEY"`
	Company           FlightCompany `gorm:"FOREIGNKEY:CompanyId"`
	Model             string        `gorm:"VARCHAR(100); NOT NULL;"`
	CompanyId         int           `gorm:"INTEGER; NOT NULL;"`
	FromAirport       FlightAirport `gorm:"FOREIGNKEY:FromAirportId"`
	FromAirportId     int           `gorm:"INTEGER"`
	FromLocationId    int           `gorm:"INTEGER"`
	ArrivalTime       time.Time
	ToAirport         FlightAirport `gorm:"FOREIGNKEY:ToAirportId"`
	ToAirportId       int           `gorm:"INTEGER"`
	ToLocationId      int           `gorm:"INTEGER"`
	DepartureTime     time.Time
	Price             int              `gorm:"INTEGER"`
	Facility          []FlightFacility `gorm:"FOREIGNKEY:FlightId"`
	Transit           FlightAirport    `gorm:"FOREIGNKEY:TransitId"`
	TransitId         int              `gorm:"INTEGER"`
	TransitLocationId int              `gorm:"INTEGER"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&Flight{}).
		AddForeignKey("company_Id", "flight_companies(Id)", "CASCADE", "CASCADE").
		AddForeignKey("from_airport_Id", "flight_airports(Id)", "CASCADE", "CASCADE").
		AddForeignKey("to_airport_Id", "flight_airports(Id)", "CASCADE", "CASCADE")

	log.Println("Initialize Flight Airport Success")
}

func GetAllFlight() []Flight {
	db := Connection.Connect()
	defer db.Close()

	var flight []Flight
	db.Find(&flight)

	for i, _ := range flight {
		db.Model(&flight[i]).Related(&flight[i].Company, "company_Id")
		db.Model(&flight[i]).Related(&flight[i].FromAirport, "from_airport_Id")
		db.Model(&flight[i]).Related(&flight[i].FromAirport.Location, "from_location_Id")
		db.Model(&flight[i]).Related(&flight[i].ToAirport, "to_airport_Id")
		db.Model(&flight[i]).Related(&flight[i].ToAirport.Location, "to_location_Id")
		db.Model(&flight[i]).Related(&flight[i].Facility, "flight_Id")
		db.Model(&flight[i]).Related(&flight[i].Transit, "transit_Id")
		db.Model(&flight[i]).Related(&flight[i].Transit.Location, "transit_location_Id")
	}

	return flight
}

func GetFlightById(Id int) Flight {
	db := Connection.Connect()
	defer db.Close()

	var flight Flight
	db.
		Where("Id = ?", Id).
		First(&flight)

	return flight
}

func InsertFlight(companyName string, fromAirportName string, arrivalTime time.Time,
	toAirportName string, departureTime time.Time, price int, model string, transitAirportName string) *Flight {

	db := Connection.Connect()
	defer db.Close()

	company := GetFlightCompanyByName(companyName)
	fromAirport := GetFlightAirportByCode(fromAirportName)
	toAirport := GetFlightAirportByCode(toAirportName)
	transitId := 1000
	transitLocationId := 1
	var transitAirport FlightAirport
	if len(transitAirportName) != 0 {
		transitAirport = GetFlightAirportByCode(transitAirportName)
		transitId = transitAirport.Id
		transitLocationId = transitAirport.LocationId
	}

	newFlight := &Flight{
		CompanyId:         company.Id,
		FromAirportId:     fromAirport.Id,
		ArrivalTime:       arrivalTime,
		ToAirportId:       toAirport.Id,
		DepartureTime:     departureTime,
		TransitId:         transitId,
		Price:             price,
		Model:             model,
		FromLocationId:    fromAirport.LocationId,
		ToLocationId:      toAirport.LocationId,
		TransitLocationId: transitLocationId,
	}
	db.Save(newFlight)

	log.Println("Insert New Flight Success")
	return newFlight
}

func UpdateFlight(Id int, fromAirportName string, arrivalTime time.Time,
	toAirportName string, departureTime time.Time, price int, model string, transitAirportName string) Flight {
	db := Connection.Connect()
	defer db.Close()

	fromAirport := GetFlightAirportByCode(fromAirportName)
	toAirport := GetFlightAirportByCode(toAirportName)
	transitId := 1000
	transitLocationId := 1
	var transitAirport FlightAirport
	if len(transitAirportName) != 0 {
		transitAirport = GetFlightAirportByCode(transitAirportName)
		transitId = transitAirport.Id
		transitLocationId = transitAirport.LocationId
	}

	var updateFlight Flight
	db.
		Model(&updateFlight).
		Where("Id = ?", Id).
		Updates(map[string]interface{}{
			"FromAirportId":     fromAirport.Id,
			"ArrivalTime":       arrivalTime,
			"ToAirportId":       toAirport.Id,
			"DepartureTime":     departureTime,
			"TransitId":         transitId,
			"Price":             price,
			"Model":             model,
			"FromLocationId":    fromAirport.LocationId,
			"ToLocationId":      toAirport.LocationId,
			"TransitLocationId": transitLocationId,
		})

	db.Model(&updateFlight).Related(&updateFlight.Company, "company_Id")
	db.Model(&updateFlight).Related(&updateFlight.FromAirport, "from_airport_Id")
	db.Model(&updateFlight).Related(&updateFlight.FromAirport.Location, "from_location_Id")
	db.Model(&updateFlight).Related(&updateFlight.ToAirport, "to_airport_Id")
	db.Model(&updateFlight).Related(&updateFlight.ToAirport.Location, "to_location_Id")
	db.Model(&updateFlight).Related(&updateFlight.Facility, "updateFlight_Id")
	db.Model(&updateFlight).Related(&updateFlight.Transit, "transit_Id")
	db.Model(&updateFlight).Related(&updateFlight.Transit.Location, "transit_location_Id")

	log.Println("Insert New Flight Success")
	return updateFlight
}

func DeleteFlight(Id int) *Flight {
	db := Connection.Connect()
	defer db.Close()

	var flight Flight
	flight = GetFlightById(Id)

	if flight.Id == 0 {
		log.Println("Delete Flight Failed")
		return &flight
	}

	err := db.Delete(flight).Error

	if err != nil {
		panic("Error Delete Hotel !" + err.Error())
	}

	log.Println("Delete Flight Success")
	return &flight
}
