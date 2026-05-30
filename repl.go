package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")

		scanner.Scan()
		text := scanner.Text()

		fmt.Println("echoeing", text)
	}
}

func cleanInput(str string ) []string {
    lowercase := strings.ToLower(str)
    words := strings.Fields(lowercase)
    return words
}