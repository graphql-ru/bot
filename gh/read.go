package gh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Read from versions.json
func (r *Releases) Read() {
	file, err := ioutil.ReadFile("/tmp/versions.json")

	if err != nil {
		fmt.Printf("[OOPS] Can not read versions.json")
		return
	}

	versions := map[string]string{}
	json.Unmarshal(file, &versions)

	for key, value := range versions {
		r.versions[key] = value
	}
}
