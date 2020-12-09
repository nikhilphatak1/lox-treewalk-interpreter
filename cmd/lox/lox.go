package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"text/scanner"
	"strings"
	"os"
)

func main() {
	var err error
    if len(os.Args) > 2 {
		log.Println("Usage: `lox` for prompt or `lox $SCRIPT` to run script")
		os.Exit(1)
    } else if len(os.Args) == 2 {
		err = runScript(os.Args[1])
	} else {
		err = runPrompt()
	}

	if err != nil {
		log.Println("Lox failure:", err)
		os.Exit(1)
	}
}

// RunScript run Lox script
func runScript(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("Unable to read file.", err)
	}

	err = run(string(bytes))
	if err != nil {
		return err
	}

	return nil
}

// RunPrompt Lox REPL
func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln("Unable to read commit message from Stdin.", err)
		}

		message = strings.TrimRight(message, "\r\n")
		if message == "" {
			break
		}

		err = run(message)
		err = nil
	}

	return nil
}

func run(src string) error {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "lox"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}
	return nil
}

func sendError(line int, message string) {
	reportError(line, "", message)
}

func reportError(line int, where string, message string) {
	log.Println("[line", line, "] Error", where, ":", message)
}
