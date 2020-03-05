package types

import (
	"github.com/graphql-go/graphql"
)

var flightType *graphql.Object

func GetFlightType() *graphql.Object {
	if flightType == nil {
		flightType = graphql.NewObject(graphql.ObjectConfig{
			Name: "FlightType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"company": &graphql.Field{
					Type: GetFlightCompanyType(),
				},
				"model": {
					Type: graphql.String,
				},
				"companyId": &graphql.Field{
					Type: graphql.Int,
				},
				"fromAirport": &graphql.Field{
					Type: GetFlightAirportType(),
				},
				"fromAirportId": &graphql.Field{
					Type: graphql.Int,
				},
				"toAirport": &graphql.Field{
					Type: GetFlightAirportType(),
				},
				"toAirportId": &graphql.Field{
					Type: graphql.Int,
				},
				"arrivalTime": &graphql.Field{
					Type: graphql.String,
				},
				"departureTime": &graphql.Field{
					Type: graphql.String,
				},
				"price": &graphql.Field{
					Type: graphql.String,
				},
				"facility": &graphql.Field{
					Type: graphql.NewList(GetFlightFacilityType()),
				},
				"transit": &graphql.Field{
					Type: GetFlightAirportType(),
				},
				"transitId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return flightType
}
