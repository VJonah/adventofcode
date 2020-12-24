package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("test.txt")
	flipped := runInstructions(input)
	//fmt.Println(input)
	fmt.Println(countBlackTiles(flipped))
	fmt.Println(countBlackTiles(updateTiles(flipped)))
}

func updateTiles(tiles map[string]bool) map[string]bool {
	newTiles := map[string]bool{}
	for pos, tile := range tiles {
		adjacentCount := getAdjacentTiles(tiles, pos)
		//fmt.Println(adjacentCount)
		newTile := false
		if tile {
			if adjacentCount != 0 && adjacentCount <= 2 {
				newTile = tile
			}
		} else {
			if adjacentCount == 2 {
				newTile = true
			}
		}
		newTiles[pos] = newTile
	}
	return newTiles
}

func getAdjacentTiles(tiles map[string]bool, pos string) int {
	directions := map[string][]int{
		"ne": []int{1, 2},
		"nw": []int{-1, 2},
		"sw": []int{-1, -2},
		"se": []int{1, -2},
		"e":  []int{2, 0},
		"w":  []int{-2, 0},
	}
	position := stringToPosition(pos)
	numberOfBlackTiles := 0
	for _, direct := range directions {
		adjacentX := direct[0] + position[0]
		adjacentY := direct[1] + position[1]
		adjacent := []int{adjacentX, adjacentY}
		adjacentPos := positionToString(adjacent)
		if tiles[adjacentPos] {
			numberOfBlackTiles++
		}
	}
	return numberOfBlackTiles
}

func countBlackTiles(tiles map[string]bool) int {
	count := 0
	for _, tile := range tiles {
		if tile {
			count++
		}
	}
	return count
}

func runInstructions(instructions [][]string) map[string]bool {
	tiles := map[string]bool{}
	directions := map[string][]int{
		"ne": []int{1, 2},
		"nw": []int{-1, 2},
		"sw": []int{-1, -2},
		"se": []int{1, -2},
		"e":  []int{2, 0},
		"w":  []int{-2, 0},
	}
	for _, instruct := range instructions {
		position := []int{0, 0}
		for _, direction := range instruct {
			update := directions[direction]
			position[0] += update[0]
			position[1] += update[1]
		}
		result := positionToString(position)
		tiles[result] = flipTile(tiles[result])
	}
	return tiles
}

func stringToPosition(pos string) []int {
	numbers := strings.Split(pos, ",")
	x, _ := strconv.Atoi(numbers[0])
	y, _ := strconv.Atoi(numbers[1])
	return []int{x, y}
}

func positionToString(pos []int) string {
	x := strconv.Itoa(pos[0])
	y := strconv.Itoa(pos[1])
	return x + "," + y
}

func flipTile(state bool) bool {
	if state {
		return false
	}
	return true
}

func parseInput(href string) [][]string {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	instructions := [][]string{}
	for _, line := range lines {
		lineInstruct := []string{}
		for i := 0; i < len(line); i++ {
			current := string(line[i])
			next := "x"
			if i < len(line)-1 {
				next = string(line[i+1])
			}
			instruction := ""
			if current == "s" || current == "n" {
				if next == "e" || next == "w" {
					instruction = current + next
					i++
				}
			} else {
				instruction = current
			}
			lineInstruct = append(lineInstruct, instruction)
		}
		instructions = append(instructions, lineInstruct)
	}
	return instructions
}
