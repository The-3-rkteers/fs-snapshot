package imagemake

import "os"

type directory struct {
	directoryName string
	entries       []*fileInfo
	subTrees      []*directory
}

type fileInfo struct {
	fullPath       string
	permissionMode int
	lastAccessTime int
	contentChanged bool
}

func insert(root *directory, f os.FileInfo) error {
	if f.IsDir() {
		// Create directory struct
		var newDir = new(directory)

		// Traverse to where dir should go
		// Insert into subtrees slice
	} else {
		// Create fileInfo struct
		// Traverse to where entry should go
		// Insert into entries slice
	}
	return nil
}
