package main

import (
	"fmt"
	"log"
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

// E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\test.txt E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files
func main() {
	numCPU := runtime.NumCPU()

	//args := os.Args[1:]
	//urlFilePath := args[0]
	//folderPath := args[1]
	urlFilePath := `E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\test.txt`
	folderPath := `E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files`
	if _, err := IsPathsCorrect(urlFilePath, folderPath); err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	fmt.Println("Read urls from file..")

	pictureUrls := make(chan string)

	go ReadPictureUrls(urlFilePath, pictureUrls)

	channelWithFilenames := make(chan string)
	waitGroup.Add(MAXDOWNLOADPROCESSES)

	fmt.Println("Download images..")
	for i := 0; i < MAXDOWNLOADPROCESSES; i++ {
		go LoadPictures(urlFilePath, pictureUrls, channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()
	close(channelWithFilenames)
	waitGroup.Add(numCPU)

	fmt.Println("Scale images..")
	for i := 0; i < numCPU; i++ {
		go MakeAvatars(urlFilePath, channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("All is ready")
}
