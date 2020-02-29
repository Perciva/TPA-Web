package mutation

import (
	"github.com/graphql-go/graphql"
	"graphql/Resolver"
	"graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootMutation",
		Fields: graphql.Fields{
			"createuser":{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: Resolver.CreateUser,
				Description: "Get User Based on Email",
			},
			"createhotel":{
				Type: types.GetHotelType(),
				Args:graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"lat" : &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"long": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve:Resolver.CreateHotel,
			},
			"updatehotel":{
				Type: types.GetHotelType(),
				Resolve:Resolver.UpdateHotel,
				Args:graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"deletehotel":{
				Type: types.GetHotelType(),
				Resolve:Resolver.DeleteHotel,
				Args:graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createlocation":{
				Type: types.GetLocationType(),
				Resolve: Resolver.CreateLocation,
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createhotelimage":{
				Type: types.GetHotelImageType(),
				Resolve:Resolver.CreateHotelImage,
				Args:graphql.FieldConfigArgument{
					"hotelId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"path": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createhotelfacility":{
				Type:types.GetHotelFacilityType(),
				Resolve:Resolver.CreateHotelFac,
				Args:graphql.FieldConfigArgument{
					"hotelId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"path":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createhoteltype":{
				Type: types.GetHotelTypeType(),
				Resolve:Resolver.CreateHotelType,
				Args: graphql.FieldConfigArgument{
					"hotelId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"path":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
		},
	})
}