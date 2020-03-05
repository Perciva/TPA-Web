package Resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
	"time"
)

func AllTrain(p graphql.ResolveParams) (i interface{}, err error) {
	train := models.GetAllTrain()

	if len(train) == 0 {
		return nil, errors.New("error: no data found")
	}

	return train, nil
}


func InsertTrain(p graphql.ResolveParams) (i interface{}, err error) {
	name := p.Args["name"].(string)
	code := p.Args["code"].(string)
	class := p.Args["class"].(string)
	arrival := p.Args["arrival"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	transit := p.Args["transit"].(string)
	departure := p.Args["departure"].(string)
	departureTime := p.Args["departureTime"].(string)
	seat := p.Args["seat"].(int)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)

	newTrain := models.InsertTrain(name, code, class, arrival, arrivalTimeConv, transit, departure, departureTimeConv, seat, price)

	return newTrain, nil
}
func GetTrainByLoc(p graphql.ResolveParams) (i interface{}, err error) {
	arrival := p.Args["arrival"].(string)
	departure := p.Args["departure"].(string)
	date := p.Args["date"].(string)

	dateConv, _ := time.Parse(time.RFC3339, date)

	train := models.GetTrainByLoc(arrival, departure, dateConv)
	if len(train) == 0 {
		return nil, errors.New("error: no data found")
	}
	return train, nil
}

func UpdateTrain(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	arrivalTime := p.Args["arrivalTime"].(string)
	departureTime := p.Args["departureTime"].(string)
	seat := p.Args["seat"].(int)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)
	//idConv, _ := strconv.Atoi(id)

	newTrain := models.UpdateTrain(id, arrivalTimeConv, departureTimeConv, seat, price)

	fmt.Println("Updated Train", newTrain)
	return newTrain, nil
}
func FilterTrain(p graphql.ResolveParams) (i interface{}, err error){
	//arrival string, dest string, date time.Time, classes []string, names []string
	arrival := p.Args["arrival"].(string)
	dest:= p.Args["dept"].(string)
	date:=p.Args["date"].(string)
	fmt.Println("tes")
	classes := p.Args["classes"].([]interface{})
	names := p.Args["names"].([]interface{})

	dateConv, _ := time.Parse(time.RFC3339, date)

	classConv := make([]string, len(classes))
	for i := range classes {
		classConv[i] = classes[i].(string)
	}
	nameConv := make([]string, len(names))
	for i := range names {
		nameConv[i] = names[i].(string)
	}
	fmt.Println(nameConv, classConv, dateConv,"Filtered")
	trains := models.GetFilteredTrain(arrival,dest,dateConv,classConv,nameConv)
	fmt.Println(trains,"Filtered")
	return trains,nil
}
func GetTrainName(p graphql.ResolveParams) (i interface{}, err error){
	trainName:=models.GetTrainName()

	return trainName,nil
}
func GetTrainClass(p graphql.ResolveParams) (i interface{}, err error){
	trainName:=models.GetTrainClass()

	return trainName,nil
}

func DeleteTrain(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	fmt.Println("Resolver Delete Train")
	//idInt, _ := strconv.Atoi(id)

	train := models.DeleteTrain(id)

	return train, nil
}
