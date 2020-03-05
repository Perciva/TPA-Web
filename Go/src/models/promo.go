package models

import (
	"Connection"
	"fmt"
	"time"
)

type Promo struct {
	Id        int `gorm:"PRIMARY_KEY"`
	Name      string `gorm:"VARCHAR(MAX)"`
	Detail    string `gorm:"VARCHAR(MAX)"`
	Image     string `gorm:"VARCHAR(MAX)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.AutoMigrate(&Promo{})
}

func GetPromoById(id int) Promo {
	db := Connection.Connect()
	defer db.Close()

	var res Promo
	db.Where("id = ?", id).First(&res)

	return res
}

func GetOtherPromo(id int) []Promo {
	fmt.Print(id)
	db := Connection.Connect()
	defer db.Close()

	var res []Promo
	db.Where("id != ?", id).Limit(3).Find(&res)

	fmt.Println("model",res)
	return res
}

