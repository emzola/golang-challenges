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

func getAmount() float64 {
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
func getCountry() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} 

	input = strings.TrimSpace(input)

	return input
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num*output)) / output
}

func ConvertCurrency(amount float64, exchangeRate float64) float64 {
	rateTo := 100.00
	amountTo := (amount * exchangeRate) / rateTo
	return toFixed(amountTo, 2)
}

func main() {
	exchangeRates := map[string]float64{
		"US": 137.51,
		"GB": 110.24,
		"DE": 100,
	}

	currencyNames := map[string]string{
		"US": "US Dollars",
		"GB": "British Pounds",
		"DE": "Euros",
	}

	fmt.Printf("How many euros are you exchanging? ")
	amount := getAmount()
	fmt.Print("What is the country? ")
	country := getCountry()

	_, ok := exchangeRates[country]
	if ok {
		result := ConvertCurrency(amount, exchangeRates[country])
		fmt.Printf("%.2f euros at an exchange rate of %.2f is %.2f %s.", amount, exchangeRates[country], result, currencyNames[country])
	} else {
		fmt.Printf("No country with code %s\n", country)
	}	
}