package main

import (
    "fmt"
    //"github.com/sumory/idgen"
	"github.com/plpan/test/test_glide_hierarchy"
	"github.com/fatih/color"
)

func main() {
	fmt.Println(test_glide_hierarchy.NextId())

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")
}
