package system

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

// Clone the database sessions created with the application initialisation
func (application *Application) ApplyDatabase(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session := application.DBSession.Clone()
		defer session.Close()
		c.Env["DBSession"] = session
		c.Env["DBName"] = application.Configuration.Database.Database
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
