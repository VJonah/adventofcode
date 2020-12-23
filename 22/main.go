package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	p1 := input[0]
	p2 := input[1]
	fmt.Println(p1.Deck)
	fmt.Println(p2.Deck)
	fmt.Println(playCombat(p1, p2).score())

}

func playCombat(p1, p2 Player) *Deck {
	for p1.Deck.size() > 0 && p2.Deck.size() > 0 {
		p1TopCard := p1.Deck.removeCard()
		p2TopCard := p2.Deck.removeCard()
		if p1TopCard.Value > p2TopCard.Value {
			p1.Deck.addCard(p1TopCard)
			p1.Deck.addCard(p2TopCard)
		} else {
			p2.Deck.addCard(p2TopCard)
			p2.Deck.addCard(p1TopCard)
		}
	}
	if p1.Deck.size() == 0 {
		return p2.Deck
	}
	return p1.Deck
}

type Card struct {
	Value int
}

type Deck struct {
	Cards []Card
}

func (d *Deck) peekTopCard() Card {
	if len(d.Cards) > 0 {
		return d.Cards[0]
	}
	return Card{Value: -1}
}

func (d *Deck) addCard(card Card) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck) removeCard() Card {
	topCard := Card{Value: -1}
	if len(d.Cards) > 0 {
		topCard = d.Cards[0]
		d.Cards = d.Cards[1:]
	}
	return topCard
}

func (d *Deck) score() int {
	sum := 0
	size := len(d.Cards)
	for i, card := range d.Cards {
		sum += card.Value * (size - i)
	}
	return sum
}

func (d *Deck) size() int {
	return len(d.Cards)
}

type Player struct {
	Deck *Deck
}

func parseInput(href string) []Player {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	var p1, p2 Player
	for i, line := range lines {
		if line == "" {
			p1 = generatePlayerDecks(lines[1:i])
			p2 = generatePlayerDecks(lines[i+2:])
		}
	}
	players := []Player{p1, p2}
	return players
}

func generatePlayerDecks(cards []string) Player {
	deck := &Deck{}
	player := Player{Deck: deck}
	for _, card := range cards {
		value, _ := strconv.Atoi(card)
		card := Card{Value: value}
		deck.addCard(card)
	}
	return player
}
