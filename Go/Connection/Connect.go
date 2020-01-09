package Connection

import "github.com/jinzhu/gorm"

const dbName = "TPAWeb"
const dbHost = "127.0.0.1"
const dbPort = "5432"
const dbUser = "postgres"
const dbPassword = ""

func Connect()(*gorm.DB, error){
	return gorm.Open("postgres","host=" + dbHost+" port="+dbPort +" user="+dbUser+" dbname="+dbName+" password="+dbPassword+" sslmode=disable")
}
