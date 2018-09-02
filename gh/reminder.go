package gh

import (
	"fmt"
	"strings"
)

// Reminder collect updated packages and sends to @graphql_ru chat
type Reminder struct {
	Outdated map[string]string
	Updated  map[string]string
}

// AddPackage just add package which should be remined
func (r *Reminder) AddPackage(name, prev, next string) {
	r.Outdated[name] = prev
	r.Updated[name] = next
}

// Message just render message text
func (r *Reminder) Message() string {
	var lines []string

	for pkg, version := range r.Updated {
		outdated := r.Outdated[pkg]
		tagURL := fmt.Sprintf("https://github.com/%s/releases/tag/%s", pkg, version)

		lines = append(lines, fmt.Sprintf(
			"[%s](%s) %s → *%s* 🎉",
			pkg,
			tagURL,
			outdated,
			version,
		))
	}

	return strings.Join(lines, "\n")
}

// HasUpdates retirn true if Updated has at least one key
func (r *Reminder) HasUpdates() bool {
	return len(r.Updated) != 0
}
