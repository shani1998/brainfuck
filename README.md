# brainfuck

Package `brainfuck` implements a brainfuck interpreter.


#### How to use

	// create new io.Reader from inputs
	code := strings.NewReader("++[>++.<-]>+")
	
	// Standards interface to io
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	
	// initialize the machine
	bf := brainfuck.NewBrainfuck(input, output)
	
    // Store the result in output interface 
	_ = bfm.Run(code)
	
	// print the result 
	fmt.Println (output.String())

#### How to run tests

To run all tests in the root of the project run `go test ./...` command.

#### 