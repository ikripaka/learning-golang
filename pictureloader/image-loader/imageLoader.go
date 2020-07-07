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
	"time"
)

func LoadPictures(folderPath string, urlsChannel chan string) {
	url, isChannelEmpty := <-urlsChannel
	for isChannelEmpty {
		filename := getFilename(url, cap(urlsChannel))

		response, err := http.Get(url)
		if err != nil {
			fmt.Println(folderPath + `\` + filename)
			out, err := os.Create(folderPath + `\` + filename)
			if err != nil {
				log.Fatal(err)
			}
			response.Body.Close()
			out.Close()

		} else {

			out, err := os.Create(folderPath + `\` + filename)
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(out, response.Body)
			if err != nil {
				log.Fatal(err)
			}
			response.Body.Close()
			out.Close()
		}
		url, isChannelEmpty = <-urlsChannel
		fmt.Println(url, isChannelEmpty)
	}
	fmt.Println("finish")
}

// Gets filename for file (from url/individual name)
func getFilename(url string, channelCapacity int) string {
	rand.Seed(time.Now().UnixNano())
	var regExpForFilename = regexp.MustCompile(`(?:[^/][-\w\.]+)+$`)
	var regExpForFileExtension = regexp.MustCompile(`(((-|\w)+)\.(jpg|png))$`)
	splitLineBySlash :=strings.Split(url,`/`)
	if regExpForFilename.MatchString(url) && regExpForFileExtension.MatchString(splitLineBySlash[cap(splitLineBySlash)-1]){
		fmt.Println(regExpForFilename.FindString(url))
		return regExpForFilename.FindString(url)
	}
	fmt.Println("Picture_№_" + strconv.Itoa(rand.Intn(1024)) + `.jpg`)
	return "Picture_№_" + strconv.Itoa(rand.Intn(1024)) + `.jpg`
}