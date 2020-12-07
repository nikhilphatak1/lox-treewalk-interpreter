package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"os"
)

func main() {
	var err error
    if len(os.Args) > 2 {
		log.Println("Usage: `lox` for prompt or `lox $SCRIPT` to run script")
		os.Exit(1)
    } else if len(os.Args) == 2 {
		err = RunScript(os.Args[1])
	} else {
		err = RunPrompt()
	}

	if err != nil {
		log.Println("Lox failure:", err)
		os.Exit(1)
	}
}

// RunScript run Lox script
func RunScript(path string) error {
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
func RunPrompt() error {
	for {
		reader := bufio.NewReader(os.Stdin)
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
		if err != nil {
			return err
		}
	}

	return nil
}

func run(source string) error {

	return nil
}
