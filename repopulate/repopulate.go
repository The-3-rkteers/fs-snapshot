package repopulate

import (
	"encoding/json"
	"io/ioutil"

	"github.com/The-3-rkteers/fs-snapshot/dirtree"
)

// LoadSnapshot loads the snapshot json file into memory
func LoadSnapshot(fileName string) (interface{}, error) {
	jsonString, err := ioutil.ReadFile(fileName)
	jsonBytes := []byte(jsonString)
	var root dirtree.Directory
	err = json.Unmarshal(jsonBytes, root)
	return jsonBytes, err
}

func CreateDirectory() {

}
