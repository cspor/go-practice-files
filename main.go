package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cspor/go-practice-files/config"
	"github.com/cspor/go-practice-files/errorHandler"
	"github.com/cspor/go-practice-files/filesystem"
	"github.com/cspor/go-practice-files/row"
	"github.com/cspor/go-practice-files/timer"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

func main() {
	filesystem.RemakeFolder(config.PagesFolder)
	filesystem.RemakeFolder(config.BuildsFolder)

	pagesStart := time.Now()

	writePages(config.PageCount, config.RowCount)

	waitGroup.Wait()

	timer.Took("Creating pages", pagesStart)

	// write all files in source directory to destination
	writeStart := time.Now()
	filesystem.WriteFilesInDirToDestination(config.PagesFolder, config.BuildsFolder, "export_write")
	timer.Took("Writing to export", writeStart)

	// copy all files in source directory to destination
	copyStart := time.Now()
	filesystem.CopyFilesInDirToDestination(config.PagesFolder, filesystem.OpenFileToAppend(config.BuildsFolder, "export_copy"))
	timer.Took("Copying to export", copyStart)
}

// writePages Writes rowCount rows to pageCount pages
func writePages(pageCount int, rowCount int) {
	for index := 1; index <= pageCount; index++ {
		waitGroup.Add(1)
		go writeUUIDsToFile(config.PagesFolder, fmt.Sprint("page_", index), rowCount)
	}
}

// writeUUIDsToFile writes count Rows to the file
func writeUUIDsToFile(folderName string, fileName string, count int) {
	file := filesystem.OpenFileToAppend(folderName, fileName)

	bufferedWriter := bufio.NewWriter(file)

	// Write new rows to the file
	for index := 1; index <= count; index++ {
		rowJson, err := json.Marshal(row.NewRow())
		errorHandler.Check(err)

		bytesCount, err := bufferedWriter.Write(rowJson)
		_ = bytesCount
		errorHandler.Check(err)

		bufferedWriter.WriteString("\n")
	}

	e := bufferedWriter.Flush()
	errorHandler.Check(e)

	fmt.Printf("finished writing to %s\n", fileName)

	errorHandler.Check(file.Close())

	waitGroup.Done()
}
