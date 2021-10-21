package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNumber() int8 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	number, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int8(number)
}

func DividePizza(numPizza, numSlices, numPeople int8) int8 {
	totalSlices := numSlices * numPizza
	return totalSlices / numPeople
}

func Remainder(numPizza, numSlices, numPeople int8) int8 {
	totalSlices := numSlices * numPizza
	return totalSlices % numPeople
}

func output(people, pizza, result, remainder int8) {
	resultText := "piece"
	remainderText := "piece"
	if result > 1 {
		resultText = "pieces"
	}
	if remainder > 1 {
		remainderText = "pieces"
	}
	fmt.Printf("%d people with %d pizzas.\nEach person gets %d %s of pizza.\nThere are %d leftover %s.", people, pizza, result, resultText, remainder, remainderText)
}

func main() {
	fmt.Print("How many people? ")
	people := getNumber()
	fmt.Print("How many pizzas do you have? ")
	pizza := getNumber()
	fmt.Print("How many slices per pizza? ")
	slices := getNumber()

	result := DividePizza(pizza, slices, people)
	remainder := Remainder(pizza, slices, people)
	output(people, pizza, result, remainder)
}
