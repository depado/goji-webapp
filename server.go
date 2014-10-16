package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	gojiweb "github.com/zenazn/goji/web"

	"labix.org/v2/mgo"

	//"github.com/depado/webapp-goji/controllers"
	"github.com/depado/webapp-goji/models"
	"github.com/depado/webapp-goji/system"
)

func GetHello(c gojiweb.C, w http.ResponseWriter, r *http.Request) {
	database := c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
	allEntries, err := models.AllEntries(database)
	if err != nil {
		log.Fatalf("Could not retrieve Entries : %v", err)
		panic(err)
	}
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, allEntries)
}

func main() {
	filename := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()

	var application = &system.Application{}

	// Initialize the application and connect to database
	application.Init(filename)
	application.ConnectToDatabase()

	// Setup Static Files
	static := gojiweb.New()
	static.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(application.Configuration.PublicPath))))
	http.Handle("/static/", static)

	// Setup Middleware
	goji.Use(application.ApplyDatabase)

	// Setup Routes
	goji.Get("/", GetHello)

	// PostHook Declaration
	graceful.PostHook(func() {
		application.Close()
	})
	goji.Serve()
}
