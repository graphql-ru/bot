package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Current versions for repos
var versions = map[string]string{}

// List of github repos for release remind
var repos = []string{
	"graphql/graphiql",
	"graphql/graphql-js",
	// "apollographql/apollo-client",
	"apollographql/apollo-server",
	"apollographql/graphql-tools",
	"apollographql/react-apollo",
}

// Ensure checks for new releases
func (r *Releases) Ensure() {
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

// Get just wrapper for http.Get
func (r *Releases) Get(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", r.githubToken))

	resp, err := r.httpClient.Do(req)

	return resp, err
}

// VersionOf fetch and returns version of given repo
func (r *Releases) VersionOf(repo string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	result := map[string]interface{}{}

	resp, err := r.Get(url)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		return ""
	}

	json.Unmarshal(body, &result)

	if result["tag_name"] == nil {
		return ""
	}

	return result["tag_name"].(string)
}
