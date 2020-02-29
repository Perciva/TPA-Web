package Resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
)

func GetAllLoc(p graphql.ResolveParams)(i interface{}, err error){
	loc := models.GetAllLocation()

	if len(loc) == 0{
		return nil, errors.New("No Location Found")
	}

	return loc,nil
}
func CreateLocation(p graphql.ResolveParams)(i interface{}, err error){
	city := p.Args["city"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)

	loc := models.InsertLocation(city, province, country)

	return loc,nil
}
func GetLocByProvince(p graphql.ResolveParams)(i interface{}, err error){
	province := p.Args["province"].(string)
	loc := models.GetLocByProvince(province)

	if len(loc) == 0{
		return nil, errors.New("location by province not found")
	}
	return loc,nil
}
func GetLocByCity(p graphql.ResolveParams)(i interface{}, err error){
	city := p.Args["city"].(string)
	loc := models.GetLocByCity(city)

	if loc.Id == 0{
		return nil, errors.New("location by city not found")
	}
	fmt.Println(loc,"Fetched By GetLocByCity")
	return loc,nil
}