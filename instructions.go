package brainfuck

type InstructionType string

var (
	moveLeft  InstructionType = "<"
	moveRight InstructionType = ">"
	increment InstructionType = "+"
	decrement InstructionType = "-"
	openLoop  InstructionType = "["
	closeLoop InstructionType = "]"
	input     InstructionType = ","
	output    InstructionType = "."
)

// DefaultOperationHandlers are set of commands and corresponding actions that are executed
// by default when brainfuck program starts.
var DefaultOperationHandlers = map[InstructionType]func(*Brainfuck){
	moveLeft:  (*Brainfuck).MoveLeft,
	moveRight: (*Brainfuck).MoveRight,
	increment: (*Brainfuck).Increment,
	decrement: (*Brainfuck).Decrement,
	openLoop:  (*Brainfuck).OpenLoop,
	closeLoop: (*Brainfuck).CloseLoop,
	input:     (*Brainfuck).Read,
	output:    (*Brainfuck).Write,
}

// AddInstruction adds given instruction and operation to be perform into the brainfuck program.
func (bf *Brainfuck) AddInstruction(command InstructionType, handler func(*Brainfuck)) {
	bf.CmdOperationMapping[command] = handler
}

// RemoveInstruction deletes given command from the brainfuck program.
func (bf *Brainfuck) RemoveInstruction(command InstructionType) {
	delete(bf.CmdOperationMapping, command)
}
