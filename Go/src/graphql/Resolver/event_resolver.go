package Resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
	"time"
)

func AllEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	event := models.GetAllEntertainment()
	if len(event) == 0 {
		return nil, errors.New("error: no data found")
	}

	return event, nil
}
func GetEntertainmentByCategory(p graphql.ResolveParams) (i interface{}, err error){
	cat := p.Args["category"].(string)

	event:= models.GetEntertainmentByCategory(cat)

	return event,nil
}
func InsertEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	title := p.Args["title"].(string)
	price := p.Args["price"].(int)
	location := p.Args["location"].(string)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	date := p.Args["date"].(string)
	category := p.Args["category"].(string)
	image := p.Args["image"].(string)
	content := p.Args["content"].(string)

	dateConv, _ := time.Parse(time.RFC3339, date)
	fmt.Println(category)
	newEvent := models.InsertEntertainment(title, price, location, latitude, longitude, dateConv, category, image,content)

	return newEvent, nil
}

func UpdateEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	title := p.Args["title"].(string)
	price := p.Args["price"].(int)
	location := p.Args["location"].(string)
	latitude := p.Args["latitude"].(float64)
	longitude := p.Args["longitude"].(float64)
	date := p.Args["date"].(string)
	category := p.Args["category"].(string)
	image := p.Args["image"].(string)
	content := p.Args["content"].(string)


	dateConv, _ := time.Parse(time.RFC3339, date)

	event := models.UpdateEntertainent(id, title, price, location, latitude, longitude, dateConv, category, image,content)

	return event, nil
}

func DeleteEntertainment(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	event := models.DeleteEntertainment(id)

	return event, nil
}
