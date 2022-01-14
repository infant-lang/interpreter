package main

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
	tokenType string // The type of the token or CATEGORY
	tokenValue  string // The value of the token or VALUE
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
		if line[i] == '#' {
			break
		} else  if line[i] == ' ' {
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
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'f' {
			if line[i+1] == 'o' {
				if line[i+2] == 'r' {
					tokens = append(tokens, token{"LOOP", "for"}) // for
					i = i + 2
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1)}
		} else if line[i] == 'i' {
			if line[i+1] == 'f' {
				tokens = append(tokens, token{"CONDITION", "if"}) // if
				i = i + 1
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'l' {
			if line[i+1] == 'e' {
				if line[i+2] == 'f' {
					if line[i+3] == 't' {
						tokens = append(tokens, token{"DIRECTION", "left"}) // left
						i = i + 3
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'm' {
			if line[i+1] == 'o' {
				if line[i+2] == 'v' {
					if line[i+3] == 'e' {
						tokens = append(tokens, token{"ACTION", "move"}) // move
						i = i + 3
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else if line[i+1] == 'e' {
				if line[i+2] == 'm' {
					if line[i+3] == 'o' {
						if line[i+4] == 'r' {
							if line[i+5] == 'y' {
								tokens = append(tokens, token{"MEMORY", "memory"}) // memory
								i = i + 5
							} else { lexError(line, lineNumber, i+5) }
						} else { lexError(line, lineNumber, i+4) }
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'n' {
			if line[i+1] == 'e' {
				if line[i+2] == 'w' {
					tokens = append(tokens, token{"SPECIAL", "new"}) // new
					i = i + 2
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'p' {
			if line[i+1] == 'r' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"PRINT", "print"}) // print
							i = i + 4
						} else { lexError(line, lineNumber, i+4) }
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else if line[i+1] == 'o' {
				if line[i+2] == 'i' {
					if line[i+3] == 'n' {
						if line[i+4] == 't' {
							if line[i+5] == 'e' {
								if line[i+6] == 'r' {
									tokens = append(tokens, token{"POINTER", "pointer"}) // pointer
									i = i + 6
								} else { lexError(line, lineNumber, i+6) }
							} else { lexError(line, lineNumber, i+5) }
						} else { lexError(line, lineNumber, i+4) }
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 'r' {
			if line[i+1] == 'i' {
				if line[i+2] == 'g' {
					if line[i+3] == 'h' {
						if line[i+4] == 't' {
							tokens = append(tokens, token{"DIRECTION", "right"}) // right
							i = i + 4
						} else { lexError(line, lineNumber, i+4) }
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == 's' { 
			if line[i+1] == 'p' {
				if line[i+2] == 'a' {
					if line[i+3] == 'c' {
						if line[i+4] == 'e' {
							tokens = append(tokens, token{"SPECIAL", "space"}) // space
							i = i + 4
						} else { lexError(line, lineNumber, i+4) }
					} else { lexError(line, lineNumber, i+3) }
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
			
		} else if line[i] == 't' {
			if line[i+1] == 'a' {
				if line[i+2] == 'b' {
					tokens = append(tokens, token{"SPECIAL", "tab"}) // tab
					i = i + 2
				} else { lexError(line, lineNumber, i+2) }
			} else { lexError(line, lineNumber, i+1) }
		} else if line[i] == '0' || line[i] == '1' || line[i] == '2' || line[i] == '3' ||
			line[i] == '4' || line[i] == '5' || line[i] == '6' || line[i] == '7' ||
			line[i] == '8' || line[i] == '9' {
			numeral = numeral + string(line[i])
			if i == len(line)-1 {
				tokens = append(tokens, token{"NUMBER", numeral}) // numberal
				numeral = ""
			}
		} else {
			lexError(line, lineNumber, i)
		}
	}

	return tokens
}