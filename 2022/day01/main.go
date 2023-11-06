package main

import (
	_ "embed"
	"fmt"
	"github.com/samber/lo"
	"log"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var rawInput string

func main() {
	fmt.Println("Day 01")

	fmt.Printf("Part 1: %d \r\n", part1(rawInput))
	fmt.Printf("Part 2: %d", part2(rawInput))
}

func part1(input string) int {
	elves := parseInput(input)

	return lo.Max(lo.Reduce(elves, func(agg []int, item []int, _ int) []int {
		return append(agg, lo.Sum(item))
	}, []int{}))
}

func part2(input string) int {
	elves := parseInput(input)
	sums := lo.Reduce(elves, func(agg []int, item []int, _ int) []int {
		return append(agg, lo.Sum(item))
	}, []int{})
	slices.Sort(sums)

	return lo.Sum(sums[len(sums)-3:])

}

func parseInput(input string) [][]int {
	result := [][]int{}

	input = strings.TrimRight(input, "\n")

	// First, split into groups
	for _, group := range strings.Split(input, "\n\n") {
		var intGroup []int
		for _, line := range strings.Split(group, "\n") {
			intLine, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("error converting `%s` value to int", line)
			}

			intGroup = append(intGroup, intLine)
		}

		result = append(result, intGroup)
	}

	return result
}
