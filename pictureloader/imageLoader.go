package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES = 2048

// Loads pictures from the internet, using urls from the file
// config - program configuration that contains all variables that need
func LoadPictures(config *ProgramConfig) {
	item, isChannelEmpty := <-config.pictureUrlsChan

	// works until another goroutine read file
	for isChannelEmpty {
		item.filename = getFilenameForDownloadedImages(config.imageFolderPath, item.url)
		filepath := config.imageFolderPath + `\` + item.filename

		// get response from url
		response, err := http.Get(item.url)

		if err != nil {
			// if all isn't ok, make empty file
			out, err := os.Create(filepath)
			if err != nil || out == nil {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: "Can't create file",imgFilepath: filepath }
			}
			err = out.Close()
			if err != nil {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: "Can't close file",imgFilepath: filepath }
			}

			if response != nil {
				err = response.Body.Close()
			} else {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: ("Empty response url: " + item.url) ,imgFilepath: filepath }
			}

			if err != nil {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: "Closing response url: " + item.url ,imgFilepath: filepath }
			}

		} else {
			// if all is ok create file and copy response.Body to it

			out, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			err = out.Close()
			if err != nil {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: "Closing file",imgFilepath: filepath }
			}

			err = response.Body.Close()
			if err != nil {
				item.err = &PictureLoaderError{ problemOccurrence: "DOWNLOAD",problemDescription: "Closing response url:" + item.url ,imgFilepath: filepath }
			}
		}

		config.downloadedImagesFilenames <- item
		item, isChannelEmpty = <-config.pictureUrlsChan
	}
}

// Gets filename for file (from url/individual name)
// folderPath - folder path for correct path to the file
// url - url from what extracts filename
func getFilenameForDownloadedImages(folderPath string, url string) string {
	rand.Seed(time.Now().UnixNano())
	var regExpForFilename = regexp.MustCompile(`(?:[^/][-\w\.]+)+$`)
	var regExpForFileExtension = regexp.MustCompile(`(((-|\w)+)\.(jpg|png))$`)
	var filename string
	splitLineBySlash := strings.Split(url, `/`)

	// if in url regexp find filename in the end of it like ...\town.jpg
	if regExpForFilename.MatchString(url) &&
		regExpForFileExtension.MatchString(splitLineBySlash[cap(splitLineBySlash)-1]) {
		filename = regExpForFilename.FindString(url)

		// if another image with same filename has already exist
		if _, err := os.Stat(folderPath + `\` + filename); err == nil {
			regexMath :=
				regexp.MustCompile(`(.+?)(\.[^.]*$|$)`).FindStringSubmatch(splitLineBySlash[cap(splitLineBySlash)-1])
			filename = regexMath[cap(regexMath)-2] + " (" +
				strconv.Itoa(rand.Intn(TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES)) + ")" +
				regexMath[cap(regexMath)-1]

		}
		return filename
	}

	// if regexp cannot extract filename from url
	filename = "Picture_#_" + strconv.Itoa(rand.Intn(TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES)) + `.jpg`
	return filename
}
