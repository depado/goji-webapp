package main

import (
	"flag"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/controllers"
	"github.com/depado/webapp-goji/system"
)

func main() {
	filename := flag.String("config", "config.json", "Path to configuration file")
	flag.Parse()

	var application = &system.Application{}

	// Initialize the application and connect to database
	application.Init(filename)
	application.ConnectToDatabase()

	// Setup Static Files
	static := web.New()
	static.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir(application.Configuration.PublicPath))))
	http.Handle("/static/", static)

	// Setup Middleware
	goji.Use(application.ApplyDatabase)

	// Setup Routes
	goji.Get("/", controllers.GetIndex)
	goji.Get("/entries/", controllers.GetEntries)
	goji.Get("/entries", http.RedirectHandler("/entries/", 301))
	goji.Get("/entries/:id/", controllers.GetEntry)

	// PostHook Declaration
	graceful.PostHook(func() {
		application.Close()
	})
	goji.Serve()
}
