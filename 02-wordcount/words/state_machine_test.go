package words

import (
	"testing"
)

func Test_makeTransition(t *testing.T) {
	test := func(state scannerState, input inputCategory, expected scannerState) {
		if actual := makeTransition(state, input); actual != expected {
			t.Errorf("%v, %v -> %v", state, input, actual)
		}
	}
	test(nonWord, nonWordRune, nonWord)
	test(nonWord, wordRune, inWord)
	test(nonWord, endOfInput, nonWord)
	test(inWord, nonWordRune, exitWord)
	test(inWord, wordRune, inWord)
	test(inWord, endOfInput, exitWord)
	test(exitWord, nonWordRune, nonWord)
	test(exitWord, wordRune, inWord)
	test(exitWord, endOfInput, nonWord)
}
