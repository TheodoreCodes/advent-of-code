package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 04")
	fmt.Printf("Part 1: %d \r\n", part1(input))
	fmt.Printf("Part 2: %d", part2(input))
}

func part1(input string) int {
	lines := parseInput(input)
	counter := 0

	for _, line := range lines {
		pair := getPair(line)

		if (pair[0][0] >= pair[1][0] && pair[0][1] <= pair[1][1]) ||
			(pair[1][0] >= pair[0][0] && pair[1][1] <= pair[0][1]) {
			counter++
		}
	}

	return counter
}

func part2(input string) int {
	lines := parseInput(input)
	counter := 0

	for _, line := range lines {
		pair := getPair(line)

		for i := 0; i <= 1; i++ {
			for j := 0; j <= 1; j++ {
				if isIntBetweenBounds(pair[i][j], pair[(i+1)%2][0], pair[(i+1)%2][1]) {
					counter++
					i = 2 // break for outer loop
					break
				}
			}
		}
	}

	return counter
}

func parseInput(rawInput string) []string {
	input = strings.TrimRight(rawInput, "\n")

	var result []string

	for _, line := range strings.Split(input, "\n") {
		result = append(result, line)
	}

	return result
}

func getPair(s string) [2][2]int {
	var pairs [2][2]int
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

	match := re.FindAllStringSubmatch(s, 1)[0]
	pairs[0][0], _ = strconv.Atoi(match[1])
	pairs[0][1], _ = strconv.Atoi(match[2])

	pairs[1][0], _ = strconv.Atoi(match[3])
	pairs[1][1], _ = strconv.Atoi(match[4])

	return pairs
}

func isIntBetweenBounds(num, leftBound, rightBound int) bool {
	return num >= leftBound && num <= rightBound
}
