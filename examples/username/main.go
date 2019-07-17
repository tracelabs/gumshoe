package main

import (
	"github.com/tracelabs/gumshoe/finding"
	"github.com/tracelabs/gumshoe/investigation"
)

func main() {
	input := &finding.Username{Name: "your-username"}
	output := investigation.Run(input)
	for _, f := range output {
		f.Display()
	}
}
