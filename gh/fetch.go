package gh

import (
	"fmt"
	"net/http"
)

// Fetch just wrapper for http.Get
func (r *Releases) Fetch(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", r.githubToken))

	resp, err := r.httpClient.Do(req)

	return resp, err
}
