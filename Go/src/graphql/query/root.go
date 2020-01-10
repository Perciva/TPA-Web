package query

import(
	"github.com/graphql-go/graphql"
	"graphql/query/Resolver"
	"graphql/type"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject( graphql.ObjectConfig{
		Name:        "RootQuery",
		Fields:      graphql.Fields{
			"users":{
				Type: graphql.NewList(types.GetUserType()),
				Resolve: Resolver.GetUsers,
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


		},
		IsTypeOf:    nil,
		Description: "",
	})
}
