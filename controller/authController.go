package controller

import (
	"net/http"
	"html/template"

	"../model"
)

type page struct {
	Title string
}

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		err := model.CheckAuth(w, r)
		if err != nil{
			http.Redirect(w, r, "/login", 303)
			return
		}
		http.Redirect(w, r, "/index", 303)
		return
	}
	w.Header().Set("Content-type", "text/html")
	htmlFiles := []string{"./Template/page.html", "./Template/page-login.html"}
	t, _ := template.ParseFiles(htmlFiles...)
	t.Execute(w, &page{Title: "Login",})
	return
}

func IsAdmin(w http.ResponseWriter, r *http.Request) error{
	return model.CheckIsAdmin(w, r)
}

func Logout(w http.ResponseWriter, r *http.Request){
	model.Logout(w, r)
	http.Redirect(w, r, "/login", 303)
}
