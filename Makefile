
default: build

build:
	go build regexpress.go parser.go printer.go formatter.go
