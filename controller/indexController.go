package controller

import (
	"net/http"
	"html/template"
	"../model"
)

type pageIndex struct {
	Title string
	*model.DockerInfo
	// model.Info
}

func Index(w http.ResponseWriter, r *http.Request){
	var info model.Info
	info = &model.DockerInfo{}
	info.GetGeneralInfo()

	page := &pageIndex{"Main", &model.GeneralInfo}

	w.Header().Set("Content-type", "text/html")
	htmlFiles := []string{"./Template/page.html", "./Template/menu.html", "./Template/navbar-top.html", "./Template/page-index.html"}
	t, _ := template.ParseFiles(htmlFiles...)
	t.Execute(w, page)
	return
}