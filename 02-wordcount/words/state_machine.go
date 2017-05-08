package words

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
	table[exitWord][wordRune] = inWord
	table[exitWord][endOfInput] = nonWord

	return table
}

var transitionTable = makeTransitionTable()

func makeTransition(state scannerState, cat inputCategory) scannerState {
	return transitionTable[state][cat]
}
