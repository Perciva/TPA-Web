package models
import (
	"Connection"
	"log"
	"time"
)

type Entertainment struct {
	Id        int `gorm:"primary_key"`
	Title     string  `gorm:"VARCHAR(100); NOT NULL;"`
	Price     int `gorm:"INTEGER"`
	Location  string  `gorm:"VARCHAR(100); NOT NULL;"`
	Latitude  float64 `gorm:"decimal(3,1)"`
	Longitude float64 `gorm:"decimal(3,1)"`
	Content string `gorm:"VARCHAR(100);"`
	Date      time.Time
	Type      string  `gorm:"VARCHAR(100); NOT NULL;"`
	Image     string  `gorm:"VARCHAR(100); NOT NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&Entertainment{})

	log.Println("Initialize Entertainment Success")
}

func GetAllEntertainment() []Entertainment {
	db := Connection.Connect()
	defer db.Close()

	var event []Entertainment
	if ValidateKey() == false {
		return event
	}
	db.Find(&event)

	return event
}

func GetEntertainmentByCategory(category string)[]Entertainment{
	db := Connection.Connect()
	defer db.Close()

	var res []Entertainment
	db.Where("type = ?", category).Limit(3).Find(&res)


	return res
}

func GetEntertainmentById(id int) Entertainment {
	db := Connection.Connect()
	defer db.Close()

	var event Entertainment
	if ValidateKey() == false {
		return event
	}
	db.Where("id = ?", id).First(&event)

	return event
}

func InsertEntertainment(title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string, desc string) *Entertainment {

	if len(desc)< 20 || len(title) < 5{
		return &Entertainment{Id:-1}
	}
	db := Connection.Connect()
	defer db.Close()

	newEvent := &Entertainment{
		Title:     title,
		Price:     price,
		Location:  location,
		Latitude:  latitude,
		Longitude: longitude,
		Date:      date.Add(-7 * time.Hour),
		Type:      category,
		Image:     image,
		Content:desc,
	}
	db.Save(newEvent)

	log.Println(newEvent,"Insert New Entertainment Success")
	return newEvent
}

func UpdateEntertainent(id int, title string, price int, location string, latitude float64, longitude float64,
	date time.Time, category string, image string, cont string) Entertainment {
	db := Connection.Connect()
	defer db.Close()

	var updateEvent Entertainment
	db.
		Model(&updateEvent).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"Title":     title,
			"Price":     price,
			"Location":  location,
			"Latitude":  latitude,
			"Longitude": longitude,
			"Date":      date.Add(-7 * time.Hour),
			"Type":      category,
			"Image":     image,
			"Content": cont,
		})

	log.Println("Update Entertainment Success")
	return updateEvent
}

func DeleteEntertainment(id int) *Entertainment {
	db := Connection.Connect()
	defer db.Close()

	var event Entertainment
	event = GetEntertainmentById(id)

	if event.Id == 0 {
		log.Println("Delete Blog Failed")
		return &event
	}

	err := db.Delete(event).Error

	if err != nil {
		panic("Error Delete Entertainment !" + err.Error())
	}

	log.Println("Delete Entertainment Success")
	return &event
}
