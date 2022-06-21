package brainfuck

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/shani1998/data-structures/stack"
)

//Max memory limit. Array of byte type simulating memory of max 255 bits from 0 to 255.
const size = math.MaxUint8

// Brainfuck used to implement the brainfuck algorithm
type Brainfuck struct {
	// the code to be executed
	code string
	// the tape code executes on
	InstrMemory [size]byte
	// the position of the instruction pointer
	InstrPointer uint8
	// the position of the data pointer
	DataPointer uint8
	// the stack used to store the positions of [ in code
	LoopStack stack.Stack
	// the output interface to write on
	Output io.Writer
	// the input interface to read from
	Input io.Reader
	// the buffer from which or to read and write
	buffer []byte
	// the command and corresponding func to be executed
	CmdOperationMapping map[InstructionType]func(*Brainfuck)
}

// NewBrainFuck initializes a brainfuck interpreter with given reader and a writer interface
// and default operator handler, CmdOperationMapping could be modified by using Add/Remove Instruction
func NewBrainFuck(i io.Reader, w io.Writer) *Brainfuck {
	return &Brainfuck{
		code:                "",
		InstrMemory:         [size]byte{},
		InstrPointer:        0,
		DataPointer:         0,
		LoopStack:           stack.Stack{},
		Output:              w,
		Input:               i,
		buffer:              make([]byte, 1),
		CmdOperationMapping: DefaultOperationHandlers,
	}
}

// MoveLeft represents the Brainfuck instruction "<". It moves the pointer
// left by one cell, wrapping to the end of the tape if necessary
func (bf *Brainfuck) MoveLeft() {
	//If the pointer reaches zero
	if bf.DataPointer == 0 {
		//pointer is returned to rightmost memory position
		bf.DataPointer = size - 1
	} else {
		bf.DataPointer--
	}
}

// MoveRight represents the Brainfuck instruction ">". It moves the pointer
// right by one cell, wrapping to 0 if necessary
func (bf *Brainfuck) MoveRight() {
	//If memory is full
	if bf.DataPointer == size-1 {
		//pointer is returned to zero
		bf.DataPointer = 0
	} else {
		bf.DataPointer++
	}
}

// Increment represents the Brainfuck instruction "+". It increments the memory
// cell under the pointer
func (bf *Brainfuck) Increment() {
	bf.InstrMemory[bf.DataPointer]++
}

// Decrement represents the Brainfuck instruction "-". It decrements the memory
// cell under the pointer
func (bf *Brainfuck) Decrement() {
	bf.InstrMemory[bf.DataPointer]--
}

// Output represents the Brainfuck instruction ".". It prints the character
// value of the cell under the pointer
func (bf *Brainfuck) Write() {
	_, err := fmt.Fprint(bf.Output, bf.InstrMemory[bf.DataPointer])
	if err != nil {
		log.Printf("error while writing data to Output: %v", err)
		return
	}
}

// Input represents the Brainfuck instruction ",". It takes a character from
// input and stores it in the cell under the pointer
func (bf *Brainfuck) Read() {
	_, err := fmt.Fscan(bf.Input, &bf.InstrMemory[bf.DataPointer])
	if err != nil {
		log.Printf("error while reading data from input: %s", err)
		return
	}
}

// OpenLoop represents the Brainfuck instruction "[". It forms the opening
// part of a loop. If the cell under the pointer is 0, returns true to
// indicate to jump to the next ]
func (bf *Brainfuck) OpenLoop() {
	bf.LoopStack.Push(bf.InstrPointer)

	var jumpStack stack.Stack
	jumpStack.Push(1)
	for !jumpStack.IsEmpty() || bf.InstrMemory[bf.DataPointer] == 0 { // jump to the next ']'
		fmt.Printf(" in OpenLoop index: %d, char: %c, bf.dp:%+v, m/m:%v jumpStack %v,loopStack %v\n",
			bf.InstrPointer, bf.code[bf.InstrPointer], bf.DataPointer, bf.InstrMemory[bf.DataPointer], jumpStack, bf.LoopStack)
		bf.InstrPointer += 1
		if string(bf.code[bf.InstrPointer]) == "[" {
			jumpStack.Push(1)
		} else if string(bf.code[bf.InstrPointer]) == "]" {
			jumpStack.Pop()
		}
	}
}

// CloseLoop represents the Brainfuck instruction "]".
// It closes a loop and returns the index of the matching open brace in the
// code, if the cell under the pointer is not 0. Otherwise returns -1
func (bf *Brainfuck) CloseLoop() {
	p, err := bf.LoopStack.Pop()
	if err != nil {
		panic(err)
	}
	if bf.InstrMemory[bf.DataPointer] != 0 {
		bf.InstrPointer = p.(uint8) - 1
	}
}

// Run runs the given Brainfuck program with given instructions.
func (bf *Brainfuck) Run(codeReader io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(codeReader)
	if err != nil {
		return err
	}
	bf.code = buf.String()

	// traverse over each instruction
	for i, cmd := range bf.code {
		bf.InstrPointer = uint8(i)
		if oprHandler, ok := bf.CmdOperationMapping[InstructionType(cmd)]; ok {
			fmt.Printf("in Run index: %d, char: %c, bf.dp:%+v, m/m:%v loopStack %v\n",
				bf.InstrPointer, bf.code[bf.InstrPointer], bf.DataPointer, bf.InstrMemory[bf.DataPointer], bf.LoopStack)
			oprHandler(bf)
		}
	}
	return nil
}
