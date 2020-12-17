package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	//input := parseInput("test.txt")
	//simulated := simulateCells(input, 6)
	//fmt.Println(countCells(simulated))
	input := parseInput4D("data.txt")
	simulated := simulateCells(input, 6)
	fmt.Println(countCells(simulated))
}

func countCells(state [][][]string) int {
	count := 0
	for _, cube := range state {
		for _, slice := range cube {
			for _, row := range slice {
				for _, char := range row {
					if string(char) == "#" {
						count++
					}
				}
			}
		}
	}
	return count
}
func simulateCells(state [][][]string, cycles int) [][][]string {
	finalState := state
	for i := 0; i < cycles; i++ {
		finalState = updateCubes(finalState)
	}
	return finalState
}

func updateCubes(state [][][]string) [][][]string {
	bufferCube := generateBufferCube(state)
	newState := [][][]string{bufferCube}
	for l, hypercube := range state {
		newCube := updateSlices(hypercube, state, l)
		newState = append(newState, newCube)
	}
	newState = append(newState, bufferCube)
	return newState
}

func updateSlices(state [][]string, hyperState [][][]string, cubeIdx int) [][]string {
	bufferRow := generateBufferRow(len(state[0][0]))
	bufferSlice := generateBufferSlice(state)
	newState := [][]string{bufferSlice}
	for i, slice := range state {
		newSlice := []string{bufferRow}
		for j, row := range slice {
			newRow := "."
			for k, char := range row {
				newCell := "."
				adjacentCells := getAdjacentCells4D(k, j, i, cubeIdx, hyperState)
				activeCubes := countActiveCubes(adjacentCells)
				if string(char) == "#" && (activeCubes == 2 || activeCubes == 3) {
					newCell = "#"
				} else if string(char) != "#" && activeCubes == 3 {
					newCell = "#"
				}
				newRow += newCell
			}
			newRow += "."
			newSlice = append(newSlice, newRow)
		}
		newSlice = append(newSlice, bufferRow)
		newState = append(newState, newSlice)
	}
	newState = append(newState, bufferSlice)
	return newState
}

func generateBufferCube(state [][][]string) [][]string {
	newCube := [][]string{}
	sliceLen := len(state[0])
	sliceBuffers := generateBufferSlice(state[0])
	for i := 0; i < sliceLen+2; i++ {
		newCube = append(newCube, sliceBuffers)
	}
	return newCube
}

func generateBufferRow(rowLen int) string {
	result := ""
	for i := 0; i < rowLen+2; i++ {
		result += "."
	}
	return result
}

func generateBufferSlice(state [][]string) []string {
	rowLen := len(state[0][0])
	newRow := generateBufferRow(rowLen)
	sliceLen := len(state[0])
	bufferSlice := []string{}
	for i := 0; i < sliceLen+2; i++ {
		bufferSlice = append(bufferSlice, newRow)
	}
	return bufferSlice
}

func countActiveCubes(adjacentCells string) int {
	count := 0
	for _, char := range adjacentCells {
		if string(char) == "#" {
			count++
		}
	}
	return count
}

func getAdjacentCells4D(x, y, z, w int, state [][][]string) string {
	adjacent := ""
	wLen := len(state)
	for l := w - 1; l <= w+1; l++ {
		if l >= 0 && l < wLen {
			currentCube := state[l]
			if l == w {
				adjacent += getAdjacentCells(x, y, z, currentCube, true)
			} else {
				adjacent += getAdjacentCells(x, y, z, currentCube, false)
			}
		}
	}
	return adjacent
}

func getAdjacentCells(x, y, z int, state [][]string, isSameCube bool) string {
	adjacent := ""
	xLen := len(state[0][0])
	yLen := len(state[0])
	zLen := len(state)
	//repeat for adjacent slices
	for i := z - 1; i <= z+1; i++ {
		//repeat across layers of slice
		for j := y - 1; j <= y+1; j++ {
			//repeat across columns of slice
			for k := x - 1; k <= x+1; k++ {
				if (i >= 0 && j >= 0 && k >= 0) && (i < zLen && j < yLen && k < xLen) {
					if !(i == z && j == y && k == x && isSameCube) {
						/*if string(state[i][j][k]) == "#" {
							fmt.Println(i, j, k)
						}*/
						adjacent += string(state[i][j][k])
					}
				} else {
					adjacent += "."
				}
			}
		}
	}
	return adjacent
}

func parseInput(href string) [][]string {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	sliceAt0 := strings.Split(input, "\n")
	row := ""
	adjacentSlices := []string{}
	for i := 0; i < len(sliceAt0[0])+2; i++ {
		row += "."
	}
	for i := 0; i < len(sliceAt0)+2; i++ {
		adjacentSlices = append(adjacentSlices, row)
	}
	for i, _ := range sliceAt0 {
		sliceAt0[i] = "." + sliceAt0[i] + "."
	}
	sliceAt0 = append(sliceAt0, row)
	finalSlice := []string{row}
	finalSlice = append(finalSlice, sliceAt0...)
	state := [][]string{adjacentSlices, finalSlice, adjacentSlices}
	return state
	//return [][]string{sliceAt0}
}

func parseInput4D(href string) [][][]string {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	lineSplit := strings.Split(input, "\n")
	row := ""
	slices := []string{}
	adjacentCubes := [][]string{}
	for i := 0; i < len(lineSplit[0])+2; i++ {
		row += "."
	}
	for i := 0; i < len(lineSplit)+2; i++ {
		slices = append(slices, row)
	}
	for i := 0; i < 3; i++ {
		adjacentCubes = append(adjacentCubes, slices)
	}
	for i, _ := range lineSplit {
		lineSplit[i] = "." + lineSplit[i] + "."
	}
	finalSlice := []string{row}
	finalSlice = append(finalSlice, lineSplit...)
	finalSlice = append(finalSlice, row)
	finalCube := [][]string{slices, finalSlice, slices}
	state := [][][]string{adjacentCubes, finalCube, adjacentCubes}
	return state
}
