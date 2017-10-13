package middleware

import (
	"net/http"
	"../controller"
)


func AuthMiddleware(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request){
	err := controller.IsAdmin(w, r)
	if err != nil{
		http.Redirect(w, r, "/login", 303)
	}
	return w, r
}