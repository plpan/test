package main

import (
	"fmt"
	"test/test_modern/parser"

	"github.com/modern-go/parse"
)

func main() {
	src := parse.NewSourceString(`4 /(1 +1 )+ 2`)
	parsed := parse.Parse(src, parser.NewExprLexer(), 0)
	fmt.Println(parsed)
}
