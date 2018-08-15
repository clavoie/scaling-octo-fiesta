package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func printHelpAndExit() {
	log.Fatal(`Usage: scaling-octo-fiesta input-file output-file`)
}

func main() {
	args := os.Args[1:]
	argn := len(args)

	if argn < 2 {
		printHelpAndExit()
	}

	fname := args[0]
	inputFile, err := os.Open(fname)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := inputFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(inputFile)
	radixTree := newRadixTree()
	for scanner.Scan() {
		radixTree.Add(scanner.Text())
	}

	radixTree.Print()
	fmt.Println("vim-go", args[0])
}
