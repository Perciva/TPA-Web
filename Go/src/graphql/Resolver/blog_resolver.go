package Resolver

import (
	"errors"
	"github.com/graphql-go/graphql"
	"models"
)

func AllBlog(p graphql.ResolveParams) (i interface{}, err error) {

	blog := models.GetAllBlog()
	if len(blog) == 0 {
		return nil, errors.New("error: no data found")
	}

	return blog, nil
}

func InsertBlog(p graphql.ResolveParams) (i interface{}, err error) {
	userID := p.Args["userId"].(int)
	title := p.Args["title"].(string)
	content := p.Args["content"].(string)
	image := p.Args["image"].(string)
	category := p.Args["category"].(string)

	blog := models.InsertBlog(userID, title, content, image, category)

	return blog, nil
}

func UpdateBlog(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)
	content := p.Args["content"].(string)
	image := p.Args["image"].(string)
	category := p.Args["category"].(string)

	blog := models.UpdateBlog(id, content, image, category)

	return blog, nil
}
func GetBlogById(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	blog := models.GetBlogById(id)

	return blog, nil
}

func GetRecBlog(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	blog := models.GetRecBlog(id)

	return blog, nil
}
func DeleteBlog(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	blog := models.DeleteBlog(id)

	return blog, nil
}
