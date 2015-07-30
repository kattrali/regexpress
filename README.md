# regexpress

CLI Go regular expression tester

![scrot](https://github.com/kattrali/regexpress/raw/master/scrot.png)

# usage

Pipe test input into the tester via standard input, then enter an expression to find matches:

    echo "hello" | regexpress

Press `Esc` or `Ctrl+C` to close `regexpress`

# dependencies

Regexpress depends on [termbox-go](https://github.com/nsf/termbox-go):

    go get -u github.com/nsf/termbox-go

# installation

	go get github.com/kattrali/regexpress
    go install github.com/kattrali/regexpress
