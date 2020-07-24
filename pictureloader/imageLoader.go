package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES = 1024

// Loads pictures from the internet, using urls from the file
// folderPath - path to the folder where images would be saved
// urlsChannel - channel with urls
// channelWithFilenames - channel with filenames that would be save file names
func LoadPictures(folderPath string, urlsChannel chan Item, channelWithFilenames chan Item,
	waitGroup *sync.WaitGroup) {
	item, isChannelEmpty := <-urlsChannel
	item.imageFolderPath = folderPath

	// works until another goroutine read file
	for isChannelEmpty {
		item.filename = getFilenameForDownloadedImages(folderPath, item.url)

		// get response from url
		response, err := http.Get(item.url)

		if err != nil {
			// if all isn't ok, make empty file
			fmt.Println(folderPath + `\` + item.filename)
			out, err := os.Create(folderPath + `\` + item.filename)
			if err != nil {
				log.Fatal(err)
			}
			err = out.Close()
			if err != nil {
				item.errInDownload = errors.New("Error in out.Close()" + item.filename)
			}

			if response != nil {
				err = response.Body.Close()
			} else {
				item.errInDownload = errors.New("Empty response in file" + item.filename)
			}

			if err != nil {
				item.errInDownload = errors.New("Error in response.Body.Close() " + item.filename)
			}

		} else {
			// if all is ok create file and copy response.Body to it

			out, err := os.Create(folderPath + `\` + item.filename)
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			err = out.Close()
			if err != nil {
				item.errInDownload = errors.New("Error in out.Close() " + item.filename)
			}

			err = response.Body.Close()
			if err != nil {
				item.errInDownload = errors.New("Error in response.Body.Close() " + item.filename)
			}
		}

		channelWithFilenames <- item
		item, isChannelEmpty = <-urlsChannel
	}
	waitGroup.Done()
}

// Gets filename for file (from url/individual name)
// folderPath - folder path for correct path to the file
// url - url from what extracts filename
// filenameChannel - channel with filenames
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
		fmt.Println(filename)
		return filename
	}

	// if regexp cannot extract filename from url
	filename = "Picture_№_" + strconv.Itoa(rand.Intn(TOPVALUEFORGENERATINGINDIVIDUALNAMESFORNESIMAGES)) + `.jpg`
	fmt.Println(filename)
	return filename

}
