package main

import (
	"fmt"

	"github.com/emrebdr/dirgod-parser/src/parser"
)

func main() {
	parser, err := parser.NewParser("./example/Dirgod")
	if err != nil {
		panic(err)
	}

	isValid := parser.IsValid()
	fmt.Println(isValid)
}
