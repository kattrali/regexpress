package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	Prompt        = "Regex: "
	errorFmt      = "Invalid expression: %v - %v"
	groupFmt      = "% 5d % MAXLENv %v"
	matchTitleFmt = "Index %MAXLENv Text"
	maxLen        = "MAXLEN"
	minNameLen    = 4
	nameTitle     = "Name"
)

func formatTestString(content []rune) []rune {
	lines := strings.SplitAfter(string(content), "\n")
	if len(lines) > TestStringDisplayHeight {
		return []rune(strings.Join(lines[:TestStringDisplayHeight], ""))
	}
	return content
}

func formatPrompt(termw int, regex []rune) []rune {
	const promptLen = len(Prompt)
	regexLen := len(regex)
	var regexSegment []rune
	if regexLen > termw-promptLen {
		start := regexLen - termw + promptLen
		regexSegment = regex[start : regexLen-1]
	} else {
		regexSegment = regex
	}
	return append([]rune(Prompt), regexSegment...)
}

func formatMatchTitle(nameLength int) []rune {
	format := strings.Replace(matchTitleFmt, maxLen, fmt.Sprintf("%d", nameLength), -1)
	return []rune(fmt.Sprintf(format, nameTitle))
}

func formatMatchGroup(index, nameLength int, name, content string) []rune {
	length := math.Max(float64(nameLength), float64(minNameLen))
	format := strings.Replace(groupFmt, maxLen, fmt.Sprintf("%d", int(length)), -1)
	return []rune(fmt.Sprintf(format, index, name, content))
}

func formatRegexError(err error, regex []rune) []rune {
	return []rune(fmt.Sprintf(errorFmt, err, string(regex)))
}
