package path_checker

import (
	"errors"
	"regexp"
)

func IsPathsCorrect(consoleInput []string) (bool, error) {
	var regExpForFilePath = regexp.MustCompile(`(.+\\)*(.+)\.(txt)`)
	if !regExpForFilePath.MatchString(consoleInput[0]) {
		return false, errors.New("incorrect path to the file with urls")
	}
	var regExpForFolderPath = regexp.MustCompile(`^.+\\([^.]+)[^\.]+$`)
	if !regExpForFolderPath.MatchString(consoleInput[1]) {
		return false, errors.New("incorrect path to the folder in which images would be saved")
	}

	return true, nil
}
