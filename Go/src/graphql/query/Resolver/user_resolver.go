package Resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"models"
)

func GetUsers(p graphql.ResolveParams)(i interface{}, err error){
	user := models.GetAllUser()
	if len(user) == 0{
		return nil, errors.New("no data found")
	}
	return user, nil
}

func GetUserByPhoneOrEmail(p graphql.ResolveParams)(i interface{}, err error){
	arg:= p.Args["arg"].(string)

	user := models.GetUserByEmailOrPhone(arg)
	if len(user) == 0{
		return nil, errors.New("no data found")
	}
	return user, nil
}
