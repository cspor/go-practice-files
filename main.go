package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const pagesFolder = "./pages"
const buildsFolder = "./builds"
const extension = "jsonl"

var waitGroup = sync.WaitGroup{}

type Row struct {
	Id string `json:"id" ,bson:"bsonid"`
}

func main() {
	remakeFolder(pagesFolder)
	remakeFolder(buildsFolder)

	pagesStart := time.Now()

	writePages(100, 1_000)

	waitGroup.Wait()

	fmt.Printf("Creating pages took: %s \n", time.Since(pagesStart))

	// write all files in source directory to destination
	writeStart := time.Now()
	writeFilesInDirToDestination(pagesFolder, buildsFolder, "export_write")
	fmt.Printf("Writing to export took: %s \n", time.Since(writeStart))

	// copy all files in source directory to destination
	copyStart := time.Now()
	copyFilesInDirToDestination(pagesFolder, openFileToAppend(buildsFolder, "export_copy"))
	fmt.Printf("Copying to export took: %s \n", time.Since(copyStart))
}

// writePages Writes rowCount rows to pageCount pages
func writePages(pageCount int, rowCount int) {
	for index := 1; index <= pageCount; index++ {
		waitGroup.Add(1)
		go writeUUIDsToFile(pagesFolder, fmt.Sprint("page_", index), rowCount)
	}
}

// check Checks an error and panics if it's not nil
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// remakeFolder Deletes folder and then remakes it
func remakeFolder(folderName string) {
	err := os.RemoveAll(folderName)
	check(err)

	makeFolder(folderName)
}

// makeFolder Makes a folder
func makeFolder(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	check(err)
}

// generateNewRow Generates a new Row with a UUID as id
func generateNewRow() *Row {
	id, err := uuid.NewRandom()
	check(err)

	return &Row{Id: id.String()}
}

// buildPath Builds the file path from the folderName and fileName
func buildPath(folderName string, fileName string) string {
	return folderName + "/" + fileName + "." + extension
}

// openFileToAppend Opens a file and readies it to be appended to
func openFileToAppend(folderName string, fileName string) *os.File {
	file, err := os.OpenFile(buildPath(folderName, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	check(err)

	return file
}

// openFileToRead Opens a file for reading
func openFileToRead(folderName string, fileName string) *os.File {
	file, err := os.Open(buildPath(folderName, fileName))
	check(err)

	return file
}

// writeUUIDsToFile writes count Rows to the file
func writeUUIDsToFile(folderName string, fileName string, count int) {
	file := openFileToAppend(folderName, fileName)

	bufferedWriter := bufio.NewWriter(file)

	//
	for index := 1; index <= count; index++ {

		//rowJson, err := json.Marshal(generateNewRow())
		//check(err)

		//bytesCount, err := bufferedWriter.Write(rowJson)
		//_ = bytesCount
		//check(err)

		bufferedWriter.WriteString(
			"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
				"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
				"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
				"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
				"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
				"\n",
		)
	}

	e := bufferedWriter.Flush()
	check(e)

	fmt.Printf("finished writing to %s\n", fileName)

	check(file.Close())

	waitGroup.Done()
}

func writeFilesInDirToDestination(sourceDirPath string, destinationDirPath string, destinationFileName string) {
	output := openFileToAppend(destinationDirPath, destinationFileName)

	folder, _ := os.ReadDir(sourceDirPath)

	for _, fileInFolder := range folder {
		file, err := os.Open(sourceDirPath + "/" + fileInFolder.Name())
		check(err)

		data, e := ioutil.ReadAll(file)
		check(e)

		output.Write(data)
		file.Close()
	}
}

func copyFilesInDirToDestination(sourceDirPath string, destination io.Writer) {

	folder, _ := os.ReadDir(sourceDirPath)

	for _, fileInFolder := range folder {
		file, err := os.Open(sourceDirPath + "/" + fileInFolder.Name())
		check(err)

		bytes, e := io.Copy(destination, file)
		_ = bytes
		check(e)
	}
}
