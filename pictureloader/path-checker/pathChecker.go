package path_checker

import (
	"errors"
	"fmt"
	"regexp"
)

// checks whether the path is recorded correctly with regular expressions
// consoleInput - []string
// first - filepath to .txt file with urls
// second - folder path where images would be stored
func IsPathsCorrect(consoleInput []string) (bool, error) {
	fmt.Println(consoleInput)
	var regExpForFilePath = regexp.MustCompile(`(.+\\)*(.+)\.(txt)`)

	if !regExpForFilePath.MatchString(consoleInput[0]) {
		fmt.Println(consoleInput[0], regExpForFilePath.MatchString(consoleInput[0]) )
		return false, errors.New("incorrect path to the file with urls")
	}

	var regExpForFolderPath = regexp.MustCompile(`^.+\\([^.]+)[^\.]+$`)
	fmt.Println(consoleInput[1], regExpForFolderPath.MatchString(consoleInput[1]))
	if !regExpForFolderPath.MatchString(consoleInput[1]) {
		return false, errors.New("incorrect path to the folder in which images would be saved")
	}

	return true, nil
}
