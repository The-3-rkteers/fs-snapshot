package imagemake

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var root directory

// AddToTree is called on each fs object to add to the directory tree
func AddToTree(path string, f os.FileInfo, err error) error {
	insert(&root, path, f)
	return nil
}

// BuildTree calls AddToTree on each fs object
func BuildTree() error {
	err := filepath.Walk("./", AddToTree)
	if err != nil {
		return err
	}

	for _, i := range root.Entries {
		fmt.Println(*i)
	}

	b, _ := json.MarshalIndent(&root.Entries, "", "    ")
	fmt.Println(string(b))
	return nil
}
