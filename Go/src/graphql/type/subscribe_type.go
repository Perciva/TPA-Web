package types

import "github.com/graphql-go/graphql"

var subscribeType *graphql.Object

func GetSubscribeType() *graphql.Object {
	if subscribeType == nil {
		subscribeType = graphql.NewObject(graphql.ObjectConfig{
			Name: "subscribeType",
			Fields: graphql.Fields{
				"email": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return subscribeType
}
