package lexer

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
	for l.input[l.curr] != '"' {
		l.curr++
	}

	l.saveToken(STRING, string(l.input[l.start:(l.curr+1)]))
}
