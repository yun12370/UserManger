package controller

import (
	"html/template"
	"net/http"
)

type HomeController struct {
	Tmpl *template.Template
}

func NewHomeController() *HomeController {
	return &HomeController{
		Tmpl: template.Must(template.ParseFiles("views/index.html")),
	}
}

func (hc *HomeController) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := hc.Tmpl.Execute(w, nil); err != nil {
		http.Error(w, "模板错误", http.StatusInternalServerError)
	}
}
