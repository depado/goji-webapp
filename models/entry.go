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

// Insert an entry to the database
func InsertEntry(database *mgo.Database, entry *Entry) error {
	entry.ID = bson.NewObjectId()
	return database.C("entries").Insert(entry)
}

// Find an entry by id
func GetEntryByID(database *mgo.Database, id string) (entry Entry, err error) {
	bid := bson.ObjectIdHex(id)
	err = database.C("entries").FindId(bid).One(&entry)
	return
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

func CountAllEntries(database *mgo.Database) (count int, err error) {
	count, err = database.C("entries").Find(nil).Count()
	return
}
