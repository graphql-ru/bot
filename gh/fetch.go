package gh

import (
	"fmt"
	"net/http"
)

// Get just wrapper for http.Get
func (r *Releases) Get(url string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", r.githubToken))

	resp, err := r.httpClient.Do(req)

	return resp, err
}
