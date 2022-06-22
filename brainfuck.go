package brainfuck

import (
	"bytes"
	"io"
	"log"
	"math"

	"github.com/shani1998/data-structures/stack"
)

//Max memory limit. Array of byte type simulating memory of max 255 bits from 0 to 255.
const size = math.MaxUint8

// Brainfuck used to implement the brainfuck algorithm
type Brainfuck struct {
	// the Code to be executed
	Code string
	// the tape Code executes on
	instrMemory [size]byte
	// the position of the instruction pointer
	instrPointer uint8
	// the position of the data pointer
	dataPointer uint8
	// the stack used to store the positions of [ in Code
	loopStack stack.Stack
	// the output interface to write on
	Output io.Writer
	// the input interface to read from
	Input io.Reader
	// the buffer from which or to read and write
	buffer []byte
	// the command and corresponding func to be executed
	CmdOperationMapping map[InstructionType]func(*Brainfuck)
}

// NewInterpreter initializes a brainfuck interpreter with given reader and a writer interface
// and default operator handler, CmdOperationMapping could be modified by using Add/Remove Instruction
func NewInterpreter(i io.Reader, w io.Writer) *Brainfuck {
	return &Brainfuck{
		Code:                "",
		instrMemory:         [size]byte{},
		instrPointer:        0,
		dataPointer:         0,
		loopStack:           stack.Stack{},
		buffer:              make([]byte, 1),
		Output:              w,
		Input:               i,
		CmdOperationMapping: DefaultOperationHandlers,
	}
}

// MoveLeft represents the Brainfuck instruction "<". It moves the pointer
// left by one cell, wrapping to the end of the tape if necessary
func (bf *Brainfuck) MoveLeft() {
	//If the pointer reaches zero
	if bf.dataPointer == 0 {
		//pointer is returned to rightmost memory position
		bf.dataPointer = size - 1
	} else {
		bf.dataPointer--
	}
}

// MoveRight represents the Brainfuck instruction ">". It moves the pointer
// right by one cell, wrapping to 0 if necessary
func (bf *Brainfuck) MoveRight() {
	//If memory is full
	if bf.dataPointer == size-1 {
		//pointer is returned to zero
		bf.dataPointer = 0
	} else {
		bf.dataPointer++
	}
}

// Increment represents the Brainfuck instruction "+". It increments the memory
// cell under the pointer
func (bf *Brainfuck) Increment() {
	bf.instrMemory[bf.dataPointer]++
}

// Decrement represents the Brainfuck instruction "-". It decrements the memory
// cell under the pointer
func (bf *Brainfuck) Decrement() {
	bf.instrMemory[bf.dataPointer]--
}

// Output represents the Brainfuck instruction ".". It prints the character
// value of the cell under the pointer
func (bf *Brainfuck) Write() {
	bf.buffer[0] = bf.instrMemory[bf.dataPointer]
	n, err := bf.Output.Write(bf.buffer)
	if err != nil {
		log.Printf("error while writing data to Output: %v", err)
	}
	if n != 1 {
		log.Print("wrong num bytes written")
	}
}

// Input represents the Brainfuck instruction ",". It takes a character from
// input and stores it in the cell under the pointer
func (bf *Brainfuck) Read() {
	n, err := bf.Input.Read(bf.buffer)
	if err != nil {
		log.Printf("error while reading data from input: %s", err)
	}
	if n != 1 {
		log.Printf("wrong num bytes read")
	}

	bf.instrMemory[bf.dataPointer] = bf.buffer[0]
}

// OpenLoop represents the Brainfuck instruction "[". It forms the opening
// part of a loop. If the cell under the pointer is 0, indicate to jump to the next ]
func (bf *Brainfuck) OpenLoop() {
	bf.loopStack.Push(bf.instrPointer)
	// return if the memory byte at data pointer is non-zero
	if bf.instrMemory[bf.dataPointer] != 0 {
		return
	}
	var jumpStack stack.Stack
	jumpStack.Push(1)          // push one from '[' braces
	for !jumpStack.IsEmpty() { // jump to the next ']'
		bf.instrPointer++
		// to avoid not getting closing bracket in the range
		if bf.instrPointer >= uint8(len(bf.Code)) {
			log.Printf("unable to read the Code, please set valid code")
			return
		}
		// push into stack for every opening bracket
		if string(bf.Code[bf.instrPointer]) == "[" {
			jumpStack.Push(1)
		}
		// pop from stack for every closing bracket
		if string(bf.Code[bf.instrPointer]) == "]" {
			jumpStack.Pop()
		}
	}
}

// CloseLoop represents the Brainfuck instruction "]".
// It closes a loop and returns the index of the matching open brace in the
// Code, if the cell under the pointer is not 0.
func (bf *Brainfuck) CloseLoop() {
	p, err := bf.loopStack.Pop()
	if err != nil {
		log.Printf("error occured while poping data from stack: %v", err)
		return
	}
	if bf.instrMemory[bf.dataPointer] != 0 {
		bf.instrPointer = p.(uint8) - 1
	}
}

// Run runs the given Brainfuck program with given instructions.
func (bf *Brainfuck) Run(CodeReader io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(CodeReader)
	if err != nil {
		return err
	}
	bf.Code = buf.String()

	// traverse over each instruction
	for bf.instrPointer < uint8(len(bf.Code)) {
		if oprHandler, ok := bf.CmdOperationMapping[InstructionType(bf.Code[bf.instrPointer])]; ok {
			oprHandler(bf)
		}
		bf.instrPointer++
	}
	return nil
}
