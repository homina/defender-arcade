package main

import (
	"defender-arcade/internal"
	"defender-arcade/pkg/file"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrInvalidCommand = errors.New("please input path to the file")
	ErrOutOfRange     = errors.New("invalid range")
)

const (
	LAYOUT    = "1504"
	SEPARATOR = " "
)

func main() {
	defer exception()
	lines := getLines()
	if len(lines) == 0 || len(lines) > 100 {
		panic(ErrOutOfRange)
	}

	max, err := internal.MaxSliceInTimeRange(lines, internal.TimeRangeParams{
		Layout:    LAYOUT,
		Separator: SEPARATOR,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(max)
}

// check argument and process the argument
func getLines() []string {
	lines := []string{}
	var err error
	if len(os.Args) > 1 {
		lines, err = file.GetEachLine(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		panic(ErrInvalidCommand)
	}

	return lines
}

func exception() {
	if r := recover(); r != nil {
		log.Printf("%+v\n", r)
	}
}
