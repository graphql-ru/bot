package gh

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Read from versions.json
func (r *Releases) Read() {
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
