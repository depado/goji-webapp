package helpers

import (
	"github.com/depado/webapp-goji/models"
	"github.com/microcosm-cc/bluemonday"
	"html/template"
)

func SanitizeEntry(entry models.Entry) (sanitizedEntry models.Entry) {
	sanitizedEntry = entry
	sanitizer := bluemonday.UGCPolicy()
	sanitizedEntry.Content = template.HTML(sanitizer.Sanitize(string(sanitizedEntry.Content)))
	return sanitizedEntry
}

func SanitizeEntries(entries models.Entries) (sanitizedEntries models.Entries) {
	sanitizedEntries = entries
	sanitizer := bluemonday.UGCPolicy()
	for _, entry := range sanitizedEntries {
		entry.Content = template.HTML(sanitizer.Sanitize(string(entry.Content)))
	}
	return sanitizedEntries
}
