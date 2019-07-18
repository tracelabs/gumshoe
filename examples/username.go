package main

import (
	"github.com/tracelabs/gumshoe"
	"github.com/tracelabs/gumshoe/findings"
)

func main() {
	input := &findings.GenericUsername{Name: "adrianosela"}
	output := gumshoe.Run(input)
	for _, f := range output {
		f.Display()
	}
}
