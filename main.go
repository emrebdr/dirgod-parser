package main

import (
	"fmt"

	"github.com/emrebdr/dirgod-parser/src/parser"
)

func main() {
	parser := parser.NewParser("./example/Dirgod")
	fmt.Println(parser.Parse().ToString())
}
