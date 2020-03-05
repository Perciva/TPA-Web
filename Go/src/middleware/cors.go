
package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)
var Apikey string
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars:= mux.Vars(r)
		Apikey = vars["key"]
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		next.ServeHTTP(w, r)
	})
}
