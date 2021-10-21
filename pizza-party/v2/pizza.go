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

func output(people, pieces, result int8) {
	resultText := "pizza"
	if result > 1 {
		resultText = "pizzas"
	}

	fmt.Printf("%d people with %d pieces each.\nYou need to buy %d %s.", people, pieces, result, resultText)
}

func main() {
	fmt.Print("How many people? ")
	people := getNumber()
	fmt.Print("How many pieces of pizza does each person want? ")
	pieces := getNumber()
	fmt.Print("How many slices per pizza? ")
	slices := getNumber()

	result := DividePizza(people, pieces, slices)
	output(people, pieces, result)
}
