package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	invalidNumber := findNonSum(input, 25)
	contiguousSet := sumToInvalid(input, invalidNumber)
	sort.Ints(contiguousSet)
	fmt.Printf("sum: %d\n", contiguousSet[0]+contiguousSet[len(contiguousSet)-1])
}

func sumToInvalid(numbers []int, invalidN int) []int {
	for i, n := range numbers {
		sum := n
		set := []int{n}
		for j := i + 1; j < len(numbers); j++ {
			currentN := numbers[j]
			sum += currentN
			if sum > invalidN {
				break
			}
			set = append(set, currentN)
			if sum == invalidN {
				return set
			}
		}
	}
	return []int{}
}

func findNonSum(numbers []int, preambleLength int) int {
	for i := preambleLength; i < len(numbers); i++ {
		subsection := numbers[i-preambleLength : i]
		currentNumber := numbers[i]
		if !checkIfValid(subsection, currentNumber) {
			return currentNumber
		}
	}
	return -1
}

func checkIfValid(subSection []int, val int) bool {
	for _, n := range subSection {
		difference := val - n
		if difference != val && contains(subSection, difference) {
			return true
		}
	}
	return false
}

func contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func parseInput(href string) []int {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	numbersAsStrings := strings.Split(input, "\n")
	numbers := []int{}
	for _, n := range numbersAsStrings {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
	}
	return numbers
}
