package image_reducer

import (
	"bufio"
	"fmt"
	imageLoader "github.com/ikripaka/learning-golang/pictureloader/image-loader"
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

func MakeAvatars(folderPath string, filenamesChannel chan string, waitGroup *sync.WaitGroup) {
	for imgFilename, isEmpty := <-filenamesChannel; isEmpty; {
		fmt.Println(folderPath+`\`+imgFilename, "processing")
		originalFile, err := os.Open(folderPath + `\` + imgFilename)
		if err != nil {
			originalFile.Close()
			fmt.Println("Can`t open image", imgFilename)
			log.Fatal(err)

		} else {

			decodedImage,_, err := image.Decode(bufio.NewReader(originalFile))
			if err != nil {
				originalFile.Close()
				fmt.Println("Problems with decod1e", imgFilename)
				log.Fatal(err)
			}

			configDecode, _, err := image.DecodeConfig(bufio.NewReader(originalFile))

			if err != nil {
				originalFile.Close()
				fmt.Println("Problems 2with decode config", imgFilename)
				log.Fatal(err)

			}

			width, height := getReducedPixelSize(configDecode.Width, configDecode.Height)


			fmt.Println(width, height)
			resizedImg := resize.Resize(width, height, decodedImage, resize.Lanczos3)

			fileneme := getFilename(folderPath, imgFilename) //delete
			fmt.Println(fileneme)
			outputFile, err := os.Create(folderPath + `\` + fileneme)

			if err != nil {
				originalFile.Close()
				fmt.Println("Can`t decode image", imgFilename)
				log.Fatal(err)
			}

			//switch format {
			//case "png":
			//	png.Encode(outputFile, resizedImg)
			//	fmt.Println("png")
			//case "jpeg", "jpg":
			//	jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 100})
			//	fmt.Println("jpeg")
			//default:
			//	jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 100} )
			//	fmt.Println("jpeg")
			jpeg.Encode(outputFile, resizedImg, &jpeg.Options{Quality: 100})
			fmt.Println("jpeg")
			outputFile.Close()
			originalFile.Close()
		}
		imgFilename, isEmpty = <-filenamesChannel
		fmt.Println("finish")
	}
}

func getFilename(folderPath string, filename string) string {

	rand.Seed(time.Now().UnixNano())
	regexMath := regexp.MustCompile(`(.+?)(\.[^.]*$|$)`).FindStringSubmatch(filename)

	if _, err := os.Stat(folderPath + `\` + regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]); err == nil {
		filename = regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
		fmt.Println(filename)
		return regexMath[cap(regexMath)-2] + " (avatar) " + regexMath[cap(regexMath)-1]
	}
	filename = regexMath[cap(regexMath)-2] + " (" + strconv.Itoa(rand.Intn(imageLoader.TopValueForGeneratingIndividualNamesForNewImages)) + ")" + regexMath[cap(regexMath)-1]
	fmt.Println(filename)

	return filename
}

func getReducedPixelSize(width int, height int) (uint, uint) {
	pixelsPerOneWidthPixel := uint(height / width)
	reducedImageHeight := pixelsPerOneWidthPixel * AvatarWidthSize
	return AvatarWidthSize, reducedImageHeight
}