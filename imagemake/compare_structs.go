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
		// Create directory struct, Add to subTrees
	} else {
		// Create fileInfo struct, Add to entries
	}
	return nil
}
