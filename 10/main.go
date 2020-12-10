package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("test1.txt")
	input = append(input, 0)
	sort.Ints(input)
	charger := input[len(input)-1] + 3
	input = append(input, charger)
	//fmt.Println(calculateJoltDifference(input))
	fmt.Println(calculateOmittion(input))
	fmt.Println(findCombinations(input, 0))
}

func findCombinations(adapters []int, total int) int {
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

func calculateOmittion(adapters []int) int {
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
	fmt.Println(canBeOmitted)
	return count
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
