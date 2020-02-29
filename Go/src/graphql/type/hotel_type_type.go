package types

import "github.com/graphql-go/graphql"

var hotelType *graphql.Object

func GetHotelTypeType() *graphql.Object{

	if hotelType == nil {
		hotelType = graphql.NewObject(graphql.ObjectConfig{
			Name: "HotelType",
			Fields: graphql.Fields{
				"id" : &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"path" : &graphql.Field{
					Type: graphql.String,
				},
				"hotelId": &graphql.Field{
					Type: graphql.Int,
				},
			},

		})
	}


	return hotelType
}
