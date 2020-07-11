package main

import (
	"fmt"
	"github.com/ikripaka/learning-golang/pictureloader/filereader"
	image_loader "github.com/ikripaka/learning-golang/pictureloader/image-loader"
	image_reducer "github.com/ikripaka/learning-golang/pictureloader/image-reducer"
	path_checker "github.com/ikripaka/learning-golang/pictureloader/path-checker"
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

const MaxNumberOfThreads = 4

func main() {
	args := os.Args[1:]

	if _, err := path_checker.IsPathsCorrect(args); err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	fmt.Println("*", "Read urls from file..")
	pictureUrls := filereader.ReadPictureUrls(args[0])
	channelWithFilenames := make(chan string, cap(pictureUrls))
	waitGroup.Add(15)

	fmt.Println("**", "Download images..")
	for i := 0; i < 15; i++ {
		go image_loader.LoadPictures(args[1], pictureUrls, channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()
	close(channelWithFilenames)
	waitGroup.Add(MaxNumberOfThreads)

	fmt.Println("***", "Scale images..")
	for i := 0; i < MaxNumberOfThreads; i++ {
		go image_reducer.MakeAvatars(args[1], channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("****", "All is ready")
}
