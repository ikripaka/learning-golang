package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
)

// Program reads file with urls
// then download these images
// then scale them and save to the directory
// You can write directly in arguments path to file and folder
// but you also can recomment first 4 strokes with code and

// 1 argument - path to the file with urls
// 2 argument - path to the folder where downloaded images would be stored

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

const MAXDOWNLOADPROCESSES = 15

// this struct describes picture info in order not to pass many values in functions
type Item struct {
	url             string
	filename        string
	avatarFilename  string
	imageFolderPath string

	errInDownload error
	errInResizing error
}

func main() {
	numCPU := runtime.NumCPU()
	var numOfUrls, counter int
	var waitGroup sync.WaitGroup

	pictureUrlsChan := make(chan Item)
	downloadedImagesFilenames := make(chan Item)
	resizedImageChan := make(chan Item)

	// get variables
	args := os.Args[1:]
	urlFilePath := args[0]
	folderPath := args[1]

	if _, err := IsPathsCorrect(urlFilePath, folderPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read urls from file..")

	waitGroup.Add(1)
	// read picture urls and push data to pictureUrlsChan
	go ReadPictureUrls(urlFilePath, pictureUrlsChan, &waitGroup, &numOfUrls)

	waitGroup.Add(MAXDOWNLOADPROCESSES)

	fmt.Println("Download images..")

	// load pictures from internet and push data to downloadedImagesFilenames
	for i := 0; i < MAXDOWNLOADPROCESSES; i++ {
		go LoadPictures(folderPath, pictureUrlsChan, downloadedImagesFilenames, &waitGroup)
	}

	waitGroup.Add(numCPU)

	// scale images and push data to resizedImageChan
	fmt.Println("Scale images..")
	for i := 0; i < numCPU; i++ {
		go MakeAvatars(downloadedImagesFilenames, resizedImageChan, &waitGroup, &counter, numOfUrls)
	}

	// show program results
	for counter = 0; counter < numOfUrls; counter++ {
		val, _ := <-resizedImageChan
		if val.errInDownload != nil && val.errInResizing != nil {

		} else if val.errInDownload == nil && val.errInResizing != nil {
			fmt.Println("Failure in download ", val.filename, val.errInDownload)

		} else if val.errInDownload != nil && val.errInResizing == nil {
			fmt.Println("Failure in resize ", val.avatarFilename, val.errInResizing)

		} else {
			fmt.Println("Successful download and resize ", val.filename, val.avatarFilename)
		}
	}

	waitGroup.Wait()

	// close all channels (there is no third channel. because it closed in ReadPictureUrls func)
	close(downloadedImagesFilenames)
	close(resizedImageChan)

	fmt.Println("All is ready")
}
