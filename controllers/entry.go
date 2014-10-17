package controllers

import (
	"html/template"
	"log"
	"net/http"

	"labix.org/v2/mgo"

	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/models"
)

// Get all the entries and display them
func GetEntries(c web.C, w http.ResponseWriter, r *http.Request) {
	database := c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
	allEntries, err := models.AllEntries(database)
	if err != nil {
		log.Fatalf("Could not retrieve Entries : %v", err)
		panic(err)
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, allEntries)
}

// Get a specific entry
func GetEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	database := c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
	entry, err := models.GetEntryByID(database, c.URLParams["id"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Redirect(w, r, "/entries/", 301)
			return
		} else {
			log.Fatalf("Could not retrieve Entries : %v", err)
			panic(err)
		}
	}
	count, err := models.CountAllEntries(database)
	if err != nil {
		log.Fatalf("Could not count Entries : %v", err)
		panic(err)
	}
	data := struct {
		Entry models.Entry
		Count int
	}{
		entry,
		count,
	}
	t, _ := template.ParseFiles("views/entry.html")
	t.Execute(w, data)
}
