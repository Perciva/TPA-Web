package mutation

import (
	"github.com/graphql-go/graphql"
	"graphql/mutation/Resolver"
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
		},
	})
}