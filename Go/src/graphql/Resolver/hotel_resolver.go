package Resolver

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
)

func GetAllHotel(p graphql.ResolveParams)(i interface{}, err error){
	hotels := models.GetAllHotel()

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

	hotel:= models.InsertHotel(name,address,loc, price, rating,lat,long,desc)

	fmt.Println(hotel,"Inserted -> resolver")

	return hotel, nil
}
func UpdateHotel(p graphql.ResolveParams)(i interface{}, err error){
	name := p.Args["name"].(string)
	id := p.Args["id"].(int)
	price := p.Args["price"].(int)
	rating := p.Args["rating"].(float64)
	desc := p.Args["description"].(string)

	hotel := models.UpdateHotel(id,name,price,rating,desc)

	fmt.Println(hotel, "updated -> resolver")

	return hotel, nil
}

func DeleteHotel(p graphql.ResolveParams)(i interface{}, err error){
	id := p.Args["id"].(int)

	hotel := models.DeleteHotel(id)
	fmt.Println(hotel, "deleted -> resolver")
	return hotel, nil
}
func GetHotelById(p graphql.ResolveParams)(i interface{}, err error){
	id := p.Args["id"].(int)
	hotel := models.GetHotel(id)

	return hotel, nil
}
func GetHotelByProvince(p graphql.ResolveParams)(i interface{}, err error){
	province := p.Args["province"].(string)

	hotel := models.GetHotelByProvince(province)

	if len(hotel) == 0{
		return nil, errors.New("no data found from get hotel by province")
	}

	return hotel, nil
}
func GetHotelByLatLong(p graphql.ResolveParams)(i interface{}, err error){
	lat := p.Args["lat"].(float64)
	long := p.Args["long"].(float64)

	hotel := models.GetHotelByLatLong(lat,long)


	return hotel, nil
}
func FilterHotel(p graphql.ResolveParams)(i interface{}, err error){
	fmt.Println("resolver")
	locations := p.Args["locations"].([]interface{})

	convLoc := make([]int, len(locations))
	for i := range locations {
		convLoc[i] = locations[i].(int)
	}

	ratings := p.Args["ratings"].([]interface{})
	convRating := make([]int, len(ratings))
	for i := range ratings {
		convRating[i] = ratings[i].(int)
	}

	facilities := p.Args["facilities"].([]interface{})
	convFaci := make([]string, len(facilities))
	for i := range facilities {
		convFaci[i] = facilities[i].(string)
	}

	min:= p.Args["min"].(int)
	max:= p.Args["max"].(int)


	hotel:= models.HotelFilter(convLoc,convRating, convFaci,min,max)

	return hotel,nil

}

func GetNearbyHotel(p graphql.ResolveParams)(i interface{}, err error){
	lat := p.Args["lat"].(float64)
	long := p.Args["long"].(float64)

	hotel := models.GetNearbyHotel(lat,long)

	if len(hotel) == 0{
		return nil, errors.New("No Nearby Hotel Found")
	}
	return hotel, nil
}