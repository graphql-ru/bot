package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// List of github repos for release remind
var reposToRemind = []string{
	"graphql/graphiql",
	"graphql/graphql-js",
	// "apollographql/apollo-client",
	"apollographql/apollo-server",
	"apollographql/graphql-tools",
	"apollographql/react-apollo",
}

// Client provides basic features
type Client struct {
	HTTPClient  http.Client
	GithubToken string
	Versions    map[string]string
}

// New creates instance of Releases
func New() Client {
	httpClient := http.Client{Timeout: time.Second * 10}
	githubToken := os.Getenv("GITHUB_API_TOKEN")

	if githubToken == "" {
		log.Printf("[OOPS] GITHUB_API_TOKEN not provided")
	}

	releases := Client{
		HTTPClient:  httpClient,
		GithubToken: githubToken,
		Versions:    map[string]string{},
	}

	releases.Read()

	return releases
}

// Reminder create instance of Reminder
func (r *Client) Reminder() Reminder {
	wg := sync.WaitGroup{}

	reminder := Reminder{
		Updated:  map[string]string{},
		Outdated: map[string]string{},
	}

	routine := func(repo string) {
		defer wg.Done()

		prev := r.Versions[repo]
		next := r.VersionOf(repo)

		if prev != "" && next != "" && prev != next {
			reminder.AddPackage(repo, prev, next)
		}

		r.Versions[repo] = next
	}

	for _, repo := range reposToRemind {
		wg.Add(1)
		go routine(repo)
	}

	wg.Wait()
	r.Write()

	return reminder
}

// ReminderTicker starts reminder tick loop
func (r *Client) ReminderTicker(duration time.Duration, onTick func(msg string)) *time.Ticker {
	ticker := time.NewTicker(duration)

	go func() {
		for range ticker.C {
			reminder := r.Reminder()

			if reminder.HasUpdates() {
				log.Printf("[REMINDER] %d package released", len(reminder.Updated))
				onTick(reminder.Message())
			}
		}
	}()

	return ticker
}

// Fetch just wrapper for http.Get
func (r *Client) Fetch(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", r.GithubToken))

	resp, err := r.HTTPClient.Do(req)

	return resp, err
}

// VersionOf fetch and returns version of given repo
func (r *Client) VersionOf(repo string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	result := map[string]interface{}{}

	resp, err := r.Fetch(url)

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

// Read from versions.json
func (r *Client) Read() {
	file, err := ioutil.ReadFile("/tmp/versions.json")

	if err != nil {
		log.Printf("[OOPS] Can not read versions.json")
		return
	}

	versions := map[string]string{}
	json.Unmarshal(file, &versions)

	for key, value := range versions {
		r.Versions[key] = value
	}
}

// Write versions into versions.json
func (r *Client) Write() {
	versions, err := json.Marshal(r.Versions)

	if err != nil {
		log.Printf("[OOPS] Can not write versions to versions.json")
		return
	}

	ioutil.WriteFile("/tmp/versions.json", versions, 0644)
}
