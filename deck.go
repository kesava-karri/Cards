package main

import "fmt"

//Creating a new data type called deck 

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string {"Spades", "Clovers", "Diamonds", "Hearts"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _,suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() { //receiver conventionally named with a single letter along with the data type
		for i:=0; i<len(d); i++ {  //Can also use range to iterate over a data structure
		fmt.Println(i, d[i])
	}
}

func deal(d deck,handSize int) (deck,deck) {
return d[:handSize], d[handSize:]
}