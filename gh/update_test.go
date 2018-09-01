package gh_test

import (
	"fmt"
	"testing"

	"github.com/graphql-ru/bot/gh"
)

func TestUpdate(t *testing.T) {
	g := gh.New()

	fmt.Printf("%v", g.Versions)
	// g.Update()
	// g.Write()
}
