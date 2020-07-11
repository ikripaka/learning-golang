package image_loader

import (
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

const TopValueForGeneratingIndividualNamesForNewImages = 1024

// Loads pictures from the internet, using urls from the file
// folderPath - path to the folder where images would be saved
// urlsChannel - channel with urls
// channelWithFilenames - channel with filenames that would be save file names
func LoadPictures(folderPath string, urlsChannel chan string, channelWithFilenames chan string,
	waitGroup *sync.WaitGroup) {
	url, isChannelEmpty := <-urlsChannel
	for isChannelEmpty {
		filename := getFilename(folderPath, url, channelWithFilenames)

		response, err := http.Get(url)
		if err != nil {
			fmt.Println(folderPath + `\` + filename)
			out, err := os.Create(folderPath + `\` + filename)
			if err != nil {
				log.Fatal(err)
			}
			err = out.Close()
			if err != nil{
				log.Println("Error in out.Close()", filename)
			}

			err = response.Body.Close()
			if err != nil {
				log.Println("Error in response.Body.Close() ", filename)
			}

		} else {

			out, err := os.Create(folderPath + `\` + filename)
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(out, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			err = out.Close()
			if err != nil{
				log.Println("Error in out.Close()", filename)
			}

			err = response.Body.Close()
			if err != nil {
				log.Println("Error in response.Body.Close() ", filename)
			}
		}
		url, isChannelEmpty = <-urlsChannel
	}
	waitGroup.Done()
}

// Gets filename for file (from url/individual name)
// folderPath - folder path for correct path to the file
// url - url from what extracts filename
// filenameChannel - channel with filenames
func getFilename(folderPath string, url string, filenameChannel chan string) string {
	rand.Seed(time.Now().UnixNano())
	var regExpForFilename = regexp.MustCompile(`(?:[^/][-\w\.]+)+$`)
	var regExpForFileExtension = regexp.MustCompile(`(((-|\w)+)\.(jpg|png))$`)
	splitLineBySlash := strings.Split(url, `/`)

	if regExpForFilename.MatchString(url) &&
		regExpForFileExtension.MatchString(splitLineBySlash[cap(splitLineBySlash)-1]) {
		filename := regExpForFilename.FindString(url)

		if _, err := os.Stat(folderPath + `\` + filename); err == nil {
			regexMath :=
				regexp.MustCompile(`(.+?)(\.[^.]*$|$)`).FindStringSubmatch(splitLineBySlash[cap(splitLineBySlash)-1])
			filename = regexMath[cap(regexMath)-2] + " (" +
				strconv.Itoa(rand.Intn(TopValueForGeneratingIndividualNamesForNewImages)) + ")" +
				regexMath[cap(regexMath)-1]
		}
		filenameChannel <- filename
		fmt.Println(filename)
		return filename
	}

	filename := "Picture_№_" + strconv.Itoa(rand.Intn(TopValueForGeneratingIndividualNamesForNewImages)) + `.jpg`
	fmt.Println(filename)
	filenameChannel <- filename
	return filename

}
