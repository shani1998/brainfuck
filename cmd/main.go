package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shani1998/brainfuck"
)

func main() {
	Program := strings.NewReader("++[>++.<-]>+.,>>.") //
	// Standards interface to io
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)

	brainfuck := brainfuck.NewBrainFuck(input, output)
	fmt.Println(Program)
	err := brainfuck.Run(Program)
	if err != nil {
		fmt.Printf("File %s does not contain a valid Brainfuck program:\n", err)
		return
	}
}
