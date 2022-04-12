package main

import "fmt"

func main() {
	cards := newDeck()
	cards = newDeckFromFile("my_cards")
	cards.shuffle()
	fmt.Println(cards.toString())
}
