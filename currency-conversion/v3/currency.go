package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func getCurrency() string {
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
	amountTo := amount * exchangeRate
	return toFixed(amountTo, 2)
}

func main() {

	var exchange struct {
		Disclaimer string `json:"disclaimer"`
		License string `json:"license"`
		Timestamp int64 `json:"timestamp"`
		Base string `json:"base"`
		Rates map[string]float64 `json:"rates"`
	}

	APP_ID := os.Getenv("TOKEN")

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	req, err := http.NewRequest("GET", "https://openexchangerates.org/api/latest.json?app_id="+APP_ID, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("accept", "application/json")
	response, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	err = json.Unmarshal(body, &exchange)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}

	// Create currency code map from exchange.Rates map
	var names []string
	for name := range exchange.Rates {
		names = append(names, name)
	}

	var currencyCode = make(map[string]string)

	for _, value := range names {
		currencyCode[value] = value
	}

	// Promt for input
	fmt.Printf("How much USD do you want to exchange? ")
	amount := getAmount()
	fmt.Print("Which currency do you want to exchange to? ")
	currency := getCurrency()

	_, ok := exchange.Rates[currency]
	if ok {
		result := ConvertCurrency(amount, exchange.Rates[currency])
		fmt.Printf("%.2f USD at an exchange rate of %.2f is %.2f %s.", amount, exchange.Rates[currency], result, currencyCode[currency])
	} else {
		fmt.Printf("Could not find the currency with code %s\n", currency)
	}	
}