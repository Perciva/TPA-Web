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
	Languange string `gorm: "type: varchar(100);"`
	Title	string `gorm: "type: varchar(100);"`
	Address string `gorm: "type: varchar(100);"`
	KodePos string `gorm: "type: varchar(100);"`

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
func UpdateUser(id int,fname string, lname string, email string, phone string, lang string, title string, address string,kodepos string)(User){
	db:= Connection.Connect()
	defer db.Close()

	var res User

	db.
		Model(&res).
		Where("Id = ?", id).
		Updates(map[string]interface{}{
			"FirstName": fname,
			"LastName" : lname,
			"Email": email,
			"Phone": phone,
			"Languange":lang,
			"Title":title,
			"Address": address,
			"KodePos": kodepos,
		})

	return res

}
func GetAllUser()([]User){
	db := Connection.Connect()
	defer db.Close()

	var users []User
	if ValidateKey() == false {
		return users
	}

	db.Find(&users)

	return users
}

func GetUserByEmailOrPhone(arg string)([]User){
	db := Connection.Connect()
	defer db.Close()

	var user []User
	if ValidateKey() == false {
		return user
	}

	if db.Where("email = ?", arg).Or("Phone = ?", arg).Find(&user).RecordNotFound(){
		return nil
	}

	return user
}

func GetUserByID(id int)(User){
	db := Connection.Connect()
	defer db.Close()

	var user User
	if ValidateKey() == false {
		return user
	}

	db.Where("id = ?", id).First(&user)

	return user
}
