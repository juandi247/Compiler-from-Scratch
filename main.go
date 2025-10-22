package main

import (
	"fmt"
	"os"
)

func main() {

	coso, err := os.ReadFile("myLenguage.lang")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(coso))

}
