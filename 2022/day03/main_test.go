package main

import (
	"reflect"
	"testing"
)

const testInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func Test_parseInput(t *testing.T) {
	expected := [][2]string{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
	}

	t.Run("day 03 parseInput", func(t *testing.T) {
		if actual := parseInput(testInput); !reflect.DeepEqual(expected, actual[0:3]) {
			t.Errorf("day 03 parseInput: expected: %s, actual %s", expected, actual)
		}
	})
}

func Test_part1(t *testing.T) {
	expected := 157

	t.Run("day 03 part 1", func(t *testing.T) {
		if actual := part1(testInput); actual != expected {
			t.Errorf("day 03 part 1: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {
	expected := 70

	t.Run("day 03 part 2", func(t *testing.T) {
		if actual := part2(testInput); actual != expected {
			t.Errorf("day 03 part 2: expected: %d, actual %d", expected, actual)
		}
	})
}
