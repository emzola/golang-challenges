package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	return input
}

func Quotes(author, quote string) string {
	return author + " " + "says, " + `"` + quote + `"` 
}

func main() {
	fmt.Print("What is the quote? ")
	quote := prompt()
	fmt.Print("Who said it? ")
	author := prompt()
	authorQuote := Quotes(author, quote)
	fmt.Println(authorQuote)
	
	// Extra Challenge
	var quotes = map[string]string{
		"Jenna Jameson": "My definition of courage is never letting anyone define you.",
		"Oscar Wilde": "Be yourself; everyone else is already taken.",
		"Marilyn Monroe": "I'm selfish, impatient and a little insecure. I make mistakes, I am out of control and at times hard to handle. But if you can't handle me at my worst, then you sure as hell don't deserve me at my best.",
		"Albert Einstein": "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
		"Frank Zappa": "So many books, so little time.",
		"Marcus Tullius Cicero": "A room without books is like a body without a soul.",
		"Zig Ziglar": "Courage is on display every day, and only the courageous wring the most out of life.",
		"Bernard M. Baruch": "Be who you are and say what you feel, because those who mind don't matter, and those who matter don't mind.",
	}

	displayQuotesMap(quotes)
}

// Extra Challenge
func displayQuotesMap(quotes map[string]string) {
	// Make map print in alphebetical order
	var authors []string
	for author := range quotes {
		authors = append(authors, author)
	}
	sort.Strings(authors)
	for _, author := range authors {
		fmt.Println(author + " " + "says, " + `"` + quotes[author] + `"`)
	}
}