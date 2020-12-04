package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	linears := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	input := parseInput()
	product := 1
	for _, eq := range linears {
		product *= countTrees(eq[0], eq[1], input)
	}
	fmt.Printf("product: %d\n", product)
}

func countTrees(xVal int, yVal int, trees []string) int {
	x := 0
	y := 0
	count := 0
	width := len(trees[0])
	for y < len(trees)-1 {
		x += xVal
		y += yVal
		if x > width-1 {
			x = x % width
		}
		currentSpot := string(trees[y][x])
		fmt.Println(currentSpot)
		if currentSpot == "#" {
			count++
		}
	}
	return count
}

func parseInput() []string {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	trees := strings.Split(input, "\n")
	return trees
}
