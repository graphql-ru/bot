package gh

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Releases provides basic features
type Releases struct {
	httpClient  http.Client
	githubToken string
	versions    map[string]string
}

// New creates instance of Releases
func New() Releases {
	httpClient := http.Client{Timeout: time.Second * 10}
	githubToken := os.Getenv("GITHUB_API_TOKEN")

	if githubToken == "" {
		log.Printf("[ERROR] GITHUB_API_TOKEN not provided")
	}

	return Releases{
		httpClient:  httpClient,
		githubToken: githubToken,
		versions:    map[string]string{},
	}
}
