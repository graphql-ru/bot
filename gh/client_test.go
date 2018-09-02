package gh_test

import (
	"fmt"
	"testing"

	"github.com/graphql-ru/bot/gh"
)

func TestUpdate(t *testing.T) {
	g := gh.New()

	g.Reminder()
	fmt.Printf("%v\n", g.Versions)
}
