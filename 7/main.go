package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node interface {
	String() string
}

type Bag struct {
	Colour   string
	Quantity int
}

func (b *Bag) String() string { return fmt.Sprintf("Colour: %s, Quant: %d", b.Colour, b.Quantity) }

func main() {
	input := parseInput()
	//bagsContainingGold := findAllCombinations(input, "shiny gold", []Bag{})
	//fmt.Println(input["dark teal"][0].String())
	numberofBagsContained := findNumberOfBags(input, Bag{Colour: "shiny gold", Quantity: 0}, 0)
	fmt.Println(numberofBagsContained)
}

func findNumberOfBags(allBags map[string][]Bag, desiredBag Bag, total int) int {
	bagsThatAreContained := allBags[desiredBag.Colour]
	fmt.Printf("desiredBag: {%s}, total: %v, bagsThatAreContained: %d\n", desiredBag.String(), total, len(bagsThatAreContained))
	if len(bagsThatAreContained) == 1 && bagsThatAreContained[0].Colour == "other" {
		total += desiredBag.Quantity
		return total
	}
	total += desiredBag.Quantity
	if desiredBag.Quantity == 0 {
		desiredBag.Quantity = 1
	}
	for i := 0; i < desiredBag.Quantity; i++ {
		for _, bag := range bagsThatAreContained {
			total = findNumberOfBags(allBags, bag, total)
		}
	}
	return total
}

func findAllBagsContaining(bagColour string, allBags map[string][]Bag) []Bag {
	bagsThatContainThisColour := []Bag{}
	for bag, bagsItCanContain := range allBags {
		if contains(bagsItCanContain, bagColour) {
			bagsThatContainThisColour = append(bagsThatContainThisColour, Bag{Colour: bag})
		}
	}
	return bagsThatContainThisColour
}

func findAllCombinations(allBags map[string][]Bag, desiredBag string, visitedBags []Bag) []Bag {
	bagsThatContainThisColour := findAllBagsContaining(desiredBag, allBags)
	//fmt.Printf("desiredBag: %s, visitedBags: %v, bagsThatContainThisColour: %s\n", desiredBag, visitedBags, bagsThatContainThisColour)
	if len(bagsThatContainThisColour) == 0 {
		return visitedBags
	}
	for _, bag := range bagsThatContainThisColour {
		if !contains(visitedBags, bag.Colour) {
			visitedBags = append(visitedBags, bag)
		}
		visitedBags = findAllCombinations(allBags, bag.Colour, visitedBags)
	}
	return visitedBags
}

func contains(arr []Bag, val string) bool {
	for _, bag := range arr {
		if bag.Colour == val {
			return true
		}
	}
	return false
}

func parseInput() map[string][]Bag {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	bags := strings.Split(input, "\n")
	bagsMap := map[string][]Bag{}
	for _, bag := range bags {
		bag = bag[:len(bag)-1]
		bagAndCanContain := strings.Split(bag, " contain ")
		bagColour := extractBagColor(bagAndCanContain[0], true)
		possibleBags := strings.Split(bagAndCanContain[1], ", ")
		bagsMappedTo := []Bag{}
		for _, colour := range possibleBags {
			numberOfBags, _ := strconv.Atoi(string(colour[0]))
			newBag := Bag{Quantity: numberOfBags}
			//fmt.Printf("bagColour: %s,  numberofBags: %d, colour: %s\n", bagColour, numberOfBags, colour)
			containedBagColour := ""
			if numberOfBags > 1 {
				containedBagColour = extractBagColor(colour[2:], true)
			} else {
				containedBagColour = extractBagColor(colour[2:], false)
			}
			newBag.Colour = containedBagColour
			bagsMappedTo = append(bagsMappedTo, newBag)
		}
		bagsMap[bagColour] = bagsMappedTo
	}
	return bagsMap
}

func extractBagColor(s string, isMultiple bool) string {
	answer := ""
	if isMultiple {
		answer = s[:len(s)-5]
	} else {
		answer = s[:len(s)-4]
	}
	return answer
}
