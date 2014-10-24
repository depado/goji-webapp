package helpers

import (
	"log"

	"github.com/zenazn/goji/web"

	"github.com/depado/webapp-goji/models"
)

type BaseContext struct {
	Active string
	Count  int
}

type BaseContextWrapper struct {
	Base BaseContext
}

type EntryContext struct {
	Entry models.Entry
	Base  BaseContext
}

type EntriesContext struct {
	Entries models.Entries
	Base    BaseContext
}

func GenerateEntryContext(active string, entry models.Entry, count int) EntryContext {
	return EntryContext{
		Entry: entry,
		Base:  GenerateBaseContext(active, count),
	}
}

func GenerateEntriesContext(active string, entries models.Entries) EntriesContext {
	return EntriesContext{
		Entries: entries,
		Base:    GenerateBaseContext(active, len(entries)),
	}
}

func GenerateBaseContext(active string, count int) BaseContext {
	return BaseContext{
		Active: active,
		Count:  count,
	}
}

// A basic context, containing only the necessary and doesn't need a database instance
func GenerateWrappedBaseContext(active string, c web.C) BaseContextWrapper {
	count, err := models.CountAllEntries(GetDatabaseFromEnv(c))
	if err != nil {
		log.Fatalf("Could not count Entries : %v", err)
		panic(err)
	}
	return BaseContextWrapper{
		Base: GenerateBaseContext(active, count),
	}
}
