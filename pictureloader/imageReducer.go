package main

import (
	"bufio"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	_ "image/jpeg"
	_ "image/png"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

const AvatarWidthSize = 64 //px

// MakesAvatars with using "github.com/nfnt/resize"
// folderPath - path to the folder where images would be stored
// filenamesChannel - channel that contains in it all filenames
// waitGroup - sync.WaitGroup that helps to handle goroutines
func MakeAvatars(folderPath string, filenamesChannel chan string, waitGroup *sync.WaitGroup) {
	for imgFilename, isEmpty := <-filenamesChannel; isEmpty; {
		originalFile, err := os.Open(folderPath + `\` + imgFilename)

		if err != nil {
			handleClosingErrInOriginalFile(originalFile)
			log.Println("Can`t open image", imgFilename)
			imgFilename, isEmpty = <-filenamesChannel
			continue

		} else {
			decodedImage, _, err := image.Decode(bufio.NewReader(originalFile))
			if err != nil {
				handleClosingErrInOriginalFile(originalFile)
				log.Println("Problems with decode", imgFilename)
				imgFilename, isEmpty = <-filenamesChannel
				continue
			}
			_, err = originalFile.Seek(0, 0)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile)
				log.Println("Problems with Reader.Seek()", imgFilename)
				imgFilename, isEmpty = <-filenamesChannel
				continue
			}

			configDecode, format, err := image.DecodeConfig(bufio.NewReader(originalFile))

			if err != nil {
				handleClosingErrInOriginalFile(originalFile)
				log.Println("Problems with decode config", imgFilename)
				imgFilename, isEmpty = <-filenamesChannel
				continue
			}

			width, height := getReducedPixelSize(configDecode.Width, configDecode.Height)

			resizedImg := resize.Resize(width, height, decodedImage, resize.MitchellNetravali)

			newFilename := getFilename(folderPath, imgFilename)
			outputFile, err := os.Create(folderPath + `\` + newFilename)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile)

				fmt.Println("Can`t create", newFilename)
				imgFilename, isEmpty = <-filenamesChannel
				continue
			}

			switch format {
			case "png":
				err = png.Encode(outputFile, resizedImg)
			case "jpeg", "jpg":
				err = jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 100})
			}
			if err != nil {
				log.Println("Error in encoding:", newFilename)

				handleClosingErrInOriginalFile(originalFile)

				err = outputFile.Close()
				if err != nil {
					log.Println("Error in closing output file")
				}
			}

			handleClosingErrInOriginalFile(originalFile)
			err = outputFile.Close()

		}
		imgFilename, isEmpty = <-filenamesChannel
	}
	waitGroup.Done()
}

// Get'f filename depending on filename name
// Adds to the filename '(avatar)' or if file exists gives different name
// folderPath - path where scaled images would be stored
// filename - file name
func getFilename(folderPath string, filename string) string {

	rand.Seed(time.Now().UnixNano())
	regexMath := regexp.MustCompile(`(.+?)(\.[^.]*$|$)`).FindStringSubmatch(filename)

	if _, err := os.Stat(folderPath + `\` + regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]); os.IsNotExist(err) {
		filename = regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
		fmt.Println(filename)
		return regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
	}

	filename = regexMath[cap(regexMath)-2] + " (" + strconv.Itoa(rand.Intn(TopValueForGeneratingIndividualNamesForNewImages)) + ")" + regexMath[cap(regexMath)-1]
	fmt.Println(filename)
	return filename
}

// Calculates reduced pixel size depending on const AvatarWidthSize
// width - original image width
// height - original image height
func getReducedPixelSize(width int, height int) (uint, uint) {
	pixelsPerOneWidthPixel := float32(height) / float32(width)
	reducedImageHeight := pixelsPerOneWidthPixel * AvatarWidthSize
	return AvatarWidthSize, uint(reducedImageHeight)
}

// Handles closing of opened file
// originalFile - file
func handleClosingErrInOriginalFile(originalFile *os.File) {
	err := originalFile.Close()

	if err != nil {
		log.Print("Error in closing original file")
	}
}
