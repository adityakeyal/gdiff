package main

import (
	"os"
	"path/filepath"
)

func identifyFiles(srca string, fh filesHolder) []string {
	fileNamesA := make([]string, 10)

	filepath.Walk(srca, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if fh.isValidExtn(path) {
				path = path[len(srca):]
				fileNamesA = append(fileNamesA, path)
			}
		}
		return nil
	})
	return fileNamesA
}
