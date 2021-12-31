package main

import (
	"fmt"
)

func parse(line string, lineNumber int, p int, m int) (int, int) {
	tokens := lex(line, lineNumber)

	if tokens[0].tokenType == "ACTION" {
		actionTokens := checkAction(tokens, line, lineNumber)
		if actionTokens != nil {
			p, m = pointerMovements(actionTokens, line, lineNumber, p, m)  // Movement of the pointer
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
			fmt.Println("Loop:")
		}

	} else {
		printParseError(line, lineNumber, "")
	}

	return p, m
}