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
			"createtrain": &graphql.Field{
				Type: types.GetTrainType(),
				Resolve: Resolver.InsertTrain,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"transit": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"seat": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createtrainclass": &graphql.Field{
				Type: types.GetTrainClassType(),
				Resolve: Resolver.InsertTrainClass,
				Args: graphql.FieldConfigArgument{
					"trainId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createtrainstation": &graphql.Field{
				Type: types.GetTrainStationType(),
				Resolve: Resolver.InsertTrainStation,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"code": &graphql.ArgumentConfig{
						Type:graphql.String,
					},
				},
			},
			"updatetrain": &graphql.Field{
				Type: types.GetTrainType(),
				Resolve: Resolver.UpdateTrain,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"seat": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"deletetrain": &graphql.Field{
				Type: types.GetTrainType(),
				Resolve: Resolver.DeleteTrain,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createcarmodel": &graphql.Field{
				Type: types.GetCarModelType(),
				Resolve: Resolver.InsertCarModel,
				Args: graphql.FieldConfigArgument{
					"brand": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"passenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"baggage": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createnewcar": &graphql.Field{
				Type: types.GetCarType(),
				Resolve: Resolver.InsertCar,
				Args: graphql.FieldConfigArgument{
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createflight": &graphql.Field{
				Type: types.GetFlightType(),
				Resolve: Resolver.InsertFlight,
				Args: graphql.FieldConfigArgument{
					"company": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"fromAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"toAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"transitAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createflightairport": &graphql.Field{
				Type: types.GetFlightAirportType(),
				Resolve: Resolver.InsertFlightAirport,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createflightcompany": &graphql.Field{
				Type: types.GetFlightCompanyType(),
				Resolve: Resolver.InsertFlightCompany,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"createflightfacility": &graphql.Field{
				Type: types.GetFlightType(),
				Resolve: Resolver.InsertFlightFacility,
				Args: graphql.FieldConfigArgument{
					"flightID": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"updateflight": &graphql.Field{
				Type: types.GetFlightType(),
				Resolve: Resolver.UpdateFlight,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"fromAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"toAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"transitAirport": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"arrivalTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departureTime": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"deleteflight": &graphql.Field{
				Type: types.GetFlightType(),
				Resolve: Resolver.DeleteFlight,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createblog": &graphql.Field{
				Type: types.GetBlogType(),
				Resolve: Resolver.InsertBlog,
				Args: graphql.FieldConfigArgument{
					"userId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"updateblog": &graphql.Field{
				Type: types.GetBlogType(),
				Resolve: Resolver.UpdateBlog,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"deleteblog": &graphql.Field{
				Type: types.GetBlogType(),
				Resolve: Resolver.DeleteBlog,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"createevent": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Resolve: Resolver.InsertEntertainment,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"updateevent": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Resolve: Resolver.UpdateEntertainment,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content":&graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"deleteevent": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Resolve: Resolver.DeleteEntertainment,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	})
}