package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	passports := parseInput()
	count := 0
	for _, pass := range passports {
		switch len(pass) {
		case 8:
			if testCases(pass) {
				count++
			}
		case 7:
			_, ok := pass["cid"]
			if !ok {
				if testCases(pass) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func testCases(pass map[string]string) bool {
	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	passed := 0
	fmt.Println(pass)
	if len(pass) == 8 {
		passed = 1
	}
	for key, val := range pass {
		fmt.Println(passed)
		switch key {
		case "byr":
			value, _ := strconv.Atoi(val)
			if value >= 1920 && value <= 2020 {
				fmt.Println("byr passed")
				passed++
			}
		case "iyr":
			value, _ := strconv.Atoi(val)
			if value >= 2010 && value <= 2020 {
				fmt.Println("iyr passed")
				passed++
			}
		case "eyr":
			value, _ := strconv.Atoi(val)
			if value >= 2020 && value <= 2030 {
				fmt.Println("eyr passed")
				passed++
			}
		case "hgt":
			switch len(val) {
			case 5:
				height, _ := strconv.Atoi(val[:3])
				if val[3:] == "cm" && (height >= 150 && height <= 193) {
					fmt.Println("hgt passed")
					passed++
				}
			case 4:
				height, _ := strconv.Atoi(val[:2])
				if val[2:] == "in" && (height >= 59 && height <= 76) {
					fmt.Println("hgt passed")
					passed++
				}
			}
		case "hcl":
			if len(val) == 7 && string(val[0]) == "#" {
				fmt.Println("hcl passed")
				passed++
			}
		case "ecl":
			if contains(eyeColors, val) {
				fmt.Println("ecl passed")
				passed++
			}
		case "pid":
			_, err := strconv.Atoi(val)
			if err == nil && len(val) == 9 {
				fmt.Println("pid passed")
				passed++
			}
		}
	}
	if passed == len(pass) {
		return true
	}
	return false
}

func contains(arr []string, val string) bool {
	for _, element := range arr {
		if element == val {
			return true
		}
	}
	return false
}

func parseInput() []map[string]string {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	passports := strings.Split(input, "\n")
	passportMaps := []map[string]string{}
	actualPassports := make([]string, len(passports))
	idx := 0
	numberOfPassports := 1
	for _, pass := range passports {
		if len(pass) == 0 {
			idx++
			numberOfPassports++
		} else if len(actualPassports[idx]) == 0 {
			actualPassports[idx] += pass
		} else {
			actualPassports[idx] += " " + pass
		}
	}
	actualPassports = actualPassports[:numberOfPassports]
	for _, pass := range actualPassports {
		components := strings.Split(pass, " ")
		tempMap := map[string]string{}
		for _, part := range components {
			tempParts := strings.Split(part, ":")
			tempMap[tempParts[0]] = tempParts[1]
		}
		passportMaps = append(passportMaps, tempMap)
	}
	return passportMaps
}
