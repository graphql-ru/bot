package gh

import (
	"fmt"
	"strings"
)

// Reminder collect updated packages and sends to @graphql_ru chat
type Reminder struct {
	Packages map[string]string
}

// AddPackage just add package which should be remined
func (r *Reminder) AddPackage(name, version string) {
	r.Packages[name] = version
}

// Message just render message text
func (r *Reminder) Message() string {
	var lines []string

	for pkg, version := range r.Packages {
		tagURL := fmt.Sprintf("https://github.com/%s/releases/tag/%s", pkg, version)

		lines = append(lines, fmt.Sprintf(
			"[%s](%s) **%s** released 🎉",
			pkg,
			tagURL,
			version,
		))
	}

	return strings.Join(lines, "\n")
}

func (r *Reminder) hasUpdates() bool {
	return len(r.Packages) != 0
}
