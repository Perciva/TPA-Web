package models

import (
	"Connection"
	"fmt"
	"math"
	"time"
)

type Hotel struct {
	Id          int             `gorm:"PRIMARY_KEY"`
	Name        string          `gorm:"VARCHAR(MAX)"`
	Address     string          `gorm:"VARCHAR(MAX)"`
	Location    Location   `gorm:"foreignkey:LocationId"`
	LocationId  int             `gorm:"Integer"`
	Lat         float64         `gorm:"decimal(3,1)"`
	Long        float64         `gorm:"decimal(3,1)"`
	Photo       []HotelImage    `gorm:"foreignkey:HotelId"`
	Facility    []HotelFacility `gorm:"foreignkey:HotelId"`
	Type        []HotelType     `gorm:"foreignkey:HotelId"`
	Price       int             `gorm:"Integer"`
	Rating      float64         `gorm:"DECIMAL(3,1)"`
	Description string          `gorm:"VARCHAR(MAX)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql : "index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()
	fmt.Println(":D")
	db.AutoMigrate(&Hotel{}).
		AddForeignKey("location_Id", "locations(id)", "cascade", "cascade")
	fmt.Println("Hotel Created")
}
func dropHotel() {
	db := Connection.Connect()
	defer db.Close()

	db.DropTableIfExists(&Hotel{})
	db.AutoMigrate(&Hotel{}).AddForeignKey("LocationId", "locations(id)", "cascade", "cascade")
}
func HotelFilter(locations []int,ratings []int , fac []string, min int, max int)[]Hotel{
	db := Connection.Connect()
	defer db.Close()

	fmt.Println(locations,ratings)
	var res []Hotel
	if ValidateKey() == false {
		return res
	}
	db.Joins("JOIN hotel_facilities on hotels.id = hotel_facilities.hotel_id").
		Find(&res, "(location_id IN (?) OR ?) AND (floor(rating) IN (?) OR ?) AND (hotel_facilities.name IN (?) OR ?) AND (price >= ? OR ? = -1) AND (price <= ? OR ? = -1)",locations, len(locations) == 0, ratings,len(ratings) == 0, fac, len(fac)==0,min,min,max,max)

	for z, _ := range res {
		db.Model(&res[z]).Related(&res[z].Location, "location_Id")
		db.Model(&res[z]).Related(&res[z].Photo, "HotelId")
		db.Model(&res[z]).Related(&res[z].Facility, "HotelId")
		db.Model(&res[z]).Related(&res[z].Type, "HotelId")
	}

	fmt.Println("a",res)
	return res
}

func GetAllHotel() []Hotel {
	db := Connection.Connect()
	defer db.Close()

	var res []Hotel
	if ValidateKey() == false {
		return res
	}
	db.Find(&res)

	for z, _ := range res {
		db.Model(res[z]).Related(&res[z].Location, "location_Id")
		db.Model(res[z]).Related(&res[z].Photo, "HotelId")
		db.Model(res[z]).Related(&res[z].Facility, "HotelId")
		db.Model(res[z]).Related(&res[z].Type, "HotelId")
	}
	return res
}

func GetHotel(Id int) Hotel {
	db := Connection.Connect()
	defer db.Close()

	var res Hotel
	if ValidateKey() == false {
		return res
	}

	db.Where("id=?", Id).First(&res)

	db.Model(&res).Related(&res.Location, "location_Id")
	db.Model(&res).Related(&res.Photo, "HotelId")
	db.Model(&res).Related(&res.Facility, "HotelId")
	db.Model(&res).Related(&res.Type, "HotelId")

	fmt.Println("model GetHotel",res)
	return res
}
func UpdateHotel(id int, name string, price int, rating float64, desc string) *Hotel {
	db := Connection.Connect()
	defer db.Close()

	var res Hotel
	db.Model(&res).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":        name,
		"price":       price,
		"rating":      rating,
		"Description": desc,
	})
	db.Model(res).Related(&res.Location, "location_Id")
	db.Model(res).Related(&res.Photo, "HotelId")
	db.Model(res).Related(&res.Facility, "HotelId")
	db.Model(res).Related(&res.Type, "HotelId")


	return &res
}
func DeleteHotel(id int) *Hotel {
	db := Connection.Connect()
	defer db.Close()

	var res Hotel
	res = GetHotel(id)
	if res.Id == 0 {
		fmt.Println("Delete Hotel Failed")
		return &res
	}
	err := db.Delete(res).Error
	if err != nil {
		fmt.Println("Error Occured From dleetHotel: ", err)
	}
	db.Model(res).Related(&res.Location, "location_Id")
	db.Model(res).Related(&res.Photo, "HotelId")
	db.Model(res).Related(&res.Facility, "HotelId")
	db.Model(res).Related(&res.Type, "HotelId")
	return &res
}
func InsertHotel(name string, address string, city string, price int, rating float64, lat float64, long float64, desc string) *Hotel {
	db := Connection.Connect()
	defer db.Close()

	if len(desc) < 20{
		return &Hotel{Id:-1}
	}

	loc := GetLocByCity(city)

	newHotel := &Hotel{
		Name:        name,
		Address:     address,
		LocationId:  loc.Id,
		Lat:         lat,
		Long:        long,
		Price:       price,
		Rating:      rating,
		Description: desc,
	}
	db.Save(newHotel)

	db.Model(newHotel).Related(&newHotel.Location, "location_Id")
	db.Model(newHotel).Related(&newHotel.Photo, "HotelId")
	db.Model(newHotel).Related(&newHotel.Facility, "HotelId")
	db.Model(newHotel).Related(&newHotel.Type, "HotelId")

	fmt.Println(newHotel, "inserted")
	return newHotel
}
func GetHotelByProvince(province string) []Hotel {
	db := Connection.Connect()
	defer db.Close()

	loc := GetLocByProvince(province)

	var res []Hotel
	if ValidateKey() == false {
		return res
	}

	if len(loc) == 1 {
		db.Where("location_id = ?", loc[0].Id).Find(&res)
	} else {
		var locId []int
		for z := 0; z < len(loc); z++ {
			locId = append(locId, loc[z].Id)
		}
		db.Where("location_id IN (?)", locId).Find(&res)
	}
	for z, _ := range res {
		db.Model(res[z]).Related(&res[z].Location, "location_Id")
		db.Model(res[z]).Related(&res[z].Photo, "HotelId")
		db.Model(res[z]).Related(&res[z].Facility, "HotelId")
		db.Model(res[z]).Related(&res[z].Type, "HotelId")
	}
	return res
}

func GetHotelByLatLong(lat float64, long float64) Hotel {
	db := Connection.Connect()
	defer db.Close()

	var res Hotel
	if ValidateKey() == false {
		return res
	}
	db.Where("lat = ?", lat).Where("long = ?", long).First(&res)

	db.Model(res).Related(&res.Location, "location_Id")
	db.Model(res).Related(&res.Photo, "HotelId")
	db.Model(res).Related(&res.Facility, "HotelId")
	db.Model(res).Related(&res.Type, "HotelId")

	return res
}

func distance(lat float64, long float64, tarLat float64, tarLong float64) float64{
	radLatitude := float64(math.Pi * lat/180)
	radLatitudeTar := float64(math.Pi * tarLat/180)

	theta := float64(long - tarLong)
	radTheta := float64(math.Pi * theta /180)

	dist := math.Sin(radLatitude)* math.Sin(radLatitudeTar) + math.Cos(radLatitude * math.Cos(radLatitudeTar) * math.Cos(radTheta))

	if dist>1{
		dist=1
	}
	dist = math.Acos(dist)
	dist = dist*180/math.Pi
	dist = dist * 60 * 1.1515
	return dist * 1.609344
}
func distanceCmp(lat float64, long float64, hot1 Hotel, hot2 Hotel) bool {
	dist1 := distance(lat,long,hot1.Lat, hot1.Long)
	dist2 := distance(lat,long,hot2.Lat, hot2.Long)

	return dist1 < dist2
}
func srot(lat float64, long float64, hotels []Hotel){
	len := len(hotels)

	for z := 0; z<len;z++{
		for x := len-1; x>= z+1 ; x--{
			if distanceCmp(lat,long,hotels[x],hotels[x-1]){
				hotels[x], hotels[x-1] = hotels[x-1], hotels[x]
			}
		}
	}
}

func GetNearbyHotel(lat float64, long float64)[]Hotel {
	fmt.Println("sent lat long",lat,long)
	db := Connection.Connect()
	defer db.Close()

	var res []Hotel
	if ValidateKey() == false {
		return res
	}

	db.Find(&res)

	for z,_ := range res {
		db.Model(res[z]).Related(&res[z].Location, "location_Id")
		db.Model(res[z]).Related(&res[z].Photo, "HotelId")
		db.Model(res[z]).Related(&res[z].Facility, "HotelId")
		db.Model(res[z]).Related(&res[z].Type, "HotelId")
	}

	srot(lat,long,res)
	fmt.Println(res,"Fetchednearby hotels")
	if len(res) > 8{
		fmt.Println("Res Sliced")
		res = res[0:8]
	}
	return res
}
