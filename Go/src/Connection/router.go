package Connection

import (
	"github.com/gorilla/mux"
	//"net/http"
)
//
//func CorsMiddleware(next http.Handler) http.Handler{
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "*")
//		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
//
//		next.ServeHTTP(w, r)
//	})
//}
func NewRouter() *mux.Router{
	r := mux.NewRouter()
	//r.Use(middleware.Cors)
	return r
}