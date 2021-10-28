package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func prompt() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} 

	input = strings.TrimSpace(input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

func ConvertCurrency(amount, exchangeRate float64) float64 {
	rateTo := 100.00
	amountTo := (amount * exchangeRate) / rateTo
	return toFixed(amountTo, 2)
}

func main() {
	fmt.Print("How many euros are you exchanging? ")
	amount := prompt()
	fmt.Print("What is the exchange rate? ")
	exchangeRate := prompt()

	result := ConvertCurrency(amount, exchangeRate)
	fmt.Printf("%.2f euros at an exchange rate of %.2f is %.2f U.S. dollars.", amount, exchangeRate, result)
}