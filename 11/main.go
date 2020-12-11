package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	//fmt.Println(len(input[0]))
	stableState := simulateUntilStill(input)
	fmt.Println(countOccupiedSeats(stableState))
	//fmt.Println(getFirstSeat(2, 7, input))
}

func countOccupiedSeats(state []string) int {
	count := 0
	for _, row := range state {
		for _, char := range row {
			if string(char) == "#" {
				count++
			}
		}
	}
	return count
}

func simulateUntilStill(start []string) []string {
	previousState := start
	for {
		currentState := evaluateNewSeatStates(previousState)
		if compareStates(currentState, previousState) {
			break
		}
		previousState = []string{}
		previousState = append(previousState, currentState...)
	}
	return previousState
}

func compareStates(first, second []string) bool {
	for i, row := range first {
		if row != second[i] {
			return false
		}
	}
	return true
}

func evaluateNewSeatStates(currentState []string) []string {
	newState := []string{}
	for i, row := range currentState {
		newRow := ""
		for j, char := range row {
			//adjacentCells := getAdjacentCells(i, j, currentState)
			adjacentCells := getFirstSeat(i, j, currentState)
			typeCount := countCellTypes(adjacentCells)
			newSeatState := string(char)
			switch string(char) {
			case "L":
				if typeCount["#"] == 0 {
					newSeatState = "#"
				}
			case "#":
				if typeCount["#"] >= 5 {
					newSeatState = "L"
				}
			}
			newRow += newSeatState
		}
		newState = append(newState, newRow)
	}
	return newState
}

func countCellTypes(adjacentCells string) map[string]int {
	typeCount := map[string]int{"#": 0, "L": 0}
	for _, char := range adjacentCells {
		if string(char) == "#" {
			typeCount["#"]++
		}
		if string(char) == "L" {
			typeCount["L"]++
		}
	}
	return typeCount
}

func getFirstSeat(row int, col int, seats []string) string {
	possibleDirections := [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}
	firstSeen := ""
	numberOfRows := len(seats)
	numberOfCols := len(seats[0])
	for _, direction := range possibleDirections {
		i := row
		j := col
		currentSeat := "."
		for currentSeat == "." {
			i += direction[1]
			j += direction[0]
			if (i >= 0 && j >= 0) && (i < numberOfRows && j < numberOfCols) {
				currentSeat = string(seats[i][j])
			} else {
				break
			}
		}
		if currentSeat != "." {
			firstSeen += currentSeat
		}
	}
	return firstSeen
}

func getAdjacentCells(row int, col int, seats []string) string {
	adjacent := ""
	rowLength := len(seats)
	colLength := len(seats[0])
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			//check within index of array
			if (i >= 0 && j >= 0) && (i < rowLength && j < colLength) {
				//checks if not the index of the value we want the adjacent cells to
				if !(i == row && j == col) {
					//fmt.Printf("colLength: %d, i: %d, j: %d\n", colLength, i, j)
					adjacent += string(seats[i][j])
				}

			}

		}
	}
	return adjacent
}

func parseInput(href string) []string {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	lines := strings.Split(input, "\n")
	return lines
}
