package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	tailer, err := tail.TailFile("/home/ikripaka/go/src/github.com/ikripaka/hello/first_code/file.log", tail.Config{
		Follow:    true,
		MustExist: true,
		ReOpen:    true,
	})
	if err != nil {
		tailer.Stop()
	}

	for line := range tailer.Lines {
		fmt.Println(line.Text)
	}
}

//cobra interactive
