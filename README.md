# brainfuck

Package `brainfuck` implements a [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter in Go.

Default brainfuck interpreter implements 8 commands defined here [wiki](https://en.wikipedia.org/wiki/Brainfuck#Commands)
but this library also provides a mechanism to `add`/`remove` commands.
##### Add command:
```
brainfuck.AddInstruction(<instructionType>, <Handler to be executed>)
```
##### Remove command:
```
brainfuck.RemoveInstruction(<instructionType>)
```
#### How to use

```sh
$ go get github.com/shani1998/brainfuck
```

	// create new io.Reader from inputs
	code := strings.NewReader("++[>++.<-]>+")
	
	// Standards interface to io
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	
	// initialize the brainfuck interpreter
	bf := brainfuck.NewInterpreter(input, output)
	
    // Store the result in output interface 
	_ = bf.Run(code)
	
	// print the result 
	fmt.Println (output.String())

#### How to run tests

To run all tests in the root of the project
```sh
$ make test
```
To run cli that uses brainfuck library
```sh
$ make run-cli
```

#### 