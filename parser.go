package main

func parse(line string, lineNumber int, p int, m int) (int, int) {
	tokens := lex(line, lineNumber)
	p, m = parser(tokens, line, lineNumber, p, m)

	return p, m
}

/*
Function that gets the tokens of the line that is currently being parsed
	For Conditionals and Looping, the function will call itself recursively until it
reaches the end of the line finding the statement in the current condition

Parameters:
	- tokens: []token - the tokens of the line currently being parsed
	- line: string - the line currently being parsed
	- lineNumber: int - the line number of the line currently being parsed
	- p: int - the current position of the pointer
	- m: int - the current memory value

Return Values:
	- p: int - the new position of the pointer
	- m: int - the new memory value
*/
func parser(tokens []token, line string, lineNumber int, p int, m int) (int, int) {

	if len(tokens) == 0 {
		return p, m
	}

	if tokens[0].tokenType == "ACTION" {
		actionTokens := checkAction(tokens, line, lineNumber)
		if actionTokens != nil {
			p, m = pointerMovements(actionTokens, line, lineNumber, p, m) // Movement of the pointer
		}
	} else if tokens[0].tokenType == "PRINT" {
		printTokens := checkPrint(tokens, line, lineNumber)
		if printTokens != nil {
			p, m = printStuff(printTokens, p, m)

		}
	} else if tokens[0].tokenType == "MEMORY" {

		if len(tokens) == 3 {
			assignmentTokens := checkAssignment(tokens, line, lineNumber)
			if assignmentTokens != nil {
				p, m = assignMemory(assignmentTokens, p)
			}
		} else if len(tokens) == 5 {
			arithmeticTokens := checkArithmetic(tokens, line, lineNumber)
			if arithmeticTokens != nil {
				p, m = doArithmetic(arithmeticTokens, line, lineNumber, p, m)
			}
		} else {
			printExpectedTokenError(line, lineNumber, "pointer or memory or number")
		}

	} else if tokens[0].tokenType == "CONDITION" {
		conditionTokens := checkCondition(tokens, line, lineNumber)
		if conditionTokens != nil {
			p, m = doConditionalCheck(tokens, conditionTokens, line, lineNumber, p, m)
		}

	} else if tokens[0].tokenType == "LOOP" {
		loopTokens := checkLoop(tokens, line, lineNumber)
		if loopTokens != nil {
			p, m = doLoops(tokens, loopTokens, line, lineNumber, p, m)
		}

	} else {
		printParseError(line, lineNumber, "")
	}

	return p, m
}