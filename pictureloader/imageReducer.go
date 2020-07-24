package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
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

func MakeAvatars(filenamesChannel chan Item, resizedImageChan chan Item, waitGroup *sync.WaitGroup, counter *int, numOfUrls int) {

	//works until counter < number of all urls that need to be processed
	for item, _ := <-filenamesChannel; *counter < numOfUrls; *counter++ {
		originalFile, err := os.Open(item.imageFolderPath + `\` + item.filename)

		if err != nil {
			handleClosingErrInOriginalFile(originalFile, item)
			item.errInResizing = errors.New("Can`t open image ")
			item, _ = <-filenamesChannel
			continue

		} else { //if program can open downloaded file

			// decodes image file to Image
			decodedImage, _, err := image.Decode(bufio.NewReader(originalFile))

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item)
				item.errInResizing = errors.New("Problems with decode ")
				item, _ = <-filenamesChannel
				continue
			}

			//  returns reader to the beginning of file
			_, err = originalFile.Seek(0, 0)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item)
				item.errInResizing = errors.New("Problems with Reader.Seek() ")
				item, _ = <-filenamesChannel
				continue
			}

			// finds out file format/size/color signature
			configDecode, format, err := image.DecodeConfig(bufio.NewReader(originalFile))

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item)
				item.errInResizing = errors.New("Problems with decode config ")
				item, _ = <-filenamesChannel
				continue
			}

			// resize image size according to original image proportion
			width, height := getReducedPixelSize(configDecode.Width, configDecode.Height)
			resizedImg := resize.Resize(width, height, decodedImage, resize.MitchellNetravali)

			// creates new file for avatar
			newFilename := getFilenameForAvatars(item.imageFolderPath, item.filename)
			outputFile, err := os.Create(item.imageFolderPath + `\` + newFilename)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item)

				item.errInResizing = errors.New("Can`t create " + newFilename)
				item, _ = <-filenamesChannel
				continue
			}

			// saves image depending on input file format
			switch format {
			case "png":
				err = png.Encode(outputFile, resizedImg)
			case "jpeg", "jpg":
				err = jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 100})
			}

			if err != nil {
				item.errInResizing = errors.New("Error in encoding: ")

				handleClosingErrInOriginalFile(originalFile, item)

				err = outputFile.Close()
				if err != nil {
					item.errInResizing = errors.New("Error in closing output file ")
				}
			}

			handleClosingErrInOriginalFile(originalFile, item)
			err = outputFile.Close()
		}
		resizedImageChan <- item
		item, _ = <-filenamesChannel

	}
	waitGroup.Done()
}

// Get'f filename depending on filename name
// Adds to the filename '(avatar)' or if file exists gives different name
// folderPath - path where scaled images would be stored
// filename - file name
func getFilenameForAvatars(folderPath string, filename string) string {
	rand.Seed(time.Now().UnixNano())
	regexMath := regexp.MustCompile(`(.+?)(\.[^.]*$|$)`).FindStringSubmatch(filename)

	// if file with same filename not exist
	if _, err := os.Stat(folderPath + `\` + regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]); os.IsNotExist(err) {
		filename = regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
		fmt.Println(filename)
		return regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
	}

	// if file with same filename exist
	filename = regexMath[cap(regexMath)-2] + " (" + strconv.Itoa(rand.Intn(TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES)) +
		")" + regexMath[cap(regexMath)-1]
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
// item - Item that contain all necessary info about file
func handleClosingErrInOriginalFile(originalFile *os.File, item Item) {
	err := originalFile.Close()

	if err != nil {
		item.errInResizing = errors.New("Error in closing original file ")
	}
}
