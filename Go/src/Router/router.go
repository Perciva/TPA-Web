package Router

import(

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.Use(middleware.LogMiddleware)




	return r
}