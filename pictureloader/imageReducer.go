package main

import (
	"bufio"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"regexp"
	"strconv"
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
// config - program configuration that contains all variables that need
// counter - counter for goroutines

func MakeAvatars(counter *int, config *ProgramConfig) {

	//works until counter < number of all urls that need to be processed
	for item, _ := <-config.downloadedImagesFilenames; *counter < cap(config.pictureUrlsChan); *counter++ {
		filepath := config.imageFolderPath + `\` + item.filename
		originalFile, err := os.Open(filepath)

		if err != nil {
			handleClosingErrInOriginalFile(originalFile, item, filepath)
			item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Image opening", imgFilepath: filepath}
			config.resizedImageChan <- item

			item, _ = <-config.downloadedImagesFilenames
			continue

		} else { //if program can open downloaded file

			// decodes image file to Image
			decodedImage, _, err := image.Decode(bufio.NewReader(originalFile))

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item, filepath)
				item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Decoding file to Image", imgFilepath: filepath}

				config.resizedImageChan <- item
				item, _ = <-config.downloadedImagesFilenames

				continue
			}

			//  returns reader to the beginning of file
			_, err = originalFile.Seek(0, 0)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item, filepath)
				item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Reader.Seek()", imgFilepath: filepath}

				config.resizedImageChan <- item
				item, _ = <-config.downloadedImagesFilenames
				continue
			}

			// finds out file format/size/color signature
			configDecode, format, err := image.DecodeConfig(bufio.NewReader(originalFile))

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item, filepath)
				item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Decoding image configuration", imgFilepath: filepath}

				config.resizedImageChan <- item
				item, _ = <-config.downloadedImagesFilenames

				continue
			}

			// resize image size according to original image proportion
			width, height := getReducedPixelSize(configDecode.Width, configDecode.Height)
			resizedImg := resize.Resize(width, height, decodedImage, resize.MitchellNetravali)

			// creates new file for avatar
			newFilename := getFilenameForAvatars(config.imageFolderPath, item.filename)
			outputFile, err := os.Create(config.imageFolderPath + `\` + newFilename)

			if err != nil {
				handleClosingErrInOriginalFile(originalFile, item, filepath)
				item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Creating new avatar file", imgFilepath: filepath}

				config.resizedImageChan <- item
				item, _ = <-config.downloadedImagesFilenames

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
				item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Encoding avatar file to format " + format, imgFilepath: filepath}

				handleClosingErrInOriginalFile(originalFile, item, filepath)

				err = outputFile.Close()
				if err != nil {
					item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Closing avatar file ", imgFilepath: filepath}
				}
			}

			handleClosingErrInOriginalFile(originalFile, item, filepath)
			err = outputFile.Close()

		}
		config.resizedImageChan <- item
		item, _ = <-config.downloadedImagesFilenames
	}
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
		return regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
	}

	// if file with same filename exist
	filename = regexMath[cap(regexMath)-2] + " (" + strconv.Itoa(rand.Intn(TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES)) +
		")" + regexMath[cap(regexMath)-1]
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
// filepath - filepath that used for creating error
func handleClosingErrInOriginalFile(originalFile *os.File, item Item, filepath string) {
	err := originalFile.Close()

	if err != nil {
		item.err = &PictureLoaderError{problemOccurrence: "RESIZING", problemDescription: "Closing downloaded file", imgFilepath: filepath}
	}
}
