package models

import (
	"Connection"
	"log"
	"time"
)

type Blog struct {
	Id        int
	UserId    int
	Title     string
	Content   string
	ViewCount int
	Image     string
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	db := Connection.Connect()
	defer db.Close()

	db.
		AutoMigrate(&Blog{})

	log.Println("Initialize Blog Success")
}

func DropTableBlog() {
	db := Connection.Connect()
	defer db.Close()

	db.DropTableIfExists(&Blog{})
	db.AutoMigrate(&Blog{})

	log.Println("Drop db Success")
}

func GetAllBlog() []Blog {
	db := Connection.Connect()
	defer db.Close()

	var blogs []Blog
	db.Find(&blogs)

	return blogs
}

func GetBlogByID(id int) Blog {
	db := Connection.Connect()
	defer db.Close()

	var blog Blog
	db.Where("id = ?", id).First(&blog)

	return blog
}

func InsertBlog(userID int, title string, content string, image string, category string) *Blog {
	db := Connection.Connect()
	defer db.Close()

	newBlog := &Blog{
		UserId:    userID,
		Title:     title,
		Content:   content,
		ViewCount: 0,
		Image:     image,
		Category:  category,
	}
	db.Save(newBlog)

	log.Println("Insert New Blog Success")
	return newBlog
}

func UpdateBlog(id int, content string, image string, category string) Blog {
	db := Connection.Connect()
	defer db.Close()

	var updateBlog Blog
	db.
		Model(&updateBlog).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"Content":  content,
			"Image":    image,
			"Category": category,
		})

	log.Println("Update Blog Success")
	return updateBlog
}

func DeleteBlog(id int) *Blog {
	db := Connection.Connect()
	defer db.Close()

	var blog Blog
	blog = GetBlogByID(id)

	if blog.Id == 0 {
		log.Println("Delete Blog Failed")
		return &blog
	}

	err := db.Delete(blog).Error

	if err != nil {
		panic("Error Delete Hotel !" + err.Error())
	}

	log.Println("Delete Blog Success")
	return &blog
}
