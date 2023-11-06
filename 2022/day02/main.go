package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 02")
	fmt.Printf("Part 1: %d \r\n", part1(input))
	fmt.Printf("Part 2: %d", part2(input))
}

func part1(input string) int {
	outcomeScore := map[string]int{
		"AX": 3,
		"BX": 0,
		"CX": 6,

		"AY": 6,
		"BY": 3,
		"CY": 0,

		"AZ": 0,
		"BZ": 6,
		"CZ": 3,
	}

	shapeScore := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	result := 0

	for _, round := range parseInput(input) {
		result += outcomeScore[round]
		result += shapeScore[round[1:]]
	}

	return result
}

func part2(input string) int {
	outcomeScore := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	shapeScore := map[string]int{
		"AX": 3,
		"AY": 1,
		"AZ": 2,

		"BX": 1,
		"BY": 2,
		"BZ": 3,

		"CX": 2,
		"CY": 3,
		"CZ": 1,
	}

	result := 0

	for _, round := range parseInput(input) {
		result += outcomeScore[round[1:]]
		result += shapeScore[round]
	}

	return result

}

func parseInput(rawInput string) []string {
	input = strings.TrimRight(rawInput, "\n")
	var result []string

	for _, line := range strings.Split(input, "\n") {
		result = append(result, strings.Replace(line, " ", "", 1))
	}

	return result
}
