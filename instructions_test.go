package brainfuck

import (
	"os"
	"testing"
)

func TestBrainfuck_AddInstruction(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	var TestHandler func(*Brainfuck)
	testBrainFuck.AddInstruction("*", TestHandler)
	if _, ok := testBrainFuck.CmdOperationMapping["*"]; !ok {
		t.Error("Failed to add new instruction")
	}
}

func TestBrainfuck_RemoveInstruction(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	testBrainFuck.RemoveInstruction("+")
	if _, ok := testBrainFuck.CmdOperationMapping["+"]; ok {
		t.Error("Failed to remove existing instruction")
	}
}
