package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
The User Defined Type token.
	- tokType: The type of the token.
	- tokVal: The value of the token.


POSSIBLE VALUES:
	- + ARITHMETIC		|	- ARITHMETIC
	- * ARITHMETIC		| 	/ ARITHMETIC
	- = ASSIGNMENT		|	== COMPARISON
	- < COMPARISON		|	<= COMPARISON
	- > COMPARISON		|	>= COMPARISON
	- || COMPARISON		|	&& COMPARISON
	- != COMPARISON

	- char CHAR			| 	for LOOP
	- if CONDITION		|	left DIRECTION
	- move ACTION		|	memory MEMORY
	- new SPECIAL		| 	print PRINT
	- pointer POINTER	| 	right DIRECTION
	- tab SPECIAL		|	 0-9 NUMBER
*/
type token struct {
	tokType string // The type of the token or CATEGORY
	tokVal  string // The value of the token or VALUE
}

/* 
A function which takes in a string
	- and returns a slice of strings.
	- The lex can identify the following tokens:
	- char, for, if, left, memory, move, new, pointer, print, right, tab
	- and the following operators: +, -, *, /, =, <, >, <=, >=, ==, !=, &&, ||
	- and integers

Parameters:
	- line: The line of code to be tokenized.
	- lineNumber: The line number of the line of code.
*/
func lex(line string, lineNumber int) []token {

	var tokens []token
	numeral := ""
	var i int

	for i = 0; i < len(line); i++ {
		if line[i] == ' ' {
			if numeral != "" {
				tokens = append(tokens, token{"NUMBER", numeral})
				numeral = ""
			}
			continue
		} else if line[i] == '+' {
			tokens = append(tokens, token{"ARITHMETIC", "+"}) // +
		} else if line[i] == '-' {
			tokens = append(tokens, token{"ARITHMETIC", "-"}) // -
		} else if line[i] == '*' {
			tokens = append(tokens, token{"ARITHMETIC", "*"}) // *
		} else if line[i] == '/' {
			tokens = append(tokens, token{"ARITHMETIC", "/"}) // /
		} else if line[i] == '=' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "=="}) // ==
				i++
			} else {
				tokens = append(tokens, token{"ASSIGNMENT", "="}) // =
			}
		} else if line[i] == '<' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "<="}) // <=
				i++
			} else {
				tokens = append(tokens, token{"COMPARISON", "<"}) // <
			}
		} else if line[i] == '>' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", ">="}) // >=
				i++
			} else {
				tokens = append(tokens, token{"COMPARISON", ">"}) // >
			}
		} else if line[i] == '!' {
			if line[i+1] == '=' {
				tokens = append(tokens, token{"COMPARISON", "!="}) // !=
				i++
			}
		} else if line[i] == '|' {
			if line[i+1] == '|' {
				tokens = append(tokens, token{"COMPARISON", "||"}) // ||
				i++
			}
		} else if line[i] == '&' {
			if line[i+1] == '&' {
				tokens = append(tokens, token{"COMPARISON", "&&"}) // &&
				i++
			}
		} else if line[i] == 'c' {
			if line[i+1] == 'h' {
				if line[i+2] == 'a' {
					if line[i+3] == 'r' {
						tokens = append(tokens, token{"CHAR", "char"}) // char
						i = i + 3
					}
				}
			}
		} else if line[i] == 'f' {
			if line[i+1] == 'o' {
				if line[i+2] == 'r' {
					tokens = append(tokens, token{"LOOP", "for"}) // for
					i = i + 2
				}
			}
		} else if line[i] == 'i' {
			if line[i+1] == 'f' {
				tokens = append(tokens, token{"CONDITION", "if"}) // if
				i = i + 1
			}
		} else if line[i] == 'l' {
			if line[i+1] == 'e' {
				if line[i+2] == 'f' {
					if line[i+3] == 't' {
						tokens = append(tokens, token{"DIRECTION", "left"}) // left
						i = i + 3
					}
				}
			}
		} else if line[i] == 'm' {
			if line[i+1] == 'o' {
				if line[i+2] == 'v' {
					if line[i+3] == 'e' {
						tokens = append(tokens, token{"ACTION", "move"}) // move
						i = i + 3
					}
				}
			} else if line[i+1] == 'e' {
				if line[i+2] == 'm' {
					if line[i+3] == 'o' {
						if line[i+4] == 'r' {
							if line[i+5] == 'y' {
								tokens = append(tokens, token{"MEMORY", "memory"}) // memory
								i = i + 5
							}
						}
					}
				}
			}
		} else if line[i] == 'n' {
			if line[i+1] == 'e' {
				if line[i+2] == 'w' {
					tokens = append(tokens, token{"SPECIAL", "new"}) // new
					i = i + 2
				}
			}
		} else if line[i] == 'p' {
			if line[i+1] == 'r' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"PRINT", "print"}) // print
							i = i + 4
						}
					}
				}
			} else if line[i+1] == 'o' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							if line[i+5] == 'e' {
								if line[i+6] == 'r' {
									tokens = append(tokens, token{"POINTER", "pointer"}) // pointer
									i = i + 6
								}
							}
						}
					}
				}
			}
		} else if line[i] == 'r' {
			if line[i+1] == 'i' {
				if line[i+2] == 'g' {
					if line[i+3] == 'h' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"DIRECTION", "right"}) // right
							i = i + 4
						}
					}
				}
			}
		} else if line[i] == 't' {
			if line[i+1] == 'a' {
				if line[i+2] == 'b' {
					tokens = append(tokens, token{"SPECIAL", "tab"}) // tab
				}
			}
		} else if line[i] == '0' || line[i] == '1' || line[i] == '2' || line[i] == '3' ||
			line[i] == '4' || line[i] == '5' || line[i] == '6' || line[i] == '7' ||
			line[i] == '8' || line[i] == '9' {
			numeral = numeral + string(line[i])
			if i == len(line)-1 {
				tokens = append(tokens, token{"NUMBER", numeral}) // numberal
				numeral = ""
			}
		} else {
			printError(line, lineNumber, i)
		}
	}

	return tokens
}

/*
The function which prints the error to the console.
	- line: The line which contains the error.
	- lineNumber: The line number of the error.
	- i: The index of the error at the given line

	Returns: Nothing

	Prints: The error to the console.
*/
func printError(line string, lineNumber int, i int) {
	fmt.Println()
	fmt.Println("(*_*) Tokenization Error:")
	fmt.Println(`Unknown token: "` + string(line[i]) + `"` + " at line " + strconv.Itoa(lineNumber))
	fmt.Println(line)
	for j := 0; j < i; j++ {
		fmt.Print(" ")
	}
	fmt.Print("^")
	fmt.Println()
	os.Exit(1)
}
