package controller

import (
	"net/http"
	"html/template"
	"../model"
)

type pageIndex struct {
	Title string
	*model.DockerInfo
}

func Index(w http.ResponseWriter, r *http.Request){

	modelDockerInfo := &model.DockerInfo{}
	modelDockerInfo.GetGeneralInfo()
	modelDockerInfo.GetDetailedInfo()
	page := &pageIndex{"Main", modelDockerInfo}

	w.Header().Set("Content-type", "text/html")
	htmlFiles := []string{"./Template/page.html", "./Template/menu.html", "./Template/navbar-top.html", "./Template/page-index.html"}
	t, _ := template.ParseFiles(htmlFiles...)
	t.Execute(w, page)
	return
}