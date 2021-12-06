package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cspor/go-practice-files/errorHandler"
	"github.com/cspor/go-practice-files/filesystem"
	"github.com/cspor/go-practice-files/row"
	"github.com/cspor/go-practice-files/timer"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const pagesFolder = "./results/pages"
const buildsFolder = "./results/builds"
const extension = "jsonl"

const pageCount = 100
const rowCount = 5_000

var waitGroup = sync.WaitGroup{}

func main() {
	filesystem.RemakeFolder(pagesFolder)
	filesystem.RemakeFolder(buildsFolder)

	pagesStart := time.Now()

	writePages(pageCount, rowCount)

	waitGroup.Wait()

	timer.Took("Creating pages", pagesStart)

	// write all files in source directory to destination
	writeStart := time.Now()
	writeFilesInDirToDestination(pagesFolder, buildsFolder, "export_write")
	timer.Took("Writing to export", writeStart)

	// copy all files in source directory to destination
	copyStart := time.Now()
	copyFilesInDirToDestination(pagesFolder, openFileToAppend(buildsFolder, "export_copy"))
	timer.Took("Copying to export", copyStart)
}

// writePages Writes rowCount rows to pageCount pages
func writePages(pageCount int, rowCount int) {
	for index := 1; index <= pageCount; index++ {
		waitGroup.Add(1)
		go writeUUIDsToFile(pagesFolder, fmt.Sprint("page_", index), rowCount)
	}
}

// buildPath Builds the file path from the folderName and fileName
func buildPath(folderName string, fileName string) string {
	return folderName + "/" + fileName + "." + extension
}

// openFileToAppend Opens a file and readies it to be appended to
func openFileToAppend(folderName string, fileName string) *os.File {
	file, err := os.OpenFile(buildPath(folderName, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	errorHandler.Check(err)

	return file
}

// openFileToRead Opens a file for reading
func openFileToRead(folderName string, fileName string) *os.File {
	file, err := os.Open(buildPath(folderName, fileName))
	errorHandler.Check(err)

	return file
}

// writeUUIDsToFile writes count Rows to the file
func writeUUIDsToFile(folderName string, fileName string, count int) {
	file := openFileToAppend(folderName, fileName)

	bufferedWriter := bufio.NewWriter(file)

	//
	for index := 1; index <= count; index++ {

		rowJson, err := json.Marshal(row.NewRow())
		errorHandler.Check(err)

		bytesCount, err := bufferedWriter.Write(rowJson)
		_ = bytesCount
		errorHandler.Check(err)

		bufferedWriter.WriteString("\n")

		//bufferedWriter.WriteString(
		//	"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
		//		"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
		//		"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
		//		"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
		//		"Urna blandit amet arcu ante ridiculus convallis facilisi mollis non condimentum vestibulum maecenas sodales eu sagittis porta. At mi ac elit nam sed imperdiet sagittis a taciti consequat malesuada senectus nec a a adipiscing pulvinar amet lacinia viverra pretium torquent. Sed elit sociis praesent senectus id scelerisque per proin ligula elit himenaeos sagittis eleifend aenean vehicula. Iaculis molestie et vestibulum dignissim parturient praesent risus sed suspendisse cum arcu urna urna nec vestibulum primis donec blandit. Ac ut quisque aptent nisi at scelerisque a aenean sed varius ullamcorper natoque ut euismod vehicula. Ad sem tempus curae a parturient congue tristique adipiscing fringilla massa consectetur suspendisse sed imperdiet primis nam luctus vitae eu varius ultricies integer non massa a a mus. Scelerisque ad consectetur nam viverra sem cras condimentum egestas bibendum maecenas a proin orci libero tortor nam ad dis.Per porta ac condimentum et ad placerat parturient sodales." +
		//		"\n",
		//)
	}

	e := bufferedWriter.Flush()
	errorHandler.Check(e)

	fmt.Printf("finished writing to %s\n", fileName)

	errorHandler.Check(file.Close())

	waitGroup.Done()
}

func writeFilesInDirToDestination(sourceDirPath string, destinationDirPath string, destinationFileName string) {
	output := openFileToAppend(destinationDirPath, destinationFileName)

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

func copyFilesInDirToDestination(sourceDirPath string, destination io.Writer) {

	folder, _ := os.ReadDir(sourceDirPath)

	for _, fileInFolder := range folder {
		file, err := os.Open(sourceDirPath + "/" + fileInFolder.Name())
		errorHandler.Check(err)

		bytes, e := io.Copy(destination, file)
		_ = bytes
		errorHandler.Check(e)
	}
}
