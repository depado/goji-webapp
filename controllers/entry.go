package controllers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"labix.org/v2/mgo"

	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/helpers"
	"github.com/depado/webapp-goji/models"
)

// Get all the entries and display them
func GetEntries(c web.C, w http.ResponseWriter, r *http.Request) {
	database := helpers.GetDatabaseFromEnv(c)
	allEntries, err := models.AllEntriesByDate(database)
	if err != nil {
		log.Fatalf("Could not retrieve Entries : %v", err)
		panic(err)
	}
	context := helpers.GenerateEntriesContext("entries", allEntries)
	t, _ := template.ParseFiles("views/base.html", "views/menu.html", "views/entries.html")
	t.ExecuteTemplate(w, "base", context)
}

// Get a specific entry
func GetEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	database := helpers.GetDatabaseFromEnv(c)
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
	context := helpers.GenerateEntryContext("entry", entry, count)
	t, _ := template.ParseFiles("views/base.html", "views/menu.html", "views/entry.html")
	t.ExecuteTemplate(w, "base", context)
}

// Get the New Entry page
func GetNewEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	context := helpers.GenerateWrappedBaseContext("new", c)
	t, _ := template.ParseFiles("views/base.html", "views/menu.html", "views/new_entry.html")
	t.ExecuteTemplate(w, "base", context)
}

// Retrieves data from the New Entry page
func PostNewEntry(c web.C, w http.ResponseWriter, r *http.Request) {
	database := helpers.GetDatabaseFromEnv(c)
	author, title := r.FormValue("author"), r.FormValue("title")
	short, content := r.FormValue("short"), r.FormValue("content")
	if helpers.EmptyStrings(author, title, short, content) {
		http.Redirect(w, r, "/entries/new/", 301)
		return
	}
	entry := &models.Entry{
		Author:  author,
		Title:   title,
		Short:   short,
		Content: content,
		Posted:  time.Now(),
	}
	if err := models.InsertEntry(database, entry); err != nil {
		log.Fatalf("Could not add new entry :%v", err)
		panic(err)
	}
	http.Redirect(w, r, "/entries/", 301)
}
