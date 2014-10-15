package system

import (
	"encoding/gob"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Application struct {
	Configuration *Configuration
	DBSession     *mgo.Session
}

func (application *Application) Init(filename *string) {
	gob.Register(bson.ObjectId(""))

	application.Configuration = &Configuration{}
	err := application.Configuration.Load(*filename)

	if err != nil {
		log.Fatalf("Can't read configuration file: %s", err)
		panic(err)
	}
}

func (application *Application) ConnectToDatabase() {
	var err error
	application.DBSession, err = mgo.Dial(application.Configuration.Database.Hosts)

	if err != nil {
		log.Fatalf("Can't connect to the database: %v", err)
		panic(err)
	}
}

func (application *Application) Close() {
	log.Println("Bye.")
	application.DBSession.Close()
}
