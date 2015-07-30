package main

import (
	"github.com/nsf/termbox-go"
	"regexp"
)

const (
	errColor                = termbox.ColorRed
	foreground              = termbox.ColorWhite
	titleColor              = termbox.ColorGreen
	subtitleColor           = termbox.ColorBlue
	background              = termbox.ColorBlack
	RegexpressTitle         = "Regexpress : Go Regular Expression Tester"
	RegexpressSubtitle      = "Type an expression to continue"
	TestStringTitle         = "Test String:"
	MatchesTitle            = "Match Groups:"
	TestStringDisplayHeight = 6
	PresentationHeight      = 3
)

type Printer struct {
	Test  []rune
	regex []rune
	termw int
	termh int
}

// Create a printer with a regex test string
func NewPrinter(test []rune) *Printer {
	return &Printer{Test: test}
}

// Clear the screen and display the splash
func (p *Printer) PrintSplash() {
	termbox.Clear(background, background)
	printCentered(PresentationHeight, titleColor, []rune(RegexpressTitle))
	printCentered(PresentationHeight+2, subtitleColor, []rune(RegexpressSubtitle))
	termbox.Flush()
}

// Update the screen layout by matching the test
// string against a regex
func (p *Printer) UpdateScreen(regex []rune) {
	if compare(regex, p.regex) {
		return
	}
	termbox.Clear(background, background)
	p.termw, p.termh = termbox.Size()
	p.regex = regex
	p.printTestString()
	p.printMatch()
	p.printPrompt()
	termbox.Flush()
}

func (p *Printer) printPrompt() {
	printContent(0, p.termh-1, foreground, formatPrompt(p.termw, p.regex))
}

func (p *Printer) printTestString() {
	printTitle(0, TestStringTitle)
	printContent(0, 1, foreground, formatTestString(p.Test))
}

func (p *Printer) printMatch() {
	re, err := regexp.Compile(string(p.regex))
	if err != nil {
		p.printRegexError(err)
		return
	}
	printTitle(TestStringDisplayHeight, MatchesTitle)
	groups := re.FindStringSubmatch(string(p.Test))
	p.printMatchGroups(groups, re.SubexpNames())
}

func (p *Printer) printMatchGroups(groups, names []string) {
	var maxNameLen = 0
	for _, name := range names {
		length := len(name)
		if length > maxNameLen {
			maxNameLen = length
		}
	}
	printContent(0, TestStringDisplayHeight+1, subtitleColor, formatMatchTitle(maxNameLen))
	for idx, group := range groups {
		content := formatMatchGroup(idx, maxNameLen, names[idx], group)
		printContent(0, idx+2+TestStringDisplayHeight, foreground, content)
	}
}

func (p *Printer) printRegexError(err error) {
	printCentered(TestStringDisplayHeight, errColor, formatRegexError(err, p.regex))
}

func printTitle(y int, content string) {
	printContent(0, y, titleColor, []rune(content))
}

func printCentered(y int, color termbox.Attribute, content []rune) {
	termw, _ := termbox.Size()
	printContent((termw-len(content))/2, y, color, content)
}

func printContent(x, y int, color termbox.Attribute, content []rune) {
	var indentX = x
	for _, char := range content {
		if char == '\n' {
			y++
			indentX = x
			continue
		}
		termbox.SetCell(indentX, y, char, color, background)
		indentX++
	}
}

func compare(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, char := range a {
		if b[idx] != char {
			return false
		}
	}
	return true
}
