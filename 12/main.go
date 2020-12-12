package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction string
	Value     int
}

func main() {
	input := parseInput("data.txt")
	//fmt.Println(calculateManhattanLen(input))
	fmt.Println(calculateWaypointManhattan(input))
}

func calculateWaypointManhattan(instructions []Instruction) int {
	waypoint := map[string]int{"N": 1, "S": 0, "E": 10, "W": 0}
	boat := map[string]int{"N": 0, "S": 0, "E": 0, "W": 0}
	for _, instruction := range instructions {
		//fmt.Printf("waypoint: %v\nboat: %v\n \n", waypoint, boat)
		switch instruction.Direction {
		case "L":
			waypoint = rotateWaypoint(waypoint, instruction.Value, false)
		case "R":
			waypoint = rotateWaypoint(waypoint, instruction.Value, true)
		case "F":
			for direction, val := range waypoint {
				if val != 0 {
					boat[direction] += val * instruction.Value
				}
			}
		default:
			waypoint[instruction.Direction] += instruction.Value
		}
	}
	eastWest := boat["E"] - boat["W"]
	northSouth := boat["N"] - boat["S"]
	if eastWest < 0 {
		eastWest *= -1
	}
	if northSouth < 0 {
		northSouth *= -1
	}
	return eastWest + northSouth
}

func rotateWaypoint(positioning map[string]int, angle int, isRight bool) map[string]int {
	newPosition := map[string]int{"N": 0, "S": 0, "E": 0, "W": 0}
	for direction, val := range positioning {
		if val != 0 {
			newPosition[getFacingDirection(direction, angle, isRight)] = val
		}
	}
	return newPosition
}

func calculateManhattanLen(instructions []Instruction) int {
	eastWest := 0
	northSouth := 0
	currentFacing := "E"
	//using a map for the values of the directions would make things nicer
	for _, instruction := range instructions {
		//fmt.Printf("eastW: %d, northS: %d, currentF: %s\n", eastWest, northSouth, currentFacing)
		switch instruction.Direction {
		case "N":
			northSouth += instruction.Value
		case "S":
			northSouth -= instruction.Value
		case "E":
			eastWest += instruction.Value
		case "W":
			eastWest -= instruction.Value
		case "L":
			currentFacing = getFacingDirection(currentFacing, instruction.Value, false)
		case "R":
			currentFacing = getFacingDirection(currentFacing, instruction.Value, true)
		case "F":
			switch currentFacing {
			case "N":
				northSouth += instruction.Value
			case "S":
				northSouth -= instruction.Value
			case "E":
				eastWest += instruction.Value
			case "W":
				eastWest -= instruction.Value
			}
		}
	}
	if eastWest < 0 {
		eastWest *= -1
	}
	if northSouth < 0 {
		northSouth *= -1
	}
	return eastWest + northSouth
}

func getFacingDirection(currentFacing string, angle int, isRight bool) string {
	directions := map[int]string{0: "N", 90: "E", 180: "S", 270: "W"}
	var currentAngle int
	for angle, direction := range directions {
		if direction == currentFacing {
			currentAngle = angle
		}
	}
	if isRight {
		currentAngle += angle
	} else {
		currentAngle -= angle
	}
	currentAngle = currentAngle % 360
	if currentAngle < 0 {
		currentAngle = 360 + currentAngle
	}
	return directions[currentAngle]
}

func parseInput(href string) []Instruction {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	splitByLine := strings.Split(input, "\n")
	instructions := []Instruction{}
	for _, instruction := range splitByLine {
		direction := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])
		newInstruct := Instruction{Direction: direction, Value: value}
		instructions = append(instructions, newInstruct)
	}
	return instructions
}
