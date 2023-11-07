package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

const ALPHA = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println("Day 03")
	fmt.Printf("Part 1: %d \r\n", part1(input))
	fmt.Printf("Part 2: %d", part2(input))
}

func part1(input string) int {
	rucksacks := parseInput(input)
	var priority int

	for _, rucksack := range rucksacks {
		l := len(rucksack)
		seen := make(map[rune]bool, 52)

		for _, c := range rucksack[:l/2] {
			if _, ok := seen[c]; !ok && strings.ContainsRune(rucksack[l/2:], c) {
				seen[c] = true
				priority += strings.IndexRune(ALPHA, c) + 1
			}
		}
	}

	return priority
}

func part2(input string) int {
	rucksacks := parseInput(input)
	var priority int

	for i := 0; i < len(rucksacks); i += 3 {
		map1 := stringToMap(rucksacks[i])
		map2 := stringToMap(rucksacks[i+1])
		map3 := stringToMap(rucksacks[i+2])

		for c, _ := range map1 {
			if map2[c] && map3[c] {
				priority += strings.IndexRune(ALPHA, c) + 1
			}
		}
	}

	return priority
}

func parseInput(rawInput string) []string {
	input = strings.TrimRight(rawInput, "\n")

	var result []string

	for _, line := range strings.Split(input, "\n") {
		result = append(result, line)
	}

	return result
}

func stringToMap(s string) map[rune]bool {
	m := map[rune]bool{}
	for _, c := range s {
		m[c] = true
	}

	return m
}
