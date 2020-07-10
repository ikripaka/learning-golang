package main

import (
	"bufio"
	"fmt"
	"github.com/ikripaka/learning-golang/pictureloader/filereader"
	image_loader "github.com/ikripaka/learning-golang/pictureloader/image-loader"
	image_reducer "github.com/ikripaka/learning-golang/pictureloader/image-reducer"
	"log"
	"os"
	"sync"
)

func main() {
	//args := os.Args[1:]
	//if _, err := pathChecker.IsPathsCorrect(args); err != nil {
	//	log.Fatal(err)
	//}
	pictureUrls :=filereader.ReadPictureUrls(`E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\test.txt`)
	channelWithFilenames:= make(chan string, cap(pictureUrls))
	var waitGroup sync.WaitGroup
	waitGroup.Add(15)
	for i:=0; i< 15; i++ {
		 go image_loader.LoadPictures(`E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files`, pictureUrls, channelWithFilenames, &waitGroup)
	}
	waitGroup.Wait()
	close(channelWithFilenames)
	fmt.Println("-------")
	image_reducer.MakeAvatars(`E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files`, channelWithFilenames,&waitGroup)
	fmt.Println("Finish")
}

func readFromConsole() string {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Write directory:")
	inputText, err := consoleReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return inputText
}
