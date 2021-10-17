package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func prompt() int16 {
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

	return int16(number)
}

func CalcRetirement(currentAge, retirementAge int16) int16 {
	return retirementAge - currentAge
}

func Output(yearsLeft int16) string {
	var currentYear int16 = int16(time.Now().Year())
	retirementYear := currentYear + yearsLeft
	if yearsLeft > 0 {
		return fmt.Sprintf("You have %d years left until you can retire. It's %d, so you can retire in %d.", yearsLeft, currentYear, retirementYear)
	} else {
		return fmt.Sprintf("You have no more years left to work. It's %d, so you can already retire.", currentYear)
	}
}

func main () {
	fmt.Print("What is your current Age? ")
	currentAge := prompt()
	
	fmt.Print("At what age would you like to retire? ")
	retirementAge := prompt()

	yearsLeft := CalcRetirement(currentAge, retirementAge)
	fmt.Println(Output(yearsLeft))
}
