package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("data.txt")
	//startTime := input[0][0]
	timestamps := input[1]
	//fmt.Println(timestamps)
	//fmt.Println(findEarliestBus(startTime, timestamps))
	fmt.Println(findEarliestTime(timestamps))
}

func findEarliestTime(timestamps []int) int {
	timestampAndIndex := [][]int{}
	for i, time := range timestamps {
		if time != -1 {
			timestampAndIndex = append(timestampAndIndex, []int{time, i})
		}
	}
	increaseBy := timestamps[0]
	count := increaseBy
	for i := 0; i < len(timestampAndIndex)-1; i++ {
		nextTime := timestampAndIndex[i+1][0]
		//fmt.Printf("currentTime: %d, nextTime:%d \n", currentTime, nextTime)
		nextTimeIdx := timestampAndIndex[i+1][1]
		for j := count; j > 0; j += increaseBy {
			//fmt.Printf("count: %d, j: %d, increaseBy: %d, jmodNxtT: %d\n", count, j, increaseBy, j%nextTime)
			if nextTime-(j%nextTime) == nextTimeIdx%nextTime {
				count = j
				increaseBy *= nextTime
				break
			}
		}
	}
	return count
}

func findEarliestBus(startTime int, timestamps []int) int {
	busID := timestamps[0]
	earliestBus := timestamps[0] - (startTime % timestamps[0])
	for i := 1; i < len(timestamps); i++ {
		bus := timestamps[i]
		if bus != -1 {
			currentEarliest := timestamps[i] - (startTime % timestamps[i])
			if currentEarliest < earliestBus {
				earliestBus = currentEarliest
				busID = bus
			}
		}
	}
	return busID * earliestBus
}

func parseInput(href string) [][]int {
	earliestAndTimestamps := [][]int{}
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	splitByLine := strings.Split(input, "\n")
	earliestTime, _ := strconv.Atoi(splitByLine[0])
	earliestAndTimestamps = append(earliestAndTimestamps, []int{earliestTime})
	timestamps := strings.Split(splitByLine[1], ",")
	onlyValidBuses := []int{}
	for _, time := range timestamps {
		if string(time) != "x" {
			timestamp, _ := strconv.Atoi(time)
			onlyValidBuses = append(onlyValidBuses, timestamp)
		} else {
			onlyValidBuses = append(onlyValidBuses, -1)
		}
	}
	earliestAndTimestamps = append(earliestAndTimestamps, onlyValidBuses)
	return earliestAndTimestamps
}
