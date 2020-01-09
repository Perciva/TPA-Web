package main

import (
	"fmt"
	"github.com/Perciva/TPA-Web/Go/Connection"

)

func main(){
	fmt.Print("Hello")

	db,err:= Connection.Connect()
	if err != nil{
		panic("Database failed to connect " + err.Error())
	}
	fmt.Print(db)

}
