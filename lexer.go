package main

import (
	"fmt"
	"os"
	"strconv"
)

// a function named lex which takes in a string
// and returns a slice of strings.
// The lex can identify the following tokens:
// char, false, for, if, left, memory, move, new, pointer, print, right, tab, true
// and the following operators: +, -, *, /, =, <, >, <=, >=, ==, !=, &&, ||
// and integers
func lex(line string, lineNumber int) []string {

	var tokens []string
	numeral := ""
	var i int
	for i = 0; i < len(line); i++ {

		if line[i] == ' ' {
			if numeral != "" {
				tokens = append(tokens, numeral)
				numeral = ""
			}
			continue
		} else if line[i] == '+' {
			tokens = append(tokens, "+") // +
		} else if line[i] == '-' {
			tokens = append(tokens, "-") // -
		} else if line[i] == '*' {
			tokens = append(tokens, "*") // *
		} else if line[i] == '/' {
			tokens = append(tokens, "/") // /
		} else if line[i] == '=' {
			if line[i+1] == '=' {
				tokens = append(tokens, "==") // ==
				i++
			} else {
				tokens = append(tokens, "=") // =
			}
		} else if line[i] == '<' {
			if line[i+1] == '=' {
				tokens = append(tokens, "<=") // <=
				i++
			} else {
				tokens = append(tokens, "<") // <
			}
		} else if line[i] == '>' {
			if line[i+1] == '=' {
				tokens = append(tokens, ">=") // >=
				i++
			} else {
				tokens = append(tokens, ">") // >
			}
		} else if line[i] == '!' {
			if line[i+1] == '=' {
				tokens = append(tokens, "!=") // !=
				i++
			}
		} else if line[i] == '|' {
			if line[i+1] == '|' {
				tokens = append(tokens, "||") // ||
				i++
			}
		} else if line[i] == '&' {
			if line[i+1] == '&' {
				tokens = append(tokens, "&&") // &&
				i++
			}
		} else if line[i] == 'c' {
			if line[i+1] == 'h' {
				if line[i+2] == 'a' {
					if line[i+3] == 'r' {
						tokens = append(tokens, "char") // char
						i = i + 3
					}
				}
			}
		} else if line[i] == 'f' {
			if line[i+1] == 'o' {
				if line[i+2] == 'r' {
					tokens = append(tokens, "for") // for
					i = i + 2
				}
			} else if line[i+1] == 'a' {
				if line[i+2] == 'l' {
					if line[i+3] == 's' {
						if line[i+4] == 'e' {
							tokens = append(tokens, "false") // false
							i = i + 4
						}
					}
				}
			}
		} else if line[i] == 'i' {
			if line[i+1] == 'f' {
				tokens = append(tokens, "if") // if
				i = i + 1
			}
		} else if line[i] == 'l' {
			if line[i+1] == 'e' {
				if line[i+2] == 'f' {
					if line[i+3] == 't' {
						tokens = append(tokens, "left") // left
						i = i + 3
					}
				}
			}
		} else if line[i] == 'm' {
			if line[i+1] == 'o' {
				if line[i+2] == 'v' {
					if line[i+3] == 'e' {
						tokens = append(tokens, "move") // move
						i = i + 3
					}
				}
			} else if line[i+1] == 'e' {
				if line[i+2] == 'm' {
					if line[i+3] == 'o' {
						if line[i+4] == 'r' {
							if line[i+5] == 'y' {
								tokens = append(tokens, "memory") // memory
								i = i + 5
							}
						}
					}
				}
			}
		} else if line[i] == 'n' {
			if line[i+1] == 'e' {
				if line[i+2] == 'w' {
					tokens = append(tokens, "new") // new
					i = i + 2
				}
			}
		} else if line[i] == 'p' {
			if line[i+1] == 'r' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							tokens = append(tokens, "print") // print
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
									tokens = append(tokens, "pointer") // pointer
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
							tokens = append(tokens, "right") // right
							i = i + 4
						}
					}
				}
			}
		} else if line[i] == 't' {
			if line[i+1] == 'a' {
				if line[i+2] == 'b' {
					tokens = append(tokens, "tab") // tab
				}
			} else if line[i+1] == 'r' {
				if line[i+2] == 'u' {
					if line[i+3] == 'e' {
						tokens = append(tokens, "true") // true
						i = i + 3
					}
				}
			}
		} else if line[i] == '0' || line[i] == '1' || line[i] == '2' || line[i] == '3' || 
			line[i] == '4' || line[i] == '5' || line[i] == '6' || line[i] == '7' || 
			line[i] == '8' || line[i] == '9' {
			numeral = numeral + string(line[i])
			if i == len(line)-1 {
				tokens = append(tokens, numeral) // numerals
				numeral = ""
			}
		} else {
			printError(line, lineNumber, i)
		}
	}

	return tokens
}

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
