package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(playGame(input))
}

func playGame(start []int) int {
	seenNumbers := map[int][]int{}
	previous := start[0]
	seenNumbers[previous] = append(seenNumbers[previous], 1)
	for i := 1; i < 30000000; i++ {
		//fmt.Printf("prev: %d, seenNumbers: %d\n", previous, seenNumbers)
		number := 0
		if i < len(start) {
			number = start[i]
		} else {
			if len(seenNumbers[previous]) == 1 {
				number = 0
			} else {
				numberOccurences := seenNumbers[previous]
				number = numberOccurences[len(numberOccurences)-1] - numberOccurences[len(numberOccurences)-2]
			}
		}
		seenNumbers[number] = append(seenNumbers[number], i+1)
		previous = number
	}
	return previous
}

func contains(n int, arr []int) bool {
	for _, val := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func parseInput(href string) []int {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	numbersAsStr := strings.Split(input, ",")
	startingNbs := []int{}
	for _, n := range numbersAsStr {
		number, _ := strconv.Atoi(n)
		startingNbs = append(startingNbs, number)
	}
	return startingNbs
}
