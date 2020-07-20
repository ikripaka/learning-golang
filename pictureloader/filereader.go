package main

import (
<<<<<<< HEAD
	"io/ioutil"
	"log"
	"strings"
	"sync"
=======
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
>>>>>>> parent of 2002d8b... Program run all groutines in parallel but I don't know how to close channel in imageLoader.go
)

// Reads file with urls and pushes them to the buffered channel
// filePath - path to file that contain urls
<<<<<<< HEAD
// chanWithUrls -
// waitGroup -
// sliceSize
func ReadPictureUrls(filePath string, chanWithUrls chan string, waitGroup *sync.WaitGroup, numOfUrls *int) {
	allFileInByte, err := ioutil.ReadFile(filePath)
=======
func ReadPictureUrls(filePath string, chanWithUrls chan string) {

	file, err := os.Open(filePath)
>>>>>>> parent of 2002d8b... Program run all groutines in parallel but I don't know how to close channel in imageLoader.go
	if err != nil {
		log.Fatal("Problems with reading file")
	}
	urlsSlice := strings.Split(string(allFileInByte), "\n")
	for _, val := range urlsSlice {
		chanWithUrls <- val
	}
	*numOfUrls = len(urlsSlice)
	close(chanWithUrls)
<<<<<<< HEAD
	waitGroup.Done()
=======

}

// Counts file lines for creating buffered channel
// filePath - path to file to count lines in file
func countLines(filePath string) (int, error) {
	reader, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var count int
	const lineBreak = '\n'

	buffer := make([]byte, bufio.MaxScanTokenSize)
	for {
		bufferSize, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return 0, err
		}
		var bufferPosition int
		for {
			i := bytes.IndexByte(buffer[bufferPosition:], lineBreak)
			if i == -1 || bufferSize == bufferPosition {
				break
			}
			bufferPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	return count + 1, nil
>>>>>>> parent of 2002d8b... Program run all groutines in parallel but I don't know how to close channel in imageLoader.go
}
