package gh

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Write versions into versions.json
func (r *Releases) Write() {
	versions, err := json.Marshal(r.Versions)

	if err != nil {
		log.Printf("[OOPS] Can not write versions to versions.json")
		return
	}

	ioutil.WriteFile("/tmp/versions.json", versions, 0644)
}
