package main

import (
	"CompilerJuandi/lexer"
	"fmt"
	"os"
)

func main() {

	input, err := os.ReadFile("myLenguage.lang")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(input))

	myLexer := lexer.NewLexer(string(input))
	myLexer.StartLexer()

	fmt.Println("numero de tokens", len(myLexer.Tokens))
	for _, v := range myLexer.Tokens {
		fmt.Printf("Token: %v Val: %s \n", v.TokenType, v.Value)
	}

}
