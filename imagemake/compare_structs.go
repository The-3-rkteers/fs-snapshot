package imagemake

import (
	"os"
	"strings"
)

type directory struct {
	DirectoryPath  string
	Entries        []*fileInfo
	SubDirectories []*directory
}

type fileInfo struct {
	FullPath         string
	PermissionNumber os.FileMode
	LastAccessTime   int
	ContentChanged   bool
}

func insert(root *directory, FullPath string, f os.FileInfo) error {

	splitPath := strings.Split(FullPath, "/")

	// If we're in the correct directory
	if len(splitPath) == 1 || len(splitPath) == 2 {
		if f.IsDir() {
			var newDir = new(directory)
			newDir.DirectoryPath = f.Name()
			root.SubDirectories = append(root.SubDirectories, newDir)
		} else {
			var newEntry = new(fileInfo)
			newEntry.FullPath = FullPath
			newEntry.ContentChanged = false
			newEntry.PermissionNumber = f.Mode()
			newEntry.LastAccessTime = 0
			root.Entries = append(root.Entries, newEntry)
		}
	} else {
		// Step to the correct directory, recursively call insert
		for _, subDir := range root.SubDirectories {
			if subDir.DirectoryPath == splitPath[1] {
				err := insert(subDir, strings.Join(splitPath[1:], "/"), f)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
