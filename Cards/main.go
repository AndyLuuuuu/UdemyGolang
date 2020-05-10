package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")
	cards.shuffle()
	fmt.Println(cards)
}
