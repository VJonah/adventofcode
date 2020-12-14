package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Write struct {
	Mask    *string
	Address int
	Value   int
}

func main() {
	input := parseInput("data.txt")
	fmt.Println(input)
	memoryState := runProgram2(input)
	//fmt.Println(sumOfMemory(memoryState))
	fmt.Println(sumOfMemory(memoryState))
	//fmt.Println(findPossibleValues(test, []string{}))
}

func sumOfMemory(memory map[int]int) int {
	sum := 0
	for _, address := range memory {
		sum += address
	}
	return sum
}

func runProgram(program []Write) map[int]int {
	memory := map[int]int{}
	for _, write := range program {
		applied := applyMask(write.Mask, write.Value)
		memory[write.Address] = applied
	}
	return memory
}

func runProgram2(program []Write) map[int]int {
	memory := map[int]int{}
	for i, write := range program {
		appliedAddr := applyAddressMask(write.Mask, write.Address)
		fmt.Println(i)
		//fmt.Println(appliedAddr)
		for _, addr := range appliedAddr {
			memory[addr] = write.Value
		}
	}
	return memory
}

func applyAddressMask(mask *string, address int) []int {
	byteValue := string(strconv.FormatInt(int64(address), 2))
	byteLength := len(byteValue)
	masked := ""
	byteIdx := byteLength - 1
	maskString := *mask
	for i := len(*mask) - 1; i > -1; i-- {
		current := string(maskString[i])
		if current == "0" {
			if byteIdx > -1 {
				masked = string(byteValue[byteIdx]) + masked
			} else {
				masked = "0" + masked
			}
		} else {
			masked = current + masked
		}
		byteIdx--
	}
	fmt.Println(masked)
	maskedVariations := findPossibleValues(masked, []string{})
	addresses := []int{}
	for _, variation := range maskedVariations {
		addresses = append(addresses, binaryToInt(variation))
	}
	return addresses
}

func findPossibleValues(byteS string, variations []string) []string {
	if !contains(byteS, "X") {
		return []string{byteS}
	}
	for i := len(byteS) - 1; i > -1; i-- {
		current := string(byteS[i])
		if current == "X" {
			if contains(byteS[i+1:], "X") {
				return variations
			}
			next := findPossibleValues(byteS[:i], variations)
			variation := ""
			for _, possibility := range next {
				variation = possibility + "1" + byteS[i+1:]
				//fmt.Printf("poss: %s, byteSsec: %s, var: %s\n", possibility, byteS, variation)
				variations = append(variations, variation)
				variation = possibility + "0" + byteS[i+1:]
				variations = append(variations, variation)
			}
		}
	}
	return variations
}

func contains(s, val string) bool {
	for _, char := range s {
		if string(char) == val {
			return true
		}
	}
	return false
}

func applyMask(mask *string, value int) int {
	byteValue := string(strconv.FormatInt(int64(value), 2))
	byteLength := len(byteValue)
	masked := ""
	byteIdx := byteLength - 1
	maskString := *mask
	for i := len(*mask) - 1; i > -1; i-- {
		current := string(maskString[i])
		if current != "X" {
			masked = current + masked
		} else {
			if byteIdx > -1 {
				masked = string(byteValue[byteIdx]) + masked
			} else {
				masked = "0" + masked
			}
		}
		byteIdx--
	}
	return binaryToInt(masked)
}

func binaryToInt(byteS string) int {
	val := 0
	exponent := 0.0
	for i := len(byteS) - 1; i > -1; i-- {
		digit := string(byteS[i])
		power := int(math.Pow(2.0, exponent))
		digitVal, _ := strconv.Atoi(digit)
		val += digitVal * power
		exponent++
	}
	return val
}
func parseInput(href string) []Write {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	splitByMask := strings.Split(input, "mask = ")
	splitByMask = splitByMask[1:]
	program := []Write{}
	for _, section := range splitByMask {
		eachLine := strings.Split(section, "\n")
		tempMask := eachLine[0]
		for i := 1; i < len(eachLine); i++ {
			currentWrite := eachLine[i]
			if currentWrite != "" {
				splitOperation := strings.Split(currentWrite, " = ")
				address, _ := strconv.Atoi(splitOperation[0][4 : len(splitOperation[0])-1])
				value, _ := strconv.Atoi(splitOperation[1])
				write := Write{Mask: &tempMask, Address: address, Value: value}
				program = append(program, write)
			}
		}
	}
	return program
}
