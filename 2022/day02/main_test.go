package main

import (
	"reflect"
	"testing"
)

const testInput = `A Y
B X
C Z`

func Test_parseInput(t *testing.T) {
	expected := []string{"AY", "BX", "CZ"}

	t.Run("day 02 parseInput", func(t *testing.T) {
		if actual := parseInput(testInput); !reflect.DeepEqual(expected, actual) {
			t.Errorf("day 02 parseInput: expected: %s, actual %s", expected, actual)
		}
	})
}

func Test_part1(t *testing.T) {
	expected := 15

	t.Run("day 02 part 1", func(t *testing.T) {
		if actual := part1(testInput); actual != expected {
			t.Errorf("day 02 part 1: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {
	expected := 12

	t.Run("day 02 part 2", func(t *testing.T) {
		if actual := part2(testInput); actual != expected {
			t.Errorf("day 02 part 2: expected: %d, actual %d", expected, actual)
		}
	})
}
