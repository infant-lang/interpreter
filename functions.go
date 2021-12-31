package main

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

	if tokens[0].tokenType == "PRINT" {
		if tokens[1].tokenType == "MEMORY" {
			return tokens[:2]
		} else if tokens[1].tokenType == "POINTER" {
			return tokens[:2]
		}

		if tokens[1].tokenType == "CHAR" {
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
	- Grammar 2: CONDITION POINTER COMPARISON MEMORY
	- Grammar 3: CONDITION MEMORY COMPARISON NUMBER
	- Grammar 4: CONDITION POINTER COMPARISON NUMBER
*/
func checkCondition(tokens []token, line string, lineNumber int) []token {

	if tokens[0].tokenType == "CONDITION" {
		if tokens[2].tokenType == "COMPARISON" {
			if tokens[1].tokenType == "MEMORY" {
				if tokens[3].tokenType == "POINTER" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[4].tokenType)

			} else if tokens[1].tokenType == "POINTER" {

				if tokens[3].tokenType == "MEMORY" {
					return tokens[:4]
				} else if tokens[3].tokenType == "NUMBER" {
					return tokens[:4]
				}
				printParseError(line, lineNumber, tokens[4].tokenType)
			}
			printParseError(line, lineNumber, tokens[3].tokenType)
		}
		printParseError(line, lineNumber, tokens[2].tokenValue)
	}

	printParseError(line, lineNumber, tokens[1].tokenValue)
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
