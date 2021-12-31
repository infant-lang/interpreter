package main

func main () {
	
	fileTextLines := readFile(1)

	p := 0
	m := 0

	for lineNumber, eachline := range fileTextLines {
		p, m = parse(eachline, lineNumber + 1, p, m)
	}

}