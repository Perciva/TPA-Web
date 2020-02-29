package Resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
)

func GetUsers(p graphql.ResolveParams)(i interface{}, err error){
	user := models.GetAllUser()

	return user, nil
}

func GetUserByPhoneOrEmail(p graphql.ResolveParams)(i interface{}, err error){
	arg:= p.Args["arg"].(string)

	user := models.GetUserByEmailOrPhone(arg)
	fmt.Println(user)
	//if len(user) == 0{
	//	return nil, errors.New("no data found")
	//}
	return user, nil
}
func CreateUser(p graphql.ResolveParams)(interface{},error){
	firstname:=p.Args["firstname"].(string)
	lastname:=p.Args["lastname"].(string)
	password:=p.Args["password"].(string)
	email:=p.Args["email"].(string)
	phone:=p.Args["phone"].(string)

	user := models.CreateUser(firstname,lastname,password,email,phone)

	return user, nil
}
