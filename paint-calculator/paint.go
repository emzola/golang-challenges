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

func promptShape() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} 

	input = strings.TrimSpace(input)

	return input
}

func promptDimension() float64 {
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

func SqFtRect(length, width float64) float64 {
	return length*width
}

func SqFtCircle(diameter float64) float64 {
	const pi float64 = 3.14159265
	total := diameter/2
	total *= total
	return toFixed(total*pi, 2)
}

func SqFtLShaped(lengthOne, widthOne, lengthTwo, widthTwo float64) float64 {
	areaOne := lengthOne*widthOne
	areaTwo := lengthTwo*widthTwo
	return areaOne + areaTwo
}

func PaintCalculator(size float64) int16 {
	const gallon float64 = 350
	return int16(math.Ceil(size/gallon))
}

func output(size float64, quantity int16) {
	gallonText := "gallon"
	if quantity > 1 {
		gallonText = "gallons"
	}

	fmt.Printf("You will need to purchase %d %s of paint to cover %.2f square feet.", quantity, gallonText, size)
}

func main() {
	var (
		roomSize, length, width, diameter, lengthTwo, widthTwo float64 
	)
	fmt.Print("Choose room shape: [R] Rectangle, [C] Round, [L] L-shaped? ")
	shape := promptShape()
	switch strings.ToLower(shape) {
	case "c":
		fmt.Print("What is the diameter of the room? ")
		diameter = promptDimension()
	case "l":
		fmt.Print("Length One? ")
		length = promptDimension()
		fmt.Print("Width One? ")
		width = promptDimension()
		fmt.Print("Length Two? ")
		lengthTwo = promptDimension()
		fmt.Print("Width One? ")
		widthTwo = promptDimension()
	default:
		fmt.Print("What is the length of the room? ")
		length = promptDimension()
		fmt.Print("What is the width of the room? ")
		width = promptDimension()		
	}

	if shape == "r" {
		roomSize = SqFtRect(length, width)
	} else if shape == "c" {
		roomSize = SqFtCircle(diameter)
	} else {
		roomSize = SqFtLShaped(length, width, lengthTwo, widthTwo)
	}	

	paintNeeded := PaintCalculator(roomSize)

	output(roomSize, paintNeeded)
}