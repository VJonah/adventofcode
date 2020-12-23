package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	circle := parseInput("test.txt")
	fmt.Println(circle.Cups)
	shuffled := shuffleCups(1, circle)
	//fmt.Println(shuffled.getTwoCloseToCup(1))
	fmt.Println(shuffled)
}

func shuffleCups(n int, circle *Circle) *Circle {
	seenCombinations := map[string]int{circle.getCupOrder(): 0}
	for i := 0; i < n; i++ {
		currentCup := circle.peekFirstCup()
		nextThreeCups := circle.getThreeAdjacentCups()
		curLabel := currentCup.Label
		fmt.Printf("curCp: %d, nxt: %d, curLbl: %d \n", currentCup, nextThreeCups, curLabel)
		destinationIdx := -1
		for true {
			if curLabel == 1 {
				curLabel = 1000000
			} else {
				curLabel--
			}
			if circle.getCupIndex(curLabel) != -1 {
				destinationIdx = circle.getCupIndex(curLabel)
				break
			}
		}
		fmt.Printf("dest: %d, cups: %d \n", destinationIdx, circle.Cups)
		circle.addCupsClockwise(destinationIdx, nextThreeCups)
		circle.nextCup()
		order := circle.getCupOrder()
		if seenCombinations[order] == 0 {
			seenCombinations[order] = i
		} else {
			fmt.Println(circle.Cups)
		}
	}
	fmt.Println(len(seenCombinations))
	return circle
}

type Cup struct {
	Label int
}

type Circle struct {
	Cups []Cup
}

func (c *Circle) peekFirstCup() Cup {
	return c.Cups[0]
}

func (c *Circle) getThreeAdjacentCups() []Cup {
	adjacent := append([]Cup{}, c.Cups[1:4]...)
	c.Cups = append(c.Cups[:1], c.Cups[4:]...)
	return adjacent
}

func (c *Circle) getTwoCloseToCup(n int) []Cup {
	cupIdx := c.getCupIndex(n)
	cups := append([]Cup{}, c.Cups[cupIdx+1:cupIdx+3]...)
	return cups
}
func (c *Circle) getCupIndex(label int) int {
	for i, cup := range c.Cups {
		if cup.Label == label {
			return i
		}
	}
	return -1
}

func (c *Circle) maximumLabel() int {
	max := -1
	for _, cup := range c.Cups {
		if cup.Label > max {
			max = cup.Label
		}
	}
	return max
}

func (c *Circle) minimumLabel() int {
	min := -1
	for _, cup := range c.Cups {
		if cup.Label < min {
			min = cup.Label
		}
	}
	return min
}

func (c *Circle) nextCup() {
	c.Cups = append(c.Cups[1:], c.Cups[0])
}

func (c *Circle) addCupsClockwise(destination int, cups []Cup) {
	newCups := append([]Cup{}, c.Cups[:destination+1]...)
	newCups = append(newCups, cups...)
	c.Cups = append(newCups, c.Cups[destination+1:]...)
}

func (c *Circle) getCupOrder() string {
	oneIdx := c.getCupIndex(1)
	ordered := append([]Cup{}, c.Cups[oneIdx+1:]...)
	ordered = append(ordered, c.Cups[:oneIdx]...)
	order := ""
	for _, cup := range ordered {
		val := strconv.Itoa(cup.Label)
		order += val
	}
	return order
}

func parseInput(href string) *Circle {
	data, _ := ioutil.ReadFile(href)
	cups := []Cup{}
	for _, digit := range string(data) {
		val, _ := strconv.Atoi(string(digit))
		cup := Cup{Label: val}
		cups = append(cups, cup)
	}
	for i := 10; i <= 1000000; i++ {
		cups = append(cups, Cup{Label: i})
	}
	return &Circle{Cups: cups}
}
