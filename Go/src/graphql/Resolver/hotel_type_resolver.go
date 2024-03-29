package Resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"models"
)

func GetAllHotelType(p graphql.ResolveParams)(i interface{}, err error){
	res := models.GetAllHotelType()

	if len(res) == 0{
		return nil, errors.New("No Data Found in Get All Htel Type")
	}
	return res,nil
}
func CreateHotelType(p graphql.ResolveParams)(i interface{}, err error){
	hotelId := p.Args["hotelId"].(int)
	name:= p.Args["name"].(string)
	path:= p.Args["path"].(string)
	newHotelType := models.InsertHotelType(hotelId,name,path)

	return newHotelType, nil
}


