package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	Index   int
	Content []string
}

type Data struct {
	Rules    []Rule
	Messages []string
}

func main() {
	input := parseInput("test.txt")
	//fmt.Println(input)
	fmt.Println(whatMapsToIt(input[0]))
}

func whatMapsToIt(rules []string) map[int][]int {
	whatMapToRule := map[int][]int{}
	for i, rule := range rules {
		seen := []int{}
		if string(rule[3]) != `"` {
			content := strings.Split(rule[3:], " ")
			for _, val := range content {
				if val != "|" {
					number, _ := strconv.Atoi(val)
					if !contains(seen, number) {
						seen = append(seen, number)
						whatMapToRule[number] = append(whatMapToRule[number], i)
					}
				}
			}
		}
	}
	return whatMapToRule
}

func contains(arr []int, n int) bool {
	for _, val := range arr {
		if n == val {
			return true
		}
	}
	return false
}

func indexMap(rules []string) map[int]string {
	idxMp := map[int]string{}
	for _, rule := range rules {
		index, _ := strconv.Atoi(string(rule[0]))
		idxMp[index] = rule[3:]
	}
	return idxMp
}

func parseInput(href string) [][]string {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	rules := []string{}
	messages := []string{}
	for i, line := range lines {
		if len(line) == 0 {
			rules = lines[:i]
			messages = lines[i+1:]
			break
		}
	}
	result := [][]string{rules, messages}
	return result
}
