package gh

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Releases provides basic features
type Releases struct {
	HTTPClient  http.Client
	GithubToken string
	Versions    map[string]string
}

// New creates instance of Releases
func New() Releases {
	httpClient := http.Client{Timeout: time.Second * 10}
	githubToken := os.Getenv("GITHUB_API_TOKEN")

	if githubToken == "" {
		log.Printf("[OOPS] GITHUB_API_TOKEN not provided")
	}

	releases := Releases{
		HTTPClient:  httpClient,
		GithubToken: githubToken,
		Versions:    map[string]string{},
	}

	releases.Read()

	return releases
}
