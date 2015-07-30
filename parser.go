package main

import (
	"errors"
	"io/ioutil"
	"os"
)

// Parse test string content from stdin
func parseTestString() []rune {
	file, err := os.Stdin.Stat()
	if err != nil {
		exitWithError(err)
	}
	if (file.Mode() & os.ModeCharDevice) == 0 {
		testString, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			exitWithError(err)
		}
		return []rune(string(testString))
	} else {
		exitWithError(errors.New("No input specified"))
	}
	return []rune{}
}
