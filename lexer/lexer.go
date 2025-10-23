package lexer

import (
	"fmt"
	//"time"
	"unicode"
)

/*
what do we need for a lexer?
Tokens: identifiers of our lenguage or keywords
Item: kind of token that we have? "><= != ==, etc
Lexer: that reads the tokens and saves them into, an array? could be.
*/
type item int

const (
	EOF item = iota
	IDENTIFIER
	NUMBER
	STRING
	LEFT_BRACKET  //(
	RIGHT_BRACKET //)

	LEFT_CURL
	RIGHT_CURL

	IF
	ELSE

	ASSIGMENT // =
	EQUALS
	NOT_EQUAL

	BIGGER
	SMALLER

	BIGGER_OR_EQUAL
	SMALLER_OR_EQUAL

	OR
	AND

	SEMICOLON

	SUM
	SUBSTRACT
	MULTIPLY
	DIVIDE
	MODULO

	PRINT
	FUNCTION
	RETURN
	WHILE
	FOR
	VAR
)

var keyWords = map[string]item{
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"while":  WHILE,
	"fn":     FUNCTION,
	"return": RETURN,
	"print":  PRINT,
	"and":    AND,
	"or":     OR,
	"var":    VAR,
}

type token struct {
	TokenType item
	Value     string
}

type lexer struct {
	input  string //this would be all the input from the file or files
	start  int    //mantain the start of the curr token so we can save it
	curr   int    //this is to traverse the current token
	Tokens []token
}

func NewLexer(input string) *lexer {
	return &lexer{
		input:  input,
		start:  0,
		curr:   0,
		Tokens: []token{}}
}

func (l *lexer) saveToken(tokeType item, val string) {

	l.Tokens = append(l.Tokens, token{
		TokenType: tokeType, Value: val})
	l.curr++
	fmt.Printf("Avanzamos a: %v , length: %v", l.curr, len(l.input))
}

func (l *lexer) StartLexer() {
	/*
				how?
				Evaluate the curr character. if our curr character is an special character like *,+,(,), etc, we emit the token
				IF our value is not there, it will be probably an string or a number.
				We need to check the strings that have the " at the start so it will be an evaluation there too"
				If not we check the nomral value and we save it as if its an "if", "else", "for", "whilw", "var", etc we append it,
				if its not that we append the value as an IDENTIFIER because its a variable name or something
				if we reach an EOF then we finish.

		error:
		if we find a string like "hooh, without an ending ", then we can show an error message after, this error will be taken by thte lexer

	*/
	for !l.isAtEnd() {
		character := l.input[l.curr]
		l.start = l.curr
		switch character {
		case '(':
			l.saveToken(LEFT_BRACKET, "(")
		case ')':
			l.saveToken(RIGHT_BRACKET, ")")
		case '{':
			l.saveToken(LEFT_CURL, "{")
		case '}':
			l.saveToken(RIGHT_CURL, "}")
		case '*':
			l.saveToken(MULTIPLY, "*")
		case '+':
			l.saveToken(SUM, "+")
		case '-':
			l.saveToken(SUBSTRACT, "-")
		case '%':
			l.saveToken(MODULO, "%")
		case '/':
			l.saveToken(DIVIDE, "/")
		case ';':
			l.saveToken(SEMICOLON, ";")

		//this tokens need to be evaluated into a different function
		//because it can be assignment, or >= or == or =, so it depends.
		case '=':
			l.readEqualOrAsignment()
		case '>':
			l.readBiggerOrBiggerEqual()
		case '<':
			l.readSmallerOrSmallerEqual()
		case '!':
			l.readNotEqual()

		case '"':
			l.readString()
		//okay if it wasnt one of those identifiers, means that we reached a string or a number
		//we could check here even for unvalid like @ or Runes in go but for now meh
		default:
			if unicode.IsNumber(rune(character)) {
				l.readNumber()
			} else if unicode.IsLetter(rune(character)) {
				l.readIdentifier()
			} else if character == ' ' || character == '\r' || character == '\n' {
				l.curr++
			}
		}

	}

	fmt.Println("llegamos al final del archivo")
}
