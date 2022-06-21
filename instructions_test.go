package brainfuck

import (
	"github.com/shani1998/data-structures/stack"
	"io"
	"testing"
)

func TestBrainfuck_AddInstruction(t *testing.T) {
	type fields struct {
		code                string
		InstrMemory         [size]byte
		InstrPointer        uint8
		DataPointer         uint8
		LoopStack           stack.Stack
		Output              io.Writer
		Input               io.Reader
		buffer              []byte
		CmdOperationMapping map[InstructionType]func(*Brainfuck)
	}
	type args struct {
		command InstructionType
		handler func(*Brainfuck)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf := &Brainfuck{
				code:                tt.fields.code,
				InstrMemory:         tt.fields.InstrMemory,
				InstrPointer:        tt.fields.InstrPointer,
				DataPointer:         tt.fields.DataPointer,
				LoopStack:           tt.fields.LoopStack,
				Output:              tt.fields.Output,
				Input:               tt.fields.Input,
				buffer:              tt.fields.buffer,
				CmdOperationMapping: tt.fields.CmdOperationMapping,
			}
			bf.AddInstruction(tt.args.command, tt.args.handler)
		})
	}
}

func TestBrainfuck_RemoveInstruction(t *testing.T) {
	type fields struct {
		code                string
		InstrMemory         [size]byte
		InstrPointer        uint8
		DataPointer         uint8
		LoopStack           stack.Stack
		Output              io.Writer
		Input               io.Reader
		buffer              []byte
		CmdOperationMapping map[InstructionType]func(*Brainfuck)
	}
	type args struct {
		command InstructionType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf := &Brainfuck{
				code:                tt.fields.code,
				InstrMemory:         tt.fields.InstrMemory,
				InstrPointer:        tt.fields.InstrPointer,
				DataPointer:         tt.fields.DataPointer,
				LoopStack:           tt.fields.LoopStack,
				Output:              tt.fields.Output,
				Input:               tt.fields.Input,
				buffer:              tt.fields.buffer,
				CmdOperationMapping: tt.fields.CmdOperationMapping,
			}
			bf.RemoveInstruction(tt.args.command)
		})
	}
}
