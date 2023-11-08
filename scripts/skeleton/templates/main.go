package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day {{.Day}}")
	fmt.Printf("Part 1: %d \r\n", part1(input))
	fmt.Printf("Part 2: %d", part2(input))
}

func part1(input string) int {
	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(rawInput string) []string {
	input = strings.TrimRight(rawInput, "\n")

	var result []string

	for _, line := range strings.Split(input, "\n") {
		result = append(result, line)
	}

	return result
}
