package types

import "github.com/graphql-go/graphql"

var facType *graphql.Object

func GetHotelFacilityType() *graphql.Object{
	if facType == nil {
		facType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelFacilityType",
			Fields: graphql.Fields{
				"name" : &graphql.Field{
					Type: graphql.String,
				},
				"id" : &graphql.Field{
					Type: graphql.Int,
				},
				"hotelId": &graphql.Field{
					Type: graphql.Int,
				},
				"path" : &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return facType
}
