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

	return strings.TrimSpace(input)
}

func Output(noun, verb, adjective, adverb string) string {
	return fmt.Sprintf("Do you %s your %s %s %s? That's hilarious!", verb, adjective, noun, adverb)
}

func main() {
	var noun, verb, adjective, adverb string
	
	for i := 0; i < 4; i++ {
		switch i {
		case 1:
			fmt.Print("Enter a verb: ")
			verb = prompt()
		case 2:
			fmt.Print("Enter an adjective: ")
			adjective = prompt()
		case 3: 
			fmt.Print("Enter an adverb: ")
			adverb = prompt()
		default:
			fmt.Print("Enter a noun: ")
			noun = prompt()
		}
	}

	display := Output(noun, verb, adjective, adverb)
	fmt.Println(display)
}