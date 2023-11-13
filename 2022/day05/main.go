package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day 05")
	fmt.Printf("Part 1: %s \r\n", part1(input))
	fmt.Printf("Part 2: %s", part2(input))
}

func part1(input string) string {
	stacks, moves := parseInput(input)

	for _, m := range moves {
		for i := 0; i < m.Qty; i++ {
			stacks[m.To].Push(stacks[m.From].Pop())
		}
	}

	var returnStr string
	for _, s := range stacks {
		returnStr += s.Pop()
	}

	return returnStr
}

func part2(input string) string {
	stacks, moves := parseInput(input)

	for _, m := range moves {
		stacks[m.To].PushSlice(stacks[m.From].PopSlice(m.Qty))
	}

	var returnStr string
	for _, s := range stacks {
		returnStr += s.Pop()
	}

	return returnStr
}

func parseInput(input string) ([]stack[string], []move) {
	input = strings.TrimRight(input, "\n")
	input = strings.Replace(input, "\r\n\r\n", "\n\n", 1)
	parts := strings.Split(input, "\n\n")

	// Build the stacks
	var stacks []stack[string]
	crateRe := regexp.MustCompile(`\[([A-Z])]|^ {3}| {3}$| ( {3})`)
	for _, line := range strings.Split(parts[0], "\n") {
		matches := crateRe.FindAllStringSubmatch(line, -1)
		if len(stacks) == 0 {
			stacks = make([]stack[string], len(matches))
		}

		for i, match := range matches {
			if match[1] == "" {
				continue
			}

			stacks[i].PushBottom(match[1])
		}

		if line[1] == '1' {
			break
		}
	}

	// Build the moves
	var moves []move
	moveRe := regexp.MustCompile("move (\\d+) from (\\d) to (\\d)")
	for _, line := range strings.Split(parts[1], "\n") {
		matches := moveRe.FindAllStringSubmatch(line, 3)

		for _, match := range matches {
			var m move
			m.SetQty(match[1]).
				SetFrom(match[2]).
				SetTo(match[3])

			moves = append(moves, m)
		}
	}

	return stacks, moves
}
