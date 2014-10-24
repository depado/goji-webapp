package controllers

import (
	"html/template"
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/helpers"
)

func GetIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	context := helpers.GenerateWrappedBaseContext("home", c)
	t, _ := template.ParseFiles("views/base.html", "views/menu.html", "views/index.html")
	t.ExecuteTemplate(w, "base", context)
}
