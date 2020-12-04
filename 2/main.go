package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input := parseInput()
	validPasswords := 0
	for _, pass := range input {
		components := strings.Split(pass, " ")
		charRange := strings.Split(components[0], "-")
		letter := string(components[1][0])
		fmt.Println(letter)
		password := components[2]
		if newPasswordTest(charRange, letter, password) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}

func testPassword(testRange []string, letter string, password string) bool {
	lowerbound, _ := strconv.Atoi(testRange[0])
	upperBound, _ := strconv.Atoi(testRange[1])
	count := 0
	for _, char := range password {
		if string(char) == letter {
			count++
		}
	}
	if count < lowerbound || count > upperBound {
		return false
	}
	return true
}

func newPasswordTest(testRange []string, letter string, password string) bool {
	idx1, _ := strconv.Atoi(testRange[0])
	idx2, _ := strconv.Atoi(testRange[1])
	isAt1 := false
	isAt2 := false
	if string(password[idx1-1]) == letter {
		isAt1 = true
	}
	if string(password[idx2-1]) == letter {
		isAt2 = true
	}
	if (isAt1 && !isAt2) || (isAt2 && !isAt1) {
		return true
	}
	return false
}

func parseInput() []string {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	passwordEntries := strings.Split(input, "\n")
	return passwordEntries
}
