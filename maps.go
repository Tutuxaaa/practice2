package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := `Mr and Mrs Dursley, of number four, Privet Drive, were proud to say that they were perfectly normal, thank you very much. 
	They were the last people you would expect to be involved in anything strange or mysterious, because they just did hold with such nonsense.`

	wordFrequency := countWords(text)
	printWordFrequency(wordFrequency)
}

func countWords(text string) map[string]int {
	frequency := make(map[string]int)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		frequency[lowerWord]++
	}

	return frequency
}

func printWordFrequency(freq map[string]int) {
	fmt.Println("Word frequency count:")
	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}
