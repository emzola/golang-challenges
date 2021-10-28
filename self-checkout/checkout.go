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

type Cart struct {
	price int64
	quantity int64
}

func prompt() int64 {
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

	return number
}

func SubTotal(items []Cart) float64 {
	var subtotal float64
	for _, item := range items {
		price := item.price * item.quantity
		subtotal += float64(price)
	}
	return subtotal/100
}

func TaxRate(subtotal float64) float64 {
	tax := (subtotal * 100) * (5.5/100)
	return math.Round(tax) / 100
}

func Total(subTotal, taxRate float64) float64 {
	return subTotal + taxRate
}

func output(subTotal, taxRate,  total float64) {
	fmt.Printf("Subtotal: $%.2f\n", subTotal)
	fmt.Printf("Tax: $%.2f\n", taxRate)
	fmt.Printf("Total: $%.2f\n", total)
}

func main() {
	products := []Cart{}
	var price, quantity int64

	for i := 1; i < 4; i++ {
		fmt.Printf("Enter the price of item %d: ", i)
		cents := prompt()
		price = cents * 100
		fmt.Printf("Enter the quantity of item %d: ", i)
		quantity = prompt()

		product := Cart{
			price: price,
			quantity: quantity,
		}

		products = append(products, product)

	}

	subTotal := SubTotal(products)
	taxRate := TaxRate(subTotal)
	total := Total(subTotal, taxRate)
	output(subTotal, taxRate, total)
}