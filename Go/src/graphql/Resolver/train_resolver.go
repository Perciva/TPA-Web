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
	arrival := p.Args["arrival"].(string)
	arrivalTime := p.Args["arrivalTime"].(string)
	transit := p.Args["transit"].(string)
	departure := p.Args["departure"].(string)
	departureTime := p.Args["departureTime"].(string)
	seat := p.Args["seat"].(int)
	price := p.Args["price"].(int)

	arrivalTimeConv, _ := time.Parse(time.RFC3339, arrivalTime)
	departureTimeConv, _ := time.Parse(time.RFC3339, departureTime)

	newTrain := models.InsertTrain(name, code, arrival, arrivalTimeConv, transit, departure, departureTimeConv, seat, price)

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
func GetTrainName(p graphql.ResolveParams) (i interface{}, err error){
	trainName:=models.GetTrainName()

	return trainName,nil
}

func DeleteTrain(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	fmt.Println("Resolver Delete Train")
	//idInt, _ := strconv.Atoi(id)

	train := models.DeleteTrain(id)

	return train, nil
}
