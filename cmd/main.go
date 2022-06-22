package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shani1998/brainfuck"
)

var (
	defaultCode = ",[.,]"
)

func init() {
	flag.StringVar(&defaultCode, "code", defaultCode, "Instruction to be executed by brainfuck interpreter")
}

func main() {
	flag.Parse()
	fmt.Println(defaultCode)
	code := strings.NewReader(defaultCode)
	bf := brainfuck.NewInterpreter(os.Stdin, os.Stdout)
	err := bf.Run(code)
	if err != nil {
		log.Fatalf("error while execting instructions . %v", err)
	}
}
