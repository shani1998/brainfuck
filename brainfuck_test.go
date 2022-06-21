package brainfuck

import (
	"bytes"
	"github.com/shani1998/data-structures/stack"
	"io"
	"reflect"
	"testing"
)

func TestBrainfuck_CloseLoop(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.CloseLoop()
		})
	}
}

func TestBrainfuck_Decrement(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.Decrement()
		})
	}
}

func TestBrainfuck_Increment(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.Increment()
		})
	}
}

func TestBrainfuck_MoveLeft(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.MoveLeft()
		})
	}
}

func TestBrainfuck_MoveRight(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.MoveRight()
		})
	}
}

func TestBrainfuck_OpenLoop(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.OpenLoop()
		})
	}
}

func TestBrainfuck_Read(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.Read()
		})
	}
}

func TestBrainfuck_Run(t *testing.T) {
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
		codeReader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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
			if err := bf.Run(tt.args.codeReader); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrainfuck_Write(t *testing.T) {
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
	tests := []struct {
		name   string
		fields fields
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
			bf.Write()
		})
	}
}

func TestNewBrainFuck(t *testing.T) {
	type args struct {
		i io.Reader
	}
	tests := []struct {
		name  string
		args  args
		wantW string
		want  *Brainfuck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			got := NewBrainFuck(tt.args.i, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewBrainFuck() gotW = %v, want %v", gotW, tt.wantW)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrainFuck() = %v, want %v", got, tt.want)
			}
		})
	}
}
