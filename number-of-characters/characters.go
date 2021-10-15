package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func Length(input string) int {
	return len(input)
}

func output(input string, length int) {
	if len(input) < 1 {
		fmt.Println("Input must not be empty")
	} else {
		fmt.Printf("%s has %d characters.", input, length)
	}
}

func main() {
	fmt.Print("What is the input string? ")
	input := prompt()
	length := Length(input)
	output(input, length)
}