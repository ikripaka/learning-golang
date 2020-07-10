package filereader

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

//Reads file with urls and pushes them to the buffered channel
func ReadPictureUrls( filePath string) chan string {
	channelCapacity, err:=  countLines(filePath)
	urlsChannel := make (chan string , channelCapacity)
	fmt.Println(cap(urlsChannel))

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("read: ", scanner.Text())
		urlsChannel <- scanner.Text()
		err = scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	close(urlsChannel)
	return urlsChannel

}

//Counts file lines for creating buffered channel
func countLines(filePath string) (int, error) {
	reader, err := os.Open(filePath)
	if err != nil{
		log.Fatal(err)
	}
	var count int
	const lineBreak = '\n'

	buffer := make([]byte, bufio.MaxScanTokenSize)
	for {
		bufferSize, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return 0, err
		}
		var bufferPosition int
		for {
			i := bytes.IndexByte(buffer[bufferPosition:], lineBreak)
			if i == -1 || bufferSize == bufferPosition {
				break
			}
			bufferPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	return count+1, nil
}
