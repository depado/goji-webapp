package helpers

import (
	"github.com/zenazn/goji/web"
	"labix.org/v2/mgo"
)

func GetDatabaseFromEnv(c web.C) *mgo.Database {
	return c.Env["DBSession"].(*mgo.Session).DB(c.Env["DBName"].(string))
}
