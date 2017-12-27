package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.MustCompile("\\\033\\[[0-9]+m").Split("\033[32mdogo version 0.0.3 installed successfully \033[0m", 3)[1])
}
