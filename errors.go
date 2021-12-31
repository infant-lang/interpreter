package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Function which prints the 'Unexpected String Literal' error message

Parameters:
	- line: the line of code which caused the error
	- lineNumber: the line of the code which contains the error
	- token: the token which caused the error

Returns: Nothing

Prints: The error to the console.
*/
func printParseError(line string, lineNumber int, token string) {
	fmt.Println()
	fmt.Println("💀 Unexpected String Literal")
	fmt.Println(`Unexpected String Literal '` + token + `' at line number ` + strconv.Itoa(lineNumber))
	fmt.Println("👉 " + line)
	fmt.Println()
	os.Exit(1)
}

/*
Function which prints the 'Expected Token' error message

Parameters: 
	- line: the line of code which caused the error
	- lineNumber: the line of the code which contains the error
	- token: the expectedToken

Returns: Nothing

Prints: The error to the console.
*/
func printExpectedTokenError(line string, lineNumber int, expectedToken string) {
	fmt.Println()
	fmt.Println("💀 Token was Expected")
	fmt.Println(`Expected  '` + expectedToken + `' at line number ` + strconv.Itoa(lineNumber))
	fmt.Println("👉 " + line)
	fmt.Println()
	os.Exit(1)
}

/*
The function which prints the error to the console.

Parameters:
	- line: The line which contains the error.
	- lineNumber: The line number of the error.
	- i: The index of the error at the given line

	Returns: Nothing

	Prints: The error to the console.
*/
func printError(line string, lineNumber int, i int) {
	fmt.Println()
	fmt.Println("💀 Tokenization Error:")
	fmt.Println(`Unknown token: "` + string(line[i]) + `"` + " at line " + strconv.Itoa(lineNumber))
	fmt.Println(line)
	for j := 0; j < i; j++ {
		fmt.Print(" ")
	}
	fmt.Print("👆")
	fmt.Println()
	os.Exit(1)
}

/*
A function which prints the runtime error when executing the program.

Parameters:
	- line: The line which contains the error.
	- lineNumber: The line number of the error.
	- errorMessage: The error which was thrown.

Returns: Nothing

Prints: The error to the console.
*/
func runtimeError(line string, lineNumber int, errorMessage string) {
	fmt.Println()
	fmt.Println("💀 Runtime Error:")
	fmt.Println(errorMessage + " at line " + strconv.Itoa(lineNumber))
	fmt.Println("👉 " + line)
	fmt.Println()
	os.Exit(1)
}

/*
A Function which prints the file access error when executing the program.

It exits the program with a code of 1.

Parameters: 
	- errorMessage: The error message which should be printed.
	- err: The error which was thrown.
*/
func printFileAccessError(errorMessage string, err string) {
	fmt.Println()
	fmt.Println(errorMessage)
	fmt.Println("Error:", err)
	fmt.Println()
	os.Exit(1)
}

/*
A function which will throw an error if no command line arguments was passed
*/
func noArguments() {
	fmt.Println("\n💀 No file name provided. Please provide a file name.")
	fmt.Println("Usage: infant <filename>.infant")
	fmt.Println("\n💀 Exiting...")
	os.Exit(1)
}
