package main

import (
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
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Unexpected String Literal\n"
	functionMessage += `Unexpected String Literal '` + token + `' at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
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
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Token was Expected\n"
	functionMessage += `Expected Token '` + expectedToken + `' at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
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
func lexError(line string, lineNumber int, i int) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Tokenization Error\n"
	functionMessage += `Unknown token: "` + string(line[i]) + `" at line number ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += line + "\n"
	for j := 0; j < i; j++ {
		functionMessage += " "
	}
	functionMessage += "👆\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
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
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += "💀 Runtime Error\n"
	functionMessage += errorMessage + ` at line ` + strconv.Itoa(lineNumber) + "\n"
	functionMessage += "👉 " + line + "\n"
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

/*
A Function which prints the file access error when executing the program.

It exits the program with a code of 1.

Parameters:
	- errorMessage: The error message which should be printed.
	- err: The error which was thrown.
*/
func printFileAccessError(errorMessage string, err string) {
	functionMessage := ""
	functionMessage += "\n"
	functionMessage += errorMessage + "\n"
	functionMessage += "Error: " + err
	functionMessage += "\n"
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}

/*
A function which will throw an error if no command line arguments was passed
*/
func noArguments() {
	functionMessage := ""
	functionMessage += "\n💀 No file name provided. Please provide a file name."
	functionMessage += "Usage: infant <filename>.infant"
	functionMessage += "\n💀 Exiting..."
	panicMessage += functionMessage
	panic("💀 ERROR 💀")
}
