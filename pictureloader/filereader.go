package main

import (
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

// Reads file with urls and pushes them to the buffered channel
// filePath - path to file that contain urls
func ReadPictureUrls(filePath string, chanWithUrls chan string, waitGroup *sync.WaitGroup) {
	allFileInByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Problems with reading file")
	}
	urlsSlice := strings.Split(string(allFileInByte), "\n")
	for _, val := range urlsSlice {
		chanWithUrls <- val
	}
	close(chanWithUrls)
	waitGroup.Done()
}