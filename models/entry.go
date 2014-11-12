package models

import (
	"time"

	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	Entries []Entry
	Entry   struct {
		ID      bson.ObjectId `bson:"_id,omitempty"`
		Author  string        `bson:"author"`
		Title   string        `bson:"title"`
		Short   string        `bson:"short"`
		Content template.HTML `bson:"content"`
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

// Retrieve all the entries sorted by date.
func AllEntriesByDate(database *mgo.Database) (entries Entries, err error) {
	err = database.C("entries").Find(nil).Sort("-posted").All(&entries)
	return
}

// Counts all the entries.
func CountAllEntries(database *mgo.Database) (count int, err error) {
	count, err = database.C("entries").Find(nil).Count()
	return
}
