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
// chanWithUrls - channel in what stores Items what has in it all file info
// waitGroup - for marking when process done
// numOfUrls - reference for number in what stored number of all urls that exists in file

func ReadPictureUrls(filePath string, chanWithUrls chan Item, waitGroup *sync.WaitGroup, numOfUrls *int) {
	// reads all file in byte -> convert it to string -> splits it with new line symbol

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
