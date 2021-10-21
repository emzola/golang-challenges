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

func getUnit() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func getNumber() int64 {
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

func SqUnit(length, width int64) int64 {
	return length * width
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num * output)) / output
}

func SqFeetToSqMeter(unit int64) float64 {
	const conversionFactor float64 = 0.09290304
	result := float64(unit) * conversionFactor
	return toFixed(result, 3)
}

func SqMeterToSqFeet(unit int64) float64 {
	const conversionFactor float64 = 10.76391041671
	result := float64(unit) * conversionFactor
	return toFixed(result, 3)
}

func output(sqUnit int64, convSqUnit float64, unit string) {
	var unitText = "feet"
	if unitText == unit {
		unitText = "meters"
	} 
	fmt.Printf("You entered dimensions of 15 %s by 20 %s.\n", unit, unit)
	fmt.Printf("The area is \n%d square %s\n%.3f square %s", sqUnit, unit, convSqUnit, unitText)
}

func main() {
	fmt.Print("Choose unit: f or m? ")
	u := getUnit()
	var unit string
	switch strings.ToLower(u) {
	case "f":
		unit = "feet"
	case "m":
		unit = "meters"
	default:
		unit = "feet"
	}
	fmt.Printf("What is the length of the room in %s? ", unit)
	length := getNumber()
	fmt.Printf("What is the width of the room in %s? ", unit)
	width := getNumber()

	sqUnit := SqUnit(length, width)

	var sqMeter float64
	var sqFeet float64
	var convSqUnit float64

	if unit == "feet" {
		sqMeter = SqFeetToSqMeter(sqUnit)
		convSqUnit = sqMeter
	}

	if unit == "meters" {
		sqFeet = SqMeterToSqFeet(sqUnit)
		convSqUnit = sqFeet
	}
	
	output(sqUnit, convSqUnit, unit)
}