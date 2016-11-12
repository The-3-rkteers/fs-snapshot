package imagemake

import (
	"fmt"
	"os"
	"path/filepath"
)

var root directory

func addToTree(path string, f os.FileInfo, err error) error {
	insert(&root, f)
	return nil
}

func buildTree() {
	err := filepath.Walk("/", addToTree)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
