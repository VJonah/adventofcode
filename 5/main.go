package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	boardingPasses := parseInput()
	highest := 0
	fmt.Printf("Test gives: %d", calculateID("FBFBBFF", "RLR"))
	ids := []int{}
	for _, pass := range boardingPasses {
		row := pass[:7]
		column := pass[7:]
		currentID := calculateID(row, column)
		ids = append(ids, currentID)
		if currentID > highest {
			highest = currentID
		}
		if currentID == -1 {
			fmt.Println("-1 has occured")
		}
	}
	sort.Ints(ids)
	fmt.Println(ids)
	seatID := 0
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] == 2 {
			seatID = ids[i] - 1
		}
	}
	fmt.Println(seatID)
	//fmt.Println(highest)
}

func calculateID(row, column string) int {
	rowRange := []int{0, 127}
	columnRange := []int{0, 7}
	//fmt.Printf("row: %s, col: %s\n", row, column)
	for _, letter := range row {
		mid := int(math.Floor(float64((rowRange[0] + rowRange[1]) / 2)))
		if string(letter) == "F" {
			rowRange[1] = mid
		} else {
			rowRange[0] = mid + 1
		}
	}
	for _, letter := range column {
		mid := int(math.Floor(float64((columnRange[0] + columnRange[1]) / 2)))
		if string(letter) == "L" {
			columnRange[1] = mid
		} else {
			columnRange[0] = mid + 1
		}
		//fmt.Printf("col: %s, mid: %d, colR: %d\n", column, mid, columnRange)
	}
	if columnRange[0] != columnRange[1] || rowRange[0] != rowRange[1] {
		return -1
	}
	return rowRange[0]*8 + columnRange[0]
}

func parseInput() []string {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	seats := strings.Split(input, "\n")
	return seats
}
