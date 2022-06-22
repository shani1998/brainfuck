package brainfuck

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestBrainfuck_Decrement(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		want byte
	}{
		{
			name: "Decrement m/m byte having value zero",
			want: 255, // decrement from zero should go to max of uint8
		},
		{
			name: "Decrement m/m byte having value non zero",
			want: 254, // it was set to 255 in previous test
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBrainFuck.Decrement()
			if tt.want != testBrainFuck.instrMemory[testBrainFuck.dataPointer] {
				t.Errorf("Decrement(): expected %v, got %v", tt.want, testBrainFuck.instrMemory[testBrainFuck.dataPointer])
			}
		})
	}
}

func TestBrainfuck_Increment(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		want byte
	}{
		{
			name: "Increment m/m byte having value zero",
			want: 1, // default set to zero
		},
		{
			name: "Increment m/m byte having value non zero",
			want: 2, // it was set to 1 in previous test
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBrainFuck.Increment()
			if tt.want != testBrainFuck.instrMemory[testBrainFuck.dataPointer] {
				t.Errorf("Increment(): expected %v, got %v", tt.want, testBrainFuck.instrMemory[testBrainFuck.dataPointer])
			}
		})
	}
}

func TestBrainfuck_MoveLeft(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		want byte
	}{
		{
			name: "MoveLeft data pointer having value zero",
			want: size - 1,
		},
		{
			name: "MoveLeft data pointer having value non zero",
			want: size - 2, // it was set to size-1 in previous test
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBrainFuck.MoveLeft()
			if tt.want != testBrainFuck.dataPointer {
				t.Errorf("MoveLeft(): expected %v, got %v", tt.want, testBrainFuck.dataPointer)
			}
		})
	}
}

func TestBrainfuck_MoveRight(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		want byte
	}{
		{
			name: "MoveRight data pointer having value zero",
			want: 1,
		},
		{
			name: "MoveRight data pointer having value non zero",
			want: 2, //it was set to 1 in previous test
		},
		{
			name: "MoveRight data pointer having value max of uint8",
			want: 0, //it was set to 1 in previous test
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "max of uint8") {
				testBrainFuck.dataPointer = size - 1
			}
			testBrainFuck.MoveRight()
			if tt.want != testBrainFuck.dataPointer {
				t.Errorf("MoveRight(): expected %v, got %v", tt.want, testBrainFuck.dataPointer)
			}
		})
	}
}

func TestBrainfuck_OpenLoop(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		code string
		want uint8
	}{
		{
			name: "invalid code, Instruction pointer reaches at the end of the Code",
			code: "random code",
			want: uint8(len("random code")), // pointer should reach at end
		},
		{
			name: "Set a valid code",
			code: "[>++.<-]>+",
			want: uint8(len("++[>++.<-]>+")),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBrainFuck.Code = tt.code
			if strings.Contains(tt.name, "non zero") {
				testBrainFuck.instrMemory[testBrainFuck.dataPointer]++
			}
			testBrainFuck.OpenLoop()
			if i+1 != len(testBrainFuck.loopStack) {
				t.Errorf("mismatched coount for total open braces")
			}
			if tt.want != testBrainFuck.instrPointer {
				t.Errorf("MoveRight(): expected %v, got %v", tt.want, testBrainFuck.instrPointer)
			}
		})
	}
}

func TestBrainfuck_CloseLoop(t *testing.T) {
	testBrainFuck := NewInterpreter(os.Stdin, os.Stdout)
	tests := []struct {
		name string
		code string
		want int
	}{
		{
			name: "loop stack having length zero",
			want: 0, // initially loopStack will be empty, so pop will error
		},
		{
			name: "set loop stack with zero m/m byte",
			want: 0, // after pop loop stack length should be empty
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testBrainFuck.Code = tt.code
			if strings.Contains(tt.name, "set loop") {
				testBrainFuck.loopStack.Push(1)
			}
			testBrainFuck.CloseLoop()
			if tt.want != len(testBrainFuck.loopStack) {
				t.Errorf("MoveRight(): expected %v, got %v", tt.want, len(testBrainFuck.loopStack))
			}
		})
	}
}

func TestBrainfuck_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   string
		output string
	}{
		{
			name:   "",
			code:   "----[---->+<]>++.+.+.+.",
			output: "ABCD",
		},
		{
			name:   "",
			code:   "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.",
			output: "Hello World!\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codeReader := strings.NewReader(tt.code)
			i := new(bytes.Buffer)
			o := new(bytes.Buffer)
			bfm := NewInterpreter(i, o)
			_ = bfm.Run(codeReader)
			if o.String() != tt.output {
				t.Errorf("wrong output, got %s", o.String())
			}
		})
	}
}
