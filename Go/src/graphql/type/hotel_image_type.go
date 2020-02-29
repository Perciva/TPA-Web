package types

import "github.com/graphql-go/graphql"

var imageType *graphql.Object

func GetHotelImageType() *graphql.Object{
	if imageType == nil {
		imageType = graphql.NewObject(graphql.ObjectConfig{
			Name:"HotelImageType",
			Fields: graphql.Fields{
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
	return imageType
}
