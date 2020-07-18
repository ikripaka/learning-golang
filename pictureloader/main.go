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

func main() {
	numCPU := runtime.NumCPU()

	args := os.Args[1:]
	urlFilePath := args[0]
	folderPath := args[1]

	if _, err := IsPathsCorrect(urlFilePath, folderPath); err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	fmt.Println("Read urls from file..")

	pictureUrls := ReadPictureUrls(args[0])
	channelWithFilenames := make(chan string, cap(pictureUrls))
	waitGroup.Add(MAXDOWNLOADPROCESSES)

	fmt.Println("Download images..")
	for i := 0; i < MAXDOWNLOADPROCESSES; i++ {
		go LoadPictures(args[1], pictureUrls, channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()
	close(channelWithFilenames)
	waitGroup.Add(numCPU)

	fmt.Println("Scale images..")
	for i := 0; i < numCPU; i++ {
		go MakeAvatars(args[1], channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("All is ready")
}
