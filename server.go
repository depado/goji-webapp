package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	gojiweb "github.com/zenazn/goji/web"
	"labix.org/v2/mgo"

	//"github.com/depado/webapp-goji/controllers"
	"github.com/depado/webapp-goji/models"
	"github.com/depado/webapp-goji/system"
)

func GetHello(c gojiweb.C, w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	dbName := c.Env["DBName"]
	log.Println("DB Name :", dbName)
	user := models.User{
		Email:     "depado@depado.eu",
		Username:  "Depado",
		Timestamp: time.Now(),
	}
	user.HashPassword("pony")
	database := c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
	if err := models.InsertUser(database, &user); err != nil {
		log.Fatal("Could not insert")
		panic(err)
	}

	t.Execute(w, nil)
}

func main() {
	filename := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()

	var application = &system.Application{}

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
