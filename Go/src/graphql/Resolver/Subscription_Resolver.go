package Resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
)

func Subscribe(p graphql.ResolveParams) (i interface{}, err error) {
	fmt.Println("in resolver")
	email := p.Args["email"].(string)

	emailArr := []string{email}

	res := models.SubscriptionEmail(emailArr)

	fmt.Println("resolve",res)
	return res, nil
}