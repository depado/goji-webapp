package controllers

import (
	"html/template"
	"log"
	"net/http"

	"labix.org/v2/mgo"

	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/models"
)

func GetIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	database := c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
	allEntries, err := models.AllEntries(database)
	if err != nil {
		log.Fatalf("Could not retrieve Entries : %v", err)
		panic(err)
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, allEntries)
}
