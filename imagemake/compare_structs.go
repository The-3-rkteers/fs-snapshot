package imagemake

import (
	"os"
	"strings"
)

type directory struct {
	directoryPath  string
	entries        []*fileInfo
	subDirectories []*directory
}

type fileInfo struct {
	fullPath       string
	permissionMode os.FileMode
	lastAccessTime int
	contentChanged bool
}

func insert(root *directory, fullPath string, f os.FileInfo) error {

	splitPath := strings.Split(fullPath, "/")

	// If we're in the correct directory
	if root.directoryPath == splitPath[len(splitPath)-2] {
		if f.IsDir() {
			var newDir = new(directory)
			newDir.directoryPath = f.Name()
			root.subDirectories = append(root.subDirectories, newDir)
		} else {
			var newEntry = new(fileInfo)
			newEntry.fullPath = fullPath
			newEntry.contentChanged = false
			newEntry.permissionMode = f.Mode()
			newEntry.lastAccessTime = 0
			root.entries = append(root.entries, newEntry)
		}
	} else {
		// Step to the correct directory, recursively call insert
		for _, subDir := range root.subDirectories {
			if subDir.directoryPath == splitPath[1] {
				err := insert(subDir, strings.Join(splitPath[1:], ""), f)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
