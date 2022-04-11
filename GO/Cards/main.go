package main

import "fmt"

func main() {
	cards := newDeck()
	cards = newDeckFromFile("my_cards")
	fmt.Println(cards.toString())
}
