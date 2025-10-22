package lexer

import "unicode"

//error hanlding for EOF
func (l *lexer) isAtEnd() bool {
	if l.curr >= len(l.input) {
		return true
	}
	return false

}

//Helper functions for characters on the input
func (l *lexer) readEqualOrAsignment() {
	if l.input[l.curr+1] == '=' {
		l.saveToken(EQUALS, "==")
		return
	}
	l.saveToken(ASSIGMENT, "=")
}

func (l *lexer) readBiggerOrBiggerEqual() {
	if l.input[l.curr+1] == '=' {
		l.saveToken(BIGGER_OR_EQUAL, ">=")
		return
	}
	l.saveToken(BIGGER, ">")
}

func (l *lexer) readSmallerOrSmallerEqual() {
	if l.input[l.curr+1] == '=' {
		l.saveToken(SMALLER_OR_EQUAL, "<=")
		return
	}
	l.saveToken(BIGGER, "<")
}

func (l *lexer) readNotEqual() {
	if l.input[l.curr+1] == '=' {
		l.saveToken(NOT_EQUAL, "!=")
		return
		//should be an error here
	}
}

func (l *lexer) readString() {
	l.curr++ //so we avoid having that initial stoping becasue curr was the initial "
	for l.input[l.curr] != '"' {
		l.curr++
	}

	l.saveToken(STRING, string(l.input[l.start:(l.curr+1)]))
}

func (l *lexer) readNumber() {
	for unicode.IsNumber(rune(l.input[l.curr])) {
		l.curr++
	}
	l.curr--
	//we let the function save token to advance +1
	l.saveToken(NUMBER, string(l.input[l.start:l.curr+1]))

}

func (l *lexer) readIdentifier() {
	for unicode.IsLetter(rune(l.input[l.curr])) || unicode.IsNumber(rune(l.input[l.curr])) {
		l.curr++
	}

	word := string(l.input[l.start:l.curr])
	itemType := checkKeyWord(word)

	l.curr--
	l.saveToken(itemType, word)
}

func checkKeyWord(word string) item {

	itemFound, exists := keyWords[word]

	if !exists {
		//its a variable name or smth
		return IDENTIFIER
	}
	//its an if, else, while, etc
	return itemFound
}
