package main

import (
	"bufio"
	"os"
	"strconv"
)

/*
A function that reads the file and returns a slice of strings

Parameters:
	- args: The Number of Command Line Argument to be read.

Returns:
	- fileTextLines: A slice of strings.
*/
func readFile(args int) ([]string) {

	if len(os.Args) == 1 {
		noArguments()
	}
	
	FILE_NAME := os.Args[1]

	if FILE_NAME == "-v" || FILE_NAME == "--version" {
		message = "v1.0.3"
		panic(message)
	}

	readFile, err := os.Open(FILE_NAME)

	if err != nil {
		errorMessage := "💀 Unable to Access File: " + FILE_NAME
		printFileAccessError(errorMessage, err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
 
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	return fileTextLines
}

/*
A function to check if the following grammars are valid
	- Grammar 1: MEMORY ASSIGNMENT MEMORY ARITHMETIC MEMORY
	- Grammar 2: MEMORY ASSIGNMENT MEMORY ARITHMETIC POINTER
	- Grammar 3: MEMORY ASSIGNMENT MEMORY ARITHMETIC NUMBER

	- Grammar 4: MEMORY ASSIGNMENT POINTER ARITHMETIC MEMORY
	- Grammar 5: MEMORY ASSIGNMENT POINTER ARITHMETIC POINTER
	- Grammar 6: MEMORY ASSIGNMENT POINTER ARITHMETIC NUMBER

	- Grammar 7: MEMORY ASSIGNMENT NUMBER ARITHMETIC MEMORY
	- Grammar 8: MEMORY ASSIGNMENT NUMBER ARITHMETIC POINTER
	- Grammar 9: MEMORY ASSIGNMENT NUMBER ARITHMETIC NUMBER

*/
func checkArithmetic(tokens []token, line string, lineNumber int) []token {

	if len(tokens) < 5 {
		parseError(line, lineNumber)
		panic("💀")
	}

	if tokens[0].tokenType == "MEMORY" {
		if tokens[1].tokenType == "ASSIGNMENT" {
			if tokens[2].tokenType == "MEMORY" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			} else if tokens[2].tokenType == "POINTER" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
			} else if tokens[2].tokenType == "NUMBER" {
				if tokens[3].tokenType == "ARITHMETIC" {
					if tokens[4].tokenType == "MEMORY" {
						return tokens[:5]
					} else if tokens[4].tokenType == "POINTER" {
						return tokens[:5]
					} else if tokens[4].tokenType == "NUMBER" {
						return tokens[:5]
					}
					printParseError(line, lineNumber, tokens[4].tokenValue)
				}
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

/*
A function to check if the following grammars are valid
	- Grammar 1: MEMORY ASSIGNMENT POINTER
	- Grammar 2: MEMORY ASSIGNMENT NUMBER
*/
func checkAssignment(tokens []token, line string, lineNumber int) []token {

	if len(tokens) != 3 {
		parseError(line, lineNumber)
		panic("💀")
	}

	if tokens[0].tokenType == "MEMORY" {
		if tokens[1].tokenType == "ASSIGNMENT" {
			if tokens[2].tokenType == "POINTER" {
				return tokens[:3]
			} else if tokens[2].tokenType == "NUMBER" {
				return tokens[:3]
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

/*
A function to check if the following grammars are valid
	- Grammar 1: PRINT MEMORY
	- Grammar 2: PRINT POINTER
	- Grammar 3: PRINT CHAR POINTER
	- Grammar 4: PRINT CHAR MEMORY
	- Grammar 5: PRINT NUMBER
	- Grammar 6: PRINT CHAR NUMBER
	- Grammar 7: PRINT SPECIAL
*/
func checkPrint(tokens []token, line string, lineNumber int) []token {

	if len(tokens) < 2 {
		parseError(line, lineNumber)
		panic("💀")
	}

	if tokens[0].tokenType == "PRINT" {
		if tokens[1].tokenType == "MEMORY" {
			return tokens[:2]
		} else if tokens[1].tokenType == "POINTER" {
			return tokens[:2]
		}

		if tokens[1].tokenType == "CHAR" {

			if len(tokens) != 3 {
				parseError(line, lineNumber)
				panic("💀")
			}

			if tokens[2].tokenType == "POINTER" {
				return tokens[:3]
			} else if tokens[2].tokenType == "MEMORY" {
				return tokens[:3]
			} else if tokens[2].tokenType == "NUMBER" {
				return tokens[:3]
			}

			printParseError(line, lineNumber, tokens[2].tokenValue)
		}

		if tokens[1].tokenType == "NUMBER" {
			return tokens[:2]
		}

		if tokens[1].tokenType == "SPECIAL" {
			return tokens[:2]
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}

	return nil
}

/*
A function to check if the following grammars are valid
	- Grammar 1: ACTION POINTER DIRECTION
	- Grammar 2: ACTION POINTER DIRECTION NUMBER
*/
func checkAction(tokens []token, line string, lineNumber int) []token {

	if len(tokens) < 3 {
		parseError(line, lineNumber)
		panic("💀")
	}

	if tokens[0].tokenType == "ACTION" {
		if tokens[2].tokenType == "DIRECTION" {
			if tokens[1].tokenType != "POINTER" {
				printParseError(line, lineNumber, tokens[1].tokenValue)
			}

			if len(tokens) == 4 {
				if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
			}

			return tokens[:3]

		}

		printParseError(line, lineNumber, tokens[2].tokenValue)
	}

	return nil
}

/*
A function to check if the following grammars are valid
	- Grammar 1: CONDITION MEMORY COMPARISON POINTER
	- Grammar 2: CONDITION MEMORY COMPARISON MEMORY
	- Grammar 3: CONDITION MEMORY COMPARISON NUMBER
	- Grammar 4: CONDITION POINTER COMPARISON POINTER
	- Grammar 5: CONDITION POINTER COMPARISON MEMORY
	- Grammar 6: CONDITION POINTER COMPARISON NUMBER
	- Grammar 7: CONDITION NUMBER COMPARISON POINTER
	- Grammar 8: CONDITION NUMBER COMPARISON MEMORY
	- Grammar 9: CONDITION NUMBER COMPARISON NUMBER
*/
func checkCondition(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "CONDITION" {
		if tokens[1].tokenType == "MEMORY" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		} else if tokens[1].tokenType == "POINTER" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		} else if tokens[1].tokenType == "NUMBER" {
			if tokens[2].tokenType == "COMPARISON" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[3].tokenValue)
			}
			printParseError(line, lineNumber, tokens[2].tokenValue)
		}
		printParseError(line, lineNumber, tokens[1].tokenValue)
	}
	return nil
}

/*
A function to check if the following grammars are valid
	- Grammar 1: LOOP POINTER
	- Grammar 2: LOOP MEMORY
	- Grammar 3: LOOP NUMBER
*/
func checkLoop(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "LOOP" {
		if tokens[1].tokenType == "POINTER" {
			return tokens[:2]
		} else if tokens[1].tokenType == "MEMORY" {
			return tokens[:2]
		} else if tokens[1].tokenType == "NUMBER" {
			return tokens[:2]
		}
	}
	printParseError(line, lineNumber, tokens[1].tokenValue)
	return nil
}

/*
A function to move the pointer in the specified direction by the specified number of steps which defaults to 1

Parameters:
	- actionTokens: the tokens array containing the action tokens
	- line: the line of the program execution
	- lineNumber: the line number of the program execution
	- p: the pointer value
	- m: the memory value

Return value:
	- p: the new pointer value
	- m: the new memory value
*/
func pointerMovements(actionTokens []token, line string, lineNumber int, p int, m int) (int, int) {
	numberOfTokens := len(actionTokens)
	if numberOfTokens == 4 {
		points := actionTokens[3].tokenValue
		pointsToMove, _ := strconv.Atoi(points)

		if actionTokens[2].tokenValue == "right" {
			p = p + pointsToMove
		} else if actionTokens[2].tokenValue == "left" {

			// DO NOT ALLOW P TO BE NEGATIVE
			if p-pointsToMove < 0 {
				runtimeError(line, lineNumber, "Pointer can't point to a negative box")
			} else {
				p = p - pointsToMove
			}
		}
	} else if numberOfTokens == 3 {
		if actionTokens[2].tokenValue == "right" {
			p = p + 1
		} else if actionTokens[2].tokenValue == "left" {
			// DO NOT ALLOW P TO BE NEGATIVE
			if p - 1 < 0 {
				runtimeError(line, lineNumber, "Pointer can't point to a negative box")
			} else {
				p = p - 1
			}
		}
	}

	return p, m
}

/*
Function which takes care of all the printing stuff from the program execution

Parameters:
	- printTokens: []tokens containing the tokens to be printed
	- p: int - the current pointer value
	- m: int - the current memory value

Return Values:
	- p: int - the new pointer value
	- m: int - the new memory value
*/
func printStuff(printTokens []token, p int, m int) (int, int) {

	functionMessage := ""

	if printTokens[1].tokenType == "MEMORY" {
		functionMessage += strconv.Itoa(m)
	} else if printTokens[1].tokenType == "POINTER" {
		functionMessage += strconv.Itoa(p)
	} else if printTokens[1].tokenType == "NUMBER" {
		functionMessage += printTokens[1].tokenValue
	} else if printTokens[1].tokenType == "CHAR" {
		if printTokens[2].tokenType == "POINTER" {
			functionMessage += returnASCII(p)
		} else if printTokens[2].tokenType == "MEMORY" {
			functionMessage += returnASCII(m)
		} else if printTokens[2].tokenType == "NUMBER" {
			points := printTokens[2].tokenValue
			number, _ := strconv.Atoi(points)
			functionMessage += returnASCII(number)
		}
	} else if printTokens[1].tokenType == "SPECIAL" {
		if printTokens[1].tokenValue == "new" {
			functionMessage += "\n"
		} else if printTokens[1].tokenValue == "tab" {
			functionMessage += "\t"
		} else if printTokens[1].tokenValue == "space" {
			functionMessage += " "
		}
	}

	message += functionMessage

	return p, m
}

/*
Function which takes care of the arithmetic operations and assigns the result to the memory

Parameters:
	- arithmeticTokens: []tokens containing the tokens to be worked on
	- line: the line of the program execution
	- lineNumber: the line number of the program execution
	- p: int - the current pointer value
	- m: int - the current memory value

Return Values:
	- p: int - the new pointer value
	- m: int - the new memory value
*/
func doArithmetic(arithmeticTokens []token, line string, lineNumber int, p int, m int) (int, int) {
	var firstOperand int
	var secondOperand int

	if arithmeticTokens[2].tokenType == "MEMORY" {
		firstOperand = m
	} else if arithmeticTokens[2].tokenType == "POINTER" {
		firstOperand = p
	} else if arithmeticTokens[2].tokenType == "NUMBER" {
		firstOperand, _ = strconv.Atoi(arithmeticTokens[2].tokenValue)
	}

	if arithmeticTokens[4].tokenType == "MEMORY" {
		secondOperand = m
	} else if arithmeticTokens[4].tokenType == "POINTER" {
		secondOperand = p
	} else if arithmeticTokens[4].tokenType == "NUMBER" {
		secondOperand, _ = strconv.Atoi(arithmeticTokens[4].tokenValue)
	}

	if arithmeticTokens[3].tokenValue == "+" {
		m = firstOperand + secondOperand
	} else if arithmeticTokens[3].tokenValue == "-" {
		if firstOperand-secondOperand < 0 {
			runtimeError(line, lineNumber, "Negative number. Memory Cannot Hold Negative Numbers")
		} else {
			m = firstOperand - secondOperand
		}
	} else if arithmeticTokens[3].tokenValue == "*" {
		m = firstOperand * secondOperand
	} else if arithmeticTokens[3].tokenValue == "/" {
		if secondOperand == 0 {
			runtimeError(line, lineNumber, "Cannot Divide Entity by Zero")
		} else {
			testValue := float32(firstOperand) / float32(secondOperand)
			if isIntegral(testValue) {
				m = int(testValue)
			} else {
				runtimeError(line, lineNumber, "Memory Cannot Have Decimal Values")
			}
		}
	}

	return p, m
}

/*
A Function which takes care of the assignment of the memory value at run time

Parameters:
	- assignmentTokens: []tokens containing the tokens to be used for the assignment
	- p: int - the current pointer value
	- m: int - the current memory value

Return Values:
	- p: int - the new pointer value
	- m: int - the new memory value
*/
func assignMemory(assignmentTokens []token, p int) (int, int) {

	var memory int
	if assignmentTokens[2].tokenType == "POINTER" {
		memory = p
	} else {
		number := assignmentTokens[2].tokenValue
		memory, _ = strconv.Atoi(number)
	}

	return p, memory
}

/*
A function that will do a conditional check on the program during runtime
	If it decides to skip the execution, the program will skip the line
	It if decides to execute it, it will recursively try to find and execute the statements

Parameters:
	- conditionalTokens: []tokens containing the tokens to be used for the conditional check
	- tokens: []tokens containing the tokens of the line that is being executed
	- line: string - the line that is being executed
	- lineNumber: int - the line number of the line that is being executed
	- p: int - the current pointer value
	- m: int - the current memory value

Return Values:
	- p: int - the new pointer value
	- m: int - the new memory value

Note: This function will recursively call the parse function which calls this function to parse the statements
*/
func doConditionalCheck(tokens []token, conditionTokens []token, line string, lineNumber int, p int, m int) (int, int) {

	var shouldWeDo bool = false
	var firstOperand int = 0
	var secondOperand int = 0

	if conditionTokens[1].tokenType == "NUMBER" {
		firstOperand, _ = strconv.Atoi(conditionTokens[1].tokenValue)
	} else if conditionTokens[1].tokenType == "MEMORY" {
		firstOperand = m
	} else if conditionTokens[1].tokenType == "POINTER" {
		firstOperand = p
	}

	if conditionTokens[3].tokenType == "NUMBER" {
		secondOperand, _ = strconv.Atoi(conditionTokens[3].tokenValue)
	} else if conditionTokens[3].tokenType == "MEMORY" {
		secondOperand = m
	} else if conditionTokens[3].tokenType == "POINTER" {
		secondOperand = p
	}

	if conditionTokens[2].tokenValue == ">" {
		if firstOperand > secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "<" {
		if firstOperand < secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "==" {
		if firstOperand == secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "!=" {
		if firstOperand != secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == ">=" {
		if firstOperand >= secondOperand {
			shouldWeDo = true
		}
	} else if conditionTokens[2].tokenValue == "<=" {
		if firstOperand <= secondOperand {
			shouldWeDo = true
		}
	}

	if shouldWeDo {
		unTouchedTokens := tokens[4:]
		p, m = parser(unTouchedTokens, line, lineNumber, p, m)
	}

	return p, m
}


/*
A function that will do a loop check on the program during runtime
	If it decides to skip the execution, the program will skip the line
	It if decides to execute it, it will recursively try to find and execute the statements for the specified number of times

Eg: It can parse following similar lines:
	- for 5 if pointer == 1 print char pointer
	- for pointer print memory

Parameters: 
	- tokens: []tokens containing the tokens to be used for the line execution
	- loopTokens: []tokens containing the tokens to be used for the loop
	- line: string - the line that is being executed
	- lineNumber: int - the line number of the line that is being executed
	- p: int - the current pointer value
	- m: int - the current memory value

Return Values:
	- p: int - the new pointer value
	- m: int - the new memory value
*/
func doLoops (tokens []token, loopTokens []token, line string, lineNumber int, p int, m int ) (int, int) {
	var timesToLoop int

	if loopTokens[1].tokenType == "NUMBER" {
		timesToLoop, _ = strconv.Atoi(loopTokens[1].tokenValue)
	} else if loopTokens[1].tokenType == "MEMORY" {
		timesToLoop = m
	} else if loopTokens[1].tokenType == "POINTER" {
		timesToLoop = p
	}

	tokensToLoop := tokens[len(loopTokens):]
	
	for i := 0; i < timesToLoop; i++ {
		p, m = parser(tokensToLoop, line, lineNumber, p, m)
	}

	return p, m
}

/*
A function which returns the ASCII Character of the given number.

Parameters:
	- number(int): The number which is to be converted to ASCII Character.

Returns:
	- string: The ASCII Character of the given number.
*/
func returnASCII(num int) string {
	return string(rune(num))
}

/*
Function which returns a boolean if the passed floating number is actually an integer

Parameters:
	- number(float64): The number to be checked

Return Value:
	- bool: True if the number is an integer, false otherwise
*/
func isIntegral(val float32) bool {
	return val == float32(int(val))
}
