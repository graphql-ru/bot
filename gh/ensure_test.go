package gh_test

import (
	"testing"

	"github.com/graphql-ru/bot/gh"
)

func TestEnsure(t *testing.T) {
	g := gh.New()
	g.Ensure()
}
