package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main () {
	FILE_NAME := os.Args[1]

	readFile, err := os.Open(FILE_NAME)
 
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
 
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
 
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
 
	readFile.Close()
 
	for _, eachline := range fileTextLines {
		fmt.Println(eachline)
	}


}