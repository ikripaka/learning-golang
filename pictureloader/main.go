package main

import (
	"bufio"
	"fmt"
	"github.com/ikripaka/learning-golang/pictureloader/filereader"
	image_loader "github.com/ikripaka/learning-golang/pictureloader/image-loader"
	"log"
	"os"
)

func main() {
	//args := os.Args[1:]
	//if _, err := pathChecker.IsPathsCorrect(args); err != nil {
	//	log.Fatal(err)
	//}
	pictureUrls := make(chan string)


	pictureUrls =filereader.ReadPictureUrls(`E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\test.txt`)
	image_loader.LoadPictures(`E:\gocode\src\github.com\ikripaka\learning-golang\pictureloader\load_files`, pictureUrls)
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
