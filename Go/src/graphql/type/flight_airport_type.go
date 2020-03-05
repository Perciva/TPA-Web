package types

import (
	"github.com/graphql-go/graphql"
)

var flightAirportType *graphql.Object

func GetFlightAirportType() *graphql.Object {
	if flightAirportType == nil {
		flightAirportType = graphql.NewObject(graphql.ObjectConfig{
			Name: "FlightAirportType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"locationID": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return flightAirportType
}
