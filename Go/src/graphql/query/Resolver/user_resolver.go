package Resolver

import (
	"github.com/graphql-go/graphql"
	"models"
)

func GetUsers(p graphql.ResolveParams)(i interface{}, e error){
	user, err := models.GetAllUser()

	return user,err
}

func GetUserByPhoneOrEmail(p graphql.ResolveParams)(i interface{},e error){
	arg:= p.Args["arg"].(string)

	user,err:= models.GetUserByEmailOrPhone(arg)

	return user,err
}
