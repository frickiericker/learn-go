package words

import (
	"unicode"
)

func categorizeInput(char rune) inputCategory {
	if unicode.IsLetter(char) {
		return wordRune
	} else {
		return nonWordRune
	}
}

func Count(text string) int64 {
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
	return count
}
