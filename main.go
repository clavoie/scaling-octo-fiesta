package main

import (
	"bufio"
	"log"
	"os"
	"strings"
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

	fname := strings.TrimSpace(args[0])
	oname := strings.TrimSpace(args[1])

	if fname == "" || oname == "" {
		printHelpAndExit()
	}

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

	// radixTree.Print()
	// err = radixTree.Write(os.Stdout)
	outFile, err := os.Create(oname)

	if err != nil {
		log.Fatal(err)
	}

	err = radixTree.Write(outFile)

	if err != nil {
		log.Fatal(err)
	}
}
