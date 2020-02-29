package Resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"models/hotel"
)

func GetAllHotelfac(p graphql.ResolveParams)(i interface{}, err error){
	res:= hotel.GetAllHotelFacility()

	if(len(res) == 0){
		return nil, errors.New("No Data Found")
	}

	return res, nil
}

func CreateHotelFac(p graphql.ResolveParams)(i interface{}, err error){
	hotelid := p.Args["hotelId"].(int)
	name:= p.Args["name"].(string)
	path:= p.Args["path"].(string)

	res := hotel.InsertHotelFacility(hotelid, name, path)

	return res,nil
}

