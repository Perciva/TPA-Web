package query

import(
	"github.com/graphql-go/graphql"
	"graphql/Resolver"
	"graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject( graphql.ObjectConfig{
		Name:        "RootQuery",
		Fields:      graphql.Fields{
			"users":{
				Type:        graphql.NewList(types.GetUserType()),
				Resolve:     Resolver.GetUsers,
				Description: "All Users",
			},
			"userbyemailphone":{
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"arg": &graphql.ArgumentConfig{
						Type:         graphql.String,
						Description:  "phone or email",
					},
				},
				Resolve: Resolver.GetUserByPhoneOrEmail,
			},
			"allHotel":{
				Type: graphql.NewList(types.GetHotelType()),
				Resolve: Resolver.GetAllHotel,
			},
			"hotelbyprovince":{
				Type: graphql.NewList(types.GetHotelType()),
				Resolve: Resolver.GetHotelByProvince,
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Province",
					},
				},
			},
			"hotelbylatlong":{
				Type: types.GetHotelType(),
				Resolve: Resolver.GetHotelByLatLong,
				Args: graphql.FieldConfigArgument{
					"lat" : &graphql.ArgumentConfig{
						Type:         graphql.Float,
					},
					"long" : &graphql.ArgumentConfig{
						Type:         graphql.Float,
					},
				},
			},
			"hotelbyid":{
				Type: types.GetHotelType(),
				Resolve: Resolver.GetHotelById,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        graphql.Int,
					},
				},
			},
			"getlocbyprovince":{
				Type: graphql.NewList(types.GetLocationType()),
				Resolve: Resolver.GetLocByProvince,
				Args: graphql.FieldConfigArgument{
					"province": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "Province",
					},
				},
			},
			"getlocbycity":{
				Type: types.GetLocationType(),
				Resolve: Resolver.GetLocByCity,
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type:        graphql.String,
						Description: "city",
					},
				},
			},
			"getallloc":{
				Type: graphql.NewList(types.GetLocationType()),
				Resolve: Resolver.GetAllLoc,
			},
			"getallhotelimage":{
				Type: graphql.NewList(types.GetHotelImageType()),
				Resolve: Resolver.GetAllHotelImage,
			},
			"getallhoteltype":{
				Type: graphql.NewList(types.GetHotelTypeType()),
				Resolve: Resolver.GetAllHotelType,
			},
			"getallhotelfacility":{
				Type: graphql.NewList(types.GetHotelFacilityType()),
				Resolve: Resolver.GetAllHotelfac,
			},
			"getnearbyhotel":{
				Type: graphql.NewList(types.GetHotelType()),
				Resolve: Resolver.GetNearbyHotel,
				Args: graphql.FieldConfigArgument{
					"lat" : &graphql.ArgumentConfig{
						Type:         graphql.Float,
					},
					"long" : &graphql.ArgumentConfig{
						Type:         graphql.Float,
					},
				},
			},
			"getfilteredhotel":{
				Type: graphql.NewList(types.GetHotelType()),
				Resolve: Resolver.FilterHotel,
				Args: graphql.FieldConfigArgument{
					"locations" : &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"ratings" : &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"facilities": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"min": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"max": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"getallcarmodel": {
				Type:        graphql.NewList(types.GetCarModelType()),
				Resolve:     Resolver.AllCarModel,
			},
			"getallcar": {
				Type:        graphql.NewList(types.GetCarType()),
				Resolve:     Resolver.AllCar,
			},
			"getcarbycity": {
				Type: graphql.NewList(types.GetCarType()),
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve:     Resolver.GetCarByCity,
			},
			"getalltrain": {
				Type:        graphql.NewList(types.GetTrainType()),
				Resolve:     Resolver.AllTrain,
			},
			"gettrainbyloc": {
				Type:        graphql.NewList(types.GetTrainType()),
				Resolve:     Resolver.GetTrainByLoc,
				Args: graphql.FieldConfigArgument{
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"date": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"getrecblog": {
				Type:        graphql.NewList(types.GetBlogType()),
				Resolve:     Resolver.GetRecBlog,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"getblogbyid": {
				Type:    types.GetBlogType(),
				Resolve: Resolver.GetBlogById,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"getpromobyid": {
				Type: types.GetPromoType(),
				Resolve:  Resolver.GetPromo,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"getotherpromo": {
				Type: graphql.NewList(types.GetPromoType()),
				Resolve:   Resolver.OtherPromo,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
			"getalltrainclass": {
				Type:        graphql.NewList(types.GetTrainClassType()),
				Resolve:     Resolver.AllTrainClass,
			},
			"gettrainnames": {
				Type:        graphql.NewList(types.GetTrainType()),
				Resolve:     Resolver.GetTrainName,
			},
			"getalltrainstation": {
				Type:        graphql.NewList(types.GetTrainStationType()),
				Resolve:     Resolver.AllTrainStation,
			},
			"getallflight": {
				Type:        graphql.NewList(types.GetFlightType()),
				Resolve:     Resolver.AllFlight,
			},
			"getallentertainment": {
				Type:        graphql.NewList(types.GetEntertainmentType()),
				Resolve:     Resolver.AllEntertainment,
			},
			"getallflightfacility": {
				Type:        graphql.NewList(types.GetFlightFacilityType()),
				Resolve:     Resolver.AllFlightFacility,
			},
			"getallflightcompany": {
				Type:        graphql.NewList(types.GetFlightCompanyType()),
				Resolve:     Resolver.AllFlightCompany,
			},
			"getallflightairport": {
				Type:        graphql.NewList(types.GetFlightAirportType()),
				Resolve:     Resolver.AllFlightAirport,
			},
			"getallblog": {
				Type:        graphql.NewList(types.GetBlogType()),
				Resolve:     Resolver.AllBlog,
			},
		},
		IsTypeOf:    nil,
		Description: "",
	})
}
