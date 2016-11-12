package imagemake

import (
	"os"
	"path/filepath"
)

var root directory

// AddToTree is called on each fs object to add to the directory tree
func AddToTree(path string, f os.FileInfo, err error) error {
	insert(&root, f)
	return nil
}

// BuildTree calls AddToTree on each fs object
func BuildTree() error {
	return filepath.Walk("/", AddToTree)
}
