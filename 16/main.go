package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Data struct {
	Fields        map[string][][]int
	YourTicket    []int
	NearbyTickets [][]int
}

func main() {
	input := parseInput("data.txt")
	//fmt.Println(ticketScanningError(input))
	discarded := discardTicketErrors(input)
	//fmt.Println(discarded)
	maps := identifyFields(discarded)
	fmt.Println(maps)
	fmt.Println(len(maps))
	fmt.Println(calculateYourTicket(discarded, maps))
}

func calculateYourTicket(data Data, fieldIdxMap map[int]string) int {
	yourTicket := data.YourTicket
	product := 1
	for i := 0; i < len(yourTicket); i++ {
		currentVal := yourTicket[i]
		currentField := fieldIdxMap[i]
		if len(currentField) > 9 && currentField[:9] == "departure" {
			fmt.Printf("curV: %d, curF: %s\n", currentVal, currentField)
			product *= currentVal
		}
	}
	return product
}

func identifyFields(data Data) map[int]string {
	fields := data.Fields
	nearTickets := data.NearbyTickets
	ticketLength := len(nearTickets[0])
	fieldIdxMap := map[int]string{}
	possibleIdx := map[string][]int{}
	for fieldName, field := range fields {
		indeces := []int{}
		for i := 0; i < ticketLength; i++ {
			validForAllVals := true
			for _, ticket := range nearTickets {
				currentVal := ticket[i]
				if !testSingleField(field, currentVal) {
					validForAllVals = false
					break
				}
			}
			if validForAllVals {
				indeces = append(indeces, i)
			}
		}
		//fmt.Printf("field: %s, indeces: %d\n", fieldName, indeces)
		possibleIdx[fieldName] = indeces
	}
	takenIdx := []int{}
	for i := 1; i < ticketLength; i++ {
		for field, indeces := range possibleIdx {
			if len(indeces) == 1 {
				takenIdx = append(takenIdx, indeces[0])
				fieldIdxMap[indeces[0]] = field
			} else if len(indeces) == i {
				for _, value := range indeces {
					if !contains(takenIdx, value) {
						takenIdx = append(takenIdx, value)
						fieldIdxMap[value] = field
					}
				}
			}
		}
	}
	return fieldIdxMap
}

func contains(arr []int, n int) bool {
	for _, val := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func testSingleField(field [][]int, value int) bool {
	for _, rng := range field {
		lowerB := rng[0]
		upperB := rng[1]
		if value <= upperB && value >= lowerB {
			return true
		}
	}
	return false
}

func discardTicketErrors(data Data) Data {
	fields := data.Fields
	nearTickets := data.NearbyTickets
	newNearT := [][]int{}
	for _, ticket := range nearTickets {
		allValsValid := true
		for _, val := range ticket {
			if !testTicketValue(fields, val) {
				allValsValid = false
			}
		}
		if allValsValid {
			newNearT = append(newNearT, ticket)
		}
	}
	newData := Data{Fields: fields, YourTicket: data.YourTicket, NearbyTickets: newNearT}
	return newData
}

func ticketScanningError(data Data) int {
	fields := data.Fields
	nearTickets := data.NearbyTickets
	count := 0
	for _, ticket := range nearTickets {
		for _, val := range ticket {
			if !testTicketValue(fields, val) {
				count += val
			}
		}
	}
	return count
}

func testTicketValue(fields map[string][][]int, value int) bool {
	for _, field := range fields {
		for _, rng := range field {
			lowerB := rng[0]
			upperB := rng[1]
			if value <= upperB && value >= lowerB {
				return true
			}
		}
	}
	return false
}

func parseInput(href string) Data {
	data, _ := ioutil.ReadFile(href)
	input := string(data)
	splitByLine := strings.Split(input, "\n")
	fields := []string{}
	yourTicket := ""
	nearTickets := []string{}
	for i, line := range splitByLine {
		if line == "" {
			fields = append(fields, splitByLine[:i]...)
			yourTicket = splitByLine[i+2]
			nearTickets = append(nearTickets, splitByLine[i+5:]...)
			break
		}
	}
	fieldMap := map[string][][]int{}
	yourTicketVal := []int{}
	nearTicketVal := [][]int{}
	for _, field := range fields {
		colonSplit := strings.Split(field, ": ")
		fieldName := colonSplit[0]
		ranges := strings.Split(colonSplit[1], " or ")
		for _, r := range ranges {
			values := strings.Split(r, "-")
			lowerB, _ := strconv.Atoi(values[0])
			upperB, _ := strconv.Atoi(values[1])
			fieldMap[fieldName] = append(fieldMap[fieldName], []int{lowerB, upperB})
		}
	}
	yourTicketSplit := strings.Split(yourTicket, ",")
	for _, val := range yourTicketSplit {
		value, _ := strconv.Atoi(val)
		yourTicketVal = append(yourTicketVal, value)
	}
	for _, nearT := range nearTickets {
		values := strings.Split(nearT, ",")
		newTicket := []int{}
		for _, val := range values {
			value, _ := strconv.Atoi(val)
			newTicket = append(newTicket, value)
		}
		nearTicketVal = append(nearTicketVal, newTicket)
	}
	allData := Data{Fields: fieldMap, YourTicket: yourTicketVal, NearbyTickets: nearTicketVal}
	return allData
}
