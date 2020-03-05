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
func UpdateUser(p graphql.ResolveParams)(interface{},error){
	id:= p.Args["id"].(int)
	fname:=p.Args["firstname"].(string)
	lname:=p.Args["lastname"].(string)
	email:=p.Args["email"].(string)
	phone:=p.Args["phone"].(string)
	lang:=p.Args["languange"].(string)
	title:=p.Args["title"].(string)
	address:=p.Args["address"].(string)
	kodepos:=p.Args["kode_pos"].(string)

	user := models.UpdateUser(id,fname,lname,email,phone,lang,title,address,kodepos)

	return user,nil
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

func GetUserByID(p graphql.ResolveParams)(interface{},error){
	id:=p.Args["id"].(int)

	user := models.GetUserByID(id)

	return user, nil
}
