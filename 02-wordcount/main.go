package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
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
	return countWords(text)
}

type scannerState int

const (
	nonWord scannerState = iota
	inWord
	exitWord
)

type inputCategory int

const (
	wordRune inputCategory = iota
	nonWordRune
	endOfInput
)

func makeTransitionTable() [3][3]scannerState {
	var table [3][3]scannerState

	table[nonWord][nonWordRune] = nonWord
	table[nonWord][wordRune] = inWord
	table[nonWord][endOfInput] = nonWord

	table[inWord][nonWordRune] = exitWord
	table[inWord][wordRune] = inWord
	table[inWord][endOfInput] = exitWord

	table[exitWord][nonWordRune] = nonWord
	table[exitWord][endOfInput] = nonWord

	return table
}

var transitionTable = makeTransitionTable()

func makeTransition(state scannerState, cat inputCategory) scannerState {
	return transitionTable[state][cat]
}

func categorizeInput(char rune) inputCategory {
	if unicode.IsLetter(char) {
		return wordRune
	} else {
		return nonWordRune
	}
}

func countWords(text string) (int64, error) {
	count := int64(0)
	state := nonWord

	for _, char := range text {
		state = makeTransition(state, categorizeInput(char))
		if state == exitWord {
			count += 1
		}
	}
	if makeTransition(state, endOfInput) == exitWord {
		count += 1
	}

	return count, nil
}
