package Resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"models/hotel"
)

func GetAllHotel(p graphql.ResolveParams)(i interface{}, err error){
	hotels := hotel.GetAllHotel()

	if len(hotels) == 0{
		return nil, errors.New("no data found in get all hotel")
	}

	fmt.Println("resolver GetHotel",hotels)
	return hotels, nil
}

func CreateHotel(p graphql.ResolveParams)(i interface{}, err error){
	name := p.Args["name"].(string)
	address := p.Args["address"].(string)
	loc := p.Args["location"].(string)
	price := p.Args["price"].(int)
	rating := p.Args["rating"].(float64)
	lat := p.Args["lat"].(float64)
	long := p.Args["long"].(float64)
	desc := p.Args["description"].(string)

	hotel:= hotel.InsertHotel(name,address,loc, price, rating,lat,long,desc)

	fmt.Println(hotel,"Inserted -> resolver")

	return hotel, nil
}
func UpdateHotel(p graphql.ResolveParams)(i interface{}, err error){
	name := p.Args["name"].(string)
	id := p.Args["id"].(int)
	price := p.Args["price"].(int)
	rating := p.Args["rating"].(float64)
	desc := p.Args["description"].(string)

	hotel := hotel.UpdateHotel(id,name,price,rating,desc)

	fmt.Println(hotel, "updated -> resolver")

	return hotel, nil
}

func DeleteHotel(p graphql.ResolveParams)(i interface{}, err error){
	id := p.Args["id"].(int)

	hotel := hotel.DeleteHotel(id)
	fmt.Println(hotel, "deleted -> resolver")
	return hotel, nil
}
func GetHotelById(p graphql.ResolveParams)(i interface{}, err error){
	id := p.Args["id"].(int)
	hotel := hotel.GetHotel(id)

	return hotel, nil
}
func GetHotelByProvince(p graphql.ResolveParams)(i interface{}, err error){
	province := p.Args["province"].(string)

	hotel := hotel.GetHotelByProvince(province)

	if len(hotel) == 0{
		return nil, errors.New("no data found from get hotel by province")
	}

	return hotel, nil
}
func GetHotelByLatLong(p graphql.ResolveParams)(i interface{}, err error){
	lat := p.Args["lat"].(float64)
	long := p.Args["long"].(float64)

	hotel := hotel.GetHotelByLatLong(lat,long)


	return hotel, nil
}

func GetNearbyHotel(p graphql.ResolveParams)(i interface{}, err error){
	lat := p.Args["lat"].(float64)
	long := p.Args["long"].(float64)

	hotel := hotel.GetNearbyHotel(lat,long)

	if len(hotel) == 0{
		return nil, errors.New("No Nearby Hotel Found")
	}
	return hotel, nil
}