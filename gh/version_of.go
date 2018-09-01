package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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
