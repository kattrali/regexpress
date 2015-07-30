package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

const delay = time.Millisecond

func main() {
	if err := termbox.Init(); err != nil {
		exitWithError(err)
	}
	defer termbox.Close()

	var regex []rune
	printer := NewPrinter(parseTestString())
	queue := make(chan termbox.Event)
	go func() {
		for {
			queue <- termbox.PollEvent()
		}
	}()

	printer.PrintSplash()
	for {
		select {
		case ev := <-queue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch != 0:
					regex = append(regex, ev.Ch)
				case ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyEsc:
					return
				case ev.Key == termbox.KeyBackspace2 || ev.Key == termbox.KeyBackspace:
					if len(regex) > 0 {
						regex = regex[0 : len(regex)-1]
					}
				}
			}
		default:
			printer.UpdateScreen(regex)
			time.Sleep(delay)
		}
	}
}

func exitWithError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
