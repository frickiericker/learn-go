package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/frickiericker/learn-go/02-wordcount/words"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	filepaths := os.Args[1:]

	for _, filepath := range filepaths {
		count, err := countWordsInFile(filepath)
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%d\n", filepath, count)
	}
	return nil
}

func countWordsInFile(filepath string) (int64, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return 0, err
	}
	text := string(data)
	return words.Count(text), nil
}
