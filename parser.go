package main

import "fmt"

func parse(line string, lineNumber int) {
	tokens := lex(line, lineNumber)
	for _, token := range tokens {
		fmt.Println(token)
	}
}