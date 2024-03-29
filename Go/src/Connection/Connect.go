package Connection

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const dbName = "TPAWeb"
const dbHost = "127.0.0.1"
const dbPort = "5432"
const dbUser = "postgres"
const dbPassword = "farfetch"

func Connect()(*gorm.DB){

	db,_ := gorm.Open("postgres","host=" + dbHost+" port="+dbPort +" user="+dbUser+" dbname="+dbName+" password="+dbPassword+" sslmode=disable")
	return db
}
