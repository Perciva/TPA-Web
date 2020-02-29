package models

import (
	"Connection"
	"time"
)

type User struct{
	Id 			int 		`gorm: "primary_key"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql : "index"`
	FirstName	string		`gorm: "type: varchar(100);not null"`
	LastName	string		`gorm: "type: varchar(100);not null"`
	Password	string		`gorm: "type: varchar(100);not null"`
	Email		string		`gorm: "type: varchar(100);not null"`
	Phone		string		`gorm: "type: varchar(100);not null"`
}

func init(){
	db := Connection.Connect()

	db.AutoMigrate(&User{})

}

func CreateUser(fname string, lname string, password string, email string, phone string)(*User){
	db := Connection.Connect()

	defer db.Close()

	var user = User{
		FirstName: fname,
		LastName: lname,
		Password:password,
		Email:email,
		Phone:phone,
	}
	if(db.NewRecord(user)){
		db.Create(&user)
	}

	return &user
}
func GetAllUser()([]User){
	db := Connection.Connect()
	defer db.Close()

	var users []User

	db.Find(&users)

	return users
}

func GetUserByEmailOrPhone(arg string)([]User){
	db := Connection.Connect()

	defer db.Close()

	var user []User

	if db.Where("email = ?", arg).Or("Phone = ?", arg).Find(&user).RecordNotFound(){
		return nil
	}

	return user
}
