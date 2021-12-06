package filesystem

import (
	"github.com/cspor/go-practice-files/config"
	"github.com/cspor/go-practice-files/errorHandler"
	"io"
	"io/ioutil"
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

// BuildPath Builds the file path from the folderName and fileName
func BuildPath(folderName string, fileName string) string {
	return folderName + "/" + fileName + "." + config.Extension
}

// OpenFileToAppend Opens a file and readies it to be appended to
func OpenFileToAppend(folderName string, fileName string) *os.File {
	file, err := os.OpenFile(BuildPath(folderName, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	errorHandler.Check(err)

	return file
}

// OpenFileToRead Opens a file for reading
func OpenFileToRead(folderName string, fileName string) *os.File {
	file, err := os.Open(BuildPath(folderName, fileName))
	errorHandler.Check(err)

	return file
}

// WriteFilesInDirToDestination Stitches files from source directory to destination filename by
// writing them one by one
func WriteFilesInDirToDestination(sourceDirPath string, destinationDirPath string, destinationFileName string) {
	output := OpenFileToAppend(destinationDirPath, destinationFileName)

	folder, _ := os.ReadDir(sourceDirPath)

	for _, fileInFolder := range folder {
		file, err := os.Open(sourceDirPath + "/" + fileInFolder.Name())
		errorHandler.Check(err)

		data, e := ioutil.ReadAll(file)
		errorHandler.Check(e)

		output.Write(data)
		file.Close()
	}
}

// CopyFilesInDirToDestination Stitches files from source directory to destination filename by
// copying them on top of the destination writer
func CopyFilesInDirToDestination(sourceDirPath string, destination io.Writer) {

	folder, _ := os.ReadDir(sourceDirPath)

	for _, fileInFolder := range folder {
		file, err := os.Open(sourceDirPath + "/" + fileInFolder.Name())
		errorHandler.Check(err)

		bytes, e := io.Copy(destination, file)
		_ = bytes
		errorHandler.Check(e)
	}
}
