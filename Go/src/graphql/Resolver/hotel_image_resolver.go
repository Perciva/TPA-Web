package Resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"models"
)

func GetAllHotelImage(p graphql.ResolveParams)( i interface{}, err error){
	res :=models.GetAllHotelImages()

	if len(res)==0{
		return nil, errors.New("Hotel Image Data Not Found")
	}
	return res,nil
}

func CreateHotelImage(p graphql.ResolveParams)(i interface{}, err error){
	hotelId := p.Args["hotelId"].(int)
	path := p.Args["path"].(string)

	newImage := models.InsertHotelImage(hotelId,path)

	return  newImage,nil
}

