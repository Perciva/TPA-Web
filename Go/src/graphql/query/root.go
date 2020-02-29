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


		},
		IsTypeOf:    nil,
		Description: "",
	})
}
