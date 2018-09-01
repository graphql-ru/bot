package gh

import (
	"fmt"
	"net/http"
)

// Fetch just wrapper for http.Get
func (r *Releases) Fetch(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", r.GithubToken))

	resp, err := r.HTTPClient.Do(req)

	return resp, err
}
