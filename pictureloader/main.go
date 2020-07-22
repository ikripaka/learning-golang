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

	pictureUrls := make(chan Item)
	downloadedImagesFilenames := make(chan Item)
	resizedImageChan := make(chan Item)

	args := os.Args[1:]
	urlFilePath := args[0]
	folderPath := args[1]

	//urlFilePath := `E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\test.txt`
	//folderPath := `E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files`
	if _, err := IsPathsCorrect(urlFilePath, folderPath); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Read urls from file..")

	waitGroup.Add(1)
	go ReadPictureUrls(urlFilePath, pictureUrls, &waitGroup, &numOfUrls)

	waitGroup.Add(MAXDOWNLOADPROCESSES)

	fmt.Println("Download images..")

	for i := 0; i < MAXDOWNLOADPROCESSES; i++ {
		go LoadPictures(folderPath, pictureUrls, downloadedImagesFilenames, &waitGroup)
	}

	waitGroup.Add(numCPU)

	fmt.Println("Scale images..")
	for i := 0; i < numCPU; i++ {
		go MakeAvatars(downloadedImagesFilenames, resizedImageChan, &waitGroup, &counter, numOfUrls)
	}

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

	close(downloadedImagesFilenames)
	close(resizedImageChan)

	fmt.Println("All is ready")
}
