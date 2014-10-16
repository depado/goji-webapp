package models

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	Entries []Entry
	Entry   struct {
		ID      bson.ObjectId `bson:"_id,omitempty"`
		Title   string        `bson:"title"`
		Content string        `bson:"content"`
		Posted  time.Time     `bson:"posted"`
	}
)

func InsertEntry(database *mgo.Database, entry *Entry) error {
	entry.ID = bson.NewObjectId()
	return database.C("entries").Insert(entry)
}

// Retrieves all the entries
func AllEntries(database *mgo.Database) (entries Entries, err error) {
	err = database.C("entries").Find(nil).All(&entries)
	return
}

// Retrive all the entries sorted by date.
func AllEntriesByDate(database *mgo.Database) (entries Entries, err error) {
	err = database.C("entries").Find(nil).Sort("-posted").All(&entries)
	return
}
