package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	groups := parseInput()
	sum := 0
	for _, group := range groups {
		answers := strings.Split(group, " ")
		previousSeen := []string{}
		//count := 0
		for i, ans := range answers {
			newSeen := []string{}
			for _, char := range ans {
				if i == 0 {
					newSeen = append(newSeen, string(char))
				} else if contains(string(char), previousSeen) {
					newSeen = append(newSeen, string(char))
				}
			}
			fmt.Printf("group: %s\n ans: %s \n previous: %s, new: %s\n", group, ans, previousSeen, newSeen)
			previousSeen = newSeen
		}
		sum += len(previousSeen)
	}
	fmt.Println(sum)
}

func contains(s string, arr []string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func parseInput() []string {
	data, _ := ioutil.ReadFile("data.txt")
	input := string(data)
	splitInput := strings.Split(input, "\n")
	passGroups := make([]string, len(splitInput))
	i := 0
	numberOfGroups := 1
	for _, pass := range splitInput {
		if len(pass) == 0 {
			i++
			numberOfGroups++
		} else if len(passGroups[i]) == 0 {
			passGroups[i] += pass
		} else {
			passGroups[i] += " " + pass
		}
	}
	passGroups = passGroups[:numberOfGroups]
	return passGroups
}
