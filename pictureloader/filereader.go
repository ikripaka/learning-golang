package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

// Reads file with urls and pushes them to the buffered channel
// imageFolderPath - path to file that contain urls
// chanWithUrls -
// waitGroup -
// sliceSize
func ReadPictureUrls(filePath string, chanWithUrls chan Item, waitGroup *sync.WaitGroup, numOfUrls *int) {
	allFileInByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Problems with reading file")
	}
	sliceOfUrls := strings.Split(string(allFileInByte), "\n")
	for _, val := range sliceOfUrls {
		chanWithUrls <- Item{url: val}
	}
	*numOfUrls = len(sliceOfUrls)
	close(chanWithUrls)
	fmt.Println(*numOfUrls)
	waitGroup.Done()
}
