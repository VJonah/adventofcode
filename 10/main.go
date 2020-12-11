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
	input = append(input, 0)
	sort.Ints(input)
	charger := input[len(input)-1] + 3
	input = append(input, charger)
	//fmt.Println(calculateJoltDifference(input))
	optionalOnes := optionalAdapters(input)
	fmt.Println(findCombinations(optionalOnes))
}

/*func findCombinations(adapters []int, total int) int {
	canBeOmitted := calculateOmittion(adapters)
	//fmt.Printf("adapters: %d, total: %d\n", adapters, total)
	if canBeOmitted == 0 {
		total++
		return total
	}
	total++
	for i := len(adapters) - 3; i > 0; i-- {
		previous := adapters[i+1]
		next := adapters[i-1]
		if previous-next < 3 {
			newArr := adapters[i:]
			total = findCombinations(newArr, total)
		}
	}
	return total
}
*/

func findCombinations(optionalAdapters []int) int {
	total := 1
	for _, adapter := range optionalAdapters {
		if contains(optionalAdapters, adapter+1) && contains(optionalAdapters, adapter+2) {
			total += 3 * total / 4
		} else {
			total += total
		}
	}
	return total
}

func contains(arr []int, val int) bool {
	for _, n := range arr {
		if val == n {
			return true
		}
	}
	return false
}

func optionalAdapters(adapters []int) []int {
	count := 0
	canBeOmitted := []int{}
	for i := len(adapters) - 3; i > 0; i-- {
		previous := adapters[i+1]
		next := adapters[i-1]
		current := adapters[i]
		if previous-next <= 3 {
			canBeOmitted = append(canBeOmitted, current)
			count++
		}
	}
	//fmt.Println(canBeOmitted)
	return canBeOmitted
}

func calculateJoltDifference(adapters []int) int {
	threeJ := 0
	oneJ := 0
	for i := 0; i < len(adapters)-1; i++ {
		current := adapters[i]
		next := adapters[i+1]
		difference := next - current
		if difference == 1 {
			oneJ++
		} else if difference == 3 {
			threeJ++
		}
	}
	return oneJ * threeJ
}

func parseInput(href string) []int {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	adapters := strings.Split(input, "\n")
	adaptersInt := []int{}
	for _, adapt := range adapters {
		jolts, _ := strconv.Atoi(adapt)
		adaptersInt = append(adaptersInt, jolts)
	}

	return adaptersInt
}
