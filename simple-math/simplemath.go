package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func prompt() uint64 {
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

func Add(firstNumber, secondNumber uint64) uint64 {
	return firstNumber + secondNumber
}

func Subtract(firstNumber, secondNumber uint64) uint64 {
	return firstNumber - secondNumber
}

func Multiply(firstNumber, secondNumber uint64) uint64 {
	return firstNumber * secondNumber
}

func Divide(firstNumber, secondNumber uint64) uint64 {
	return firstNumber / secondNumber
}

type generator struct {
	name string
	fn func(arg1, arg2 uint64) uint64
}

func display(firstNumber uint64, secondNumber uint64, result uint64, funcName string) {
	var sign string
	switch funcName {
	case "add":
		sign = "+"
	case "subtract":
		sign = "-"
	case "multiply":
		sign = "*"
	case "divide":
		sign = "/"
	}
	fmt.Printf("%d %s %d = %d\n", firstNumber, sign, secondNumber, result)
}

func main() {
	var firstNumber uint64
	var secondNumber uint64

	for i := 0; i < 2; i++ {
		switch i {
		case 1:
			fmt.Print("What is the second number? ")
			secondNumber = prompt()
		default:
			fmt.Print("What is the first number? ")
			firstNumber = prompt()
		}
	}

	gen := generator{"add", Add}
	addition := Add(firstNumber, secondNumber)
	display(firstNumber, secondNumber, addition, gen.name)

	gen = generator{"subtract", Subtract}
	subtraction := Subtract(firstNumber, secondNumber)
	display(firstNumber, secondNumber, subtraction, gen.name)

	gen = generator{"multiply", Multiply}
	multiplication := Multiply(firstNumber, secondNumber)
	display(firstNumber, secondNumber, multiplication, gen.name)

	gen = generator{"divide", Divide}
	division := Divide(firstNumber, secondNumber)
	display(firstNumber, secondNumber, division, gen.name)
}