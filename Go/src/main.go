package main

import (
	"Connection"
	"middleware"

	//"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	//"github.com/gorilla/handlers"
	"graphql/mutation"
	"graphql/query"
	"net/http"
)
//func Cors(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html; charset=ascii")
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
//	w.Write([]byte("Hello, World!"))
//}


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

	tes := middleware.Cors(h)

	//models.InLoc()
	//models.InTrainStation()
	//models.InFlight()
	//models.InFlightCompany()
	router := Connection.NewRouter()
	router.Handle("/{key}", tes)
	log.Fatal(http.ListenAndServe(":4100",router))
	//router := router.NewRouter()
	//router.HandleFunc("/api",Cors)
	//router.Handle("/",h)
	//log.Fatalln(http.ListenAndServe(":4100",handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(h)))

	//http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Access-Control-Allow-Origin", "*")
	//	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	//	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	//
	//	if r.Method == "OPTIONS" {
	//		w.Write([]byte("allowed"))
	//		return
	//	}
	//
	//
	//})
	//
	//http.ListenAndServe(":4100", nil)
}
