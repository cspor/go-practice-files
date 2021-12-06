package filesystem

import (
	"github.com/cspor/go-practice-files/errorHandler"
	"os"
)

// RemakeFolder Deletes folder and then remakes it
func RemakeFolder(folderName string) {
	err := os.RemoveAll(folderName)
	errorHandler.Check(err)

	MakeFolder(folderName)
}

// MakeFolder Makes a folder
func MakeFolder(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	errorHandler.Check(err)
}
