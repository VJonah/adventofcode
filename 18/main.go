package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"./evaluator"
	"./lexer"
	"./parser"
)

func main() {
	input := parseInput("data.txt")
	fmt.Println(sumCalculations(input))

}

func sumCalculations(operations []string) int64 {
	var sum int64
	for _, calc := range operations {
		l := lexer.New(calc)
		p := parser.New(l)
		expression := p.ParseLine()
		result := evaluator.Eval(expression)
		sum += result
	}
	return sum
}

func parseInput(href string) []string {
	data, _ := ioutil.ReadFile(href)
	calculations := strings.Split(string(data), "\n")
	return calculations
}
