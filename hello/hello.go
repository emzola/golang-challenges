package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PromptInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func Output() string {
	name := PromptInput()
	return fmt.Sprintf("Hello, %s, nice to meet you!", name)
}

func main() {
	fmt.Print("What is your name? ")
	output := Output()
	fmt.Println(output)
}