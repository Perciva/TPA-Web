package types

import "github.com/graphql-go/graphql"

var hotel *graphql.Object

func GetHotelType() *graphql.Object{
	if hotel == nil{
		hotel = graphql.NewObject(graphql.ObjectConfig{
			Name: "Hotel",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"locationId": &graphql.Field{
					Type: graphql.Int,
				},
				"lat": &graphql.Field{
					Type: graphql.Float,
				},
				"long": &graphql.Field{
					Type: graphql.Float,
				},
				"photo": &graphql.Field{
					Type: graphql.NewList(GetHotelImageType()),
				},
				"facility": &graphql.Field{
					Type: graphql.NewList(GetHotelFacilityType()),
				},
				"type": &graphql.Field{
					Type: graphql.NewList(GetHotelTypeType()),
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"rating": &graphql.Field{
					Type: graphql.Float,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return hotel
}