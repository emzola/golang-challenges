package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/emzola/tipcalculator/calculator"
)

func promptInput() uint64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func display(tip, total float64) {
	fmt.Printf("Tip: $%0.2f \n", tip)
	fmt.Printf("Total: $%0.2f \n", total)
}

func main() {
	var bill, tipRate uint64

	fmt.Print("What is the bill amount? ")
	bill = promptInput()
	
	fmt.Print("What is the tip rate? ")
	tipRate = promptInput()

	tip := calculator.Calculate(bill, tipRate)

	total := tip + float64(bill)

	display(tip, total)
}