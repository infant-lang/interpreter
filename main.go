package main

import "fmt"

var message string =  ""
var panicMessage string = ""

func main () {
	
	defer func() {     
		if e := recover(); e != nil {
			fmt.Print(message)
			fmt.Print(panicMessage)
		} else {
			fmt.Print(message)
		}
	}()
	
	fileTextLines := readFile(1)

	p := 0
	m := 0

	for lineNumber, eachline := range fileTextLines {
		p, m = parse(eachline, lineNumber + 1, p, m)
	}


}