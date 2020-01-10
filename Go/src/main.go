package main

import (
	//"Middleware"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"graphql/mutation"
	"graphql/query"
	"log"
	"net/http"
)

func main(){
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query.GetRoot(),
		Mutation: mutation.GetRoot(),
	})

	if err != nil{
		panic(err.Error())
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty:	true,
		GraphiQL: true,
		Playground: true,
	})

	//wrapped := Middleware.CorsMiddleware(h)

	//router := api.NewRouter()
	//router.Handle("/api",wrapped)
	log.Fatalln(http.ListenAndServe(":4100",h))
}
