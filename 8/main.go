package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Operation string
	Argument  int
}

func main() {
	input := parseInput("data.txt")
	//fmt.Println(runCodeUntilLoop(input))
	possibleInstructions := createVariationsOfCode(input)
	for _, instructions := range possibleInstructions {
		//fmt.Printf("Instructions: %v\n", instructions)
		result := runCode(instructions)
		if result != -1 {
			fmt.Println(result)
		}
	}
}

func runCodeUntilLoop(instructions []Instruction) int {
	accumulator := 0
	alreadyExecuted := []int{}
	for i := 0; i < len(instructions); i++ {
		if contains(alreadyExecuted, i) {
			return accumulator
		}
		instruct := instructions[i]
		switch instruct.Operation {
		case "nop":
			alreadyExecuted = append(alreadyExecuted, i)
			continue
		case "acc":
			accumulator += instruct.Argument
			alreadyExecuted = append(alreadyExecuted, i)
		case "jmp":
			alreadyExecuted = append(alreadyExecuted, i)
			newIndex := i + instruct.Argument - 1
			i = newIndex
		}
	}
	return -1
}

func runCode(instructions []Instruction) int {
	accumulator := 0
	alreadyExecuted := []int{}
	for i := 0; i < len(instructions); i++ {
		if contains(alreadyExecuted, i) {
			return -1
		}
		instruct := instructions[i]
		switch instruct.Operation {
		case "nop":
			alreadyExecuted = append(alreadyExecuted, i)
			continue
		case "acc":
			accumulator += instruct.Argument
			alreadyExecuted = append(alreadyExecuted, i)
		case "jmp":
			alreadyExecuted = append(alreadyExecuted, i)
			newIndex := i + instruct.Argument - 1
			i = newIndex
		}
	}
	return accumulator
}

func createVariationsOfCode(instructions []Instruction) [][]Instruction {
	possibleCodes := [][]Instruction{}
	for i, code := range instructions {
		tempInstructions := []Instruction{}
		tempInstructions = append(tempInstructions, instructions...)
		if code.Operation == "nop" {
			tempInstructions[i].Operation = "jmp"
			possibleCodes = append(possibleCodes, tempInstructions)
		} else if code.Operation == "jmp" {
			tempInstructions[i].Operation = "nop"
			possibleCodes = append(possibleCodes, tempInstructions)
		}
	}
	return possibleCodes
}

func contains(arr []int, val int) bool {
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func parseInput(href string) []Instruction {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	individualInstructions := strings.Split(input, "\n")
	instructions := []Instruction{}
	for _, instruct := range individualInstructions {
		opAndArg := strings.Split(instruct, " ")
		operation := opAndArg[0]
		argument, _ := strconv.Atoi(opAndArg[1])
		instructions = append(instructions, Instruction{Operation: operation, Argument: argument})
	}
	return instructions
}
