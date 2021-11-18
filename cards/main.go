package main

// files with same package can be treated as single file in the end
// aka, you can call funcs from another file under the same package

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	// range is a keyword not a func

	for i, card := range cards {
		fmt.Println(i, card)
	}

	for _, card := range cards {
		fmt.Println(card)
		// every declared variable must be used
	}

}

func newCard() string {
	return "Five of Diamonds"
}
