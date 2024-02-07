package main

import "fmt"

func main() {
	greeting := greet("hh")
	fmt.Println(greeting)
}

// language represents the language’s code
type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": " שלום עולם ",       // Hebrew
	"ur": " ہیلو دنیا ",       // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in the specified language
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
