package Resolver

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"models"
)

func GetPromo(p graphql.ResolveParams) (i interface{}, err error) {
	id := p.Args["id"].(int)

	res := models.GetPromoById(id);
	return res, nil
}

func OtherPromo(p graphql.ResolveParams) (i interface{}, err error) {
	fmt.Println("in resolver")
	id := p.Args["id"].(int)

	res := models.GetOtherPromo(id)

	fmt.Println("resolve",res)
	return res, nil
}
