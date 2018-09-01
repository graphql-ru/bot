package gh

import (
	"fmt"
	"sync"
)

// List of github repos for release remind
var repos = []string{
	"graphql/graphiql",
	"graphql/graphql-js",
	// "apollographql/apollo-client",
	"apollographql/apollo-server",
	"apollographql/graphql-tools",
	"apollographql/react-apollo",
}

// Update checks for new releases
func (r *Releases) Update() {
	wg := sync.WaitGroup{}

	routine := func(repo string) {
		defer wg.Done()
		fmt.Println(r.VersionOf(repo))
	}

	for _, repo := range repos {
		wg.Add(1)
		go routine(repo)
	}

	wg.Wait()
}
