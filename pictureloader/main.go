package main

import (
	"fmt"
	"os"
	"runtime"
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
	url            string
	filename       string
	avatarFilename string

	err error
}

// this struct helps to contain all program configuration variable at one place
type ProgramConfig struct {
	imageFolderPath string

	pictureUrlsChan           chan Item
	downloadedImagesFilenames chan Item
	resizedImageChan          chan Item
}

// this error represents one error type that can appear in program execution
type PictureLoaderError struct {
	problemOccurrence  string
	problemDescription string
	imgFilepath        string
}

func (e *PictureLoaderError) Error() string {
	return "image" + e.imgFilepath + "has problems with " + e.problemOccurrence + " " + e.problemDescription
}

func main() {

	// get variables
	urlFilePath, folderPath := getArgs(os.Args[1:])

	config := ProgramConfig{
		imageFolderPath:           folderPath,
		downloadedImagesFilenames: make(chan Item),
		resizedImageChan:          make(chan Item)}

	fmt.Println("Reading urls from file..")

	// read picture urls and push data to pictureUrlsChan
	ReadPictureUrls(urlFilePath, &config)

	fmt.Println("Downloading images..")

	// load pictures from internet and push data to downloadedImagesFilenames
	for i := 0; i < MAXDOWNLOADPROCESSES; i++ {
		go LoadPictures(&config)
	}

	var counter int

	// scale images and push data to resizedImageChan
	fmt.Println("Scaling images..")
	for i := 0; i < runtime.NumCPU(); i++ {
		go MakeAvatars(&counter, &config)
	}

	// show program results
	for i := 0; i < cap(config.pictureUrlsChan); i++ {

		item, _ := <-config.resizedImageChan

		if item.err != nil {
			fmt.Println(item.err)
		} else {
			fmt.Println("Successful download and resize ", item.filename, item.avatarFilename)
		}
	}
	fmt.Println("All is ready")
}
