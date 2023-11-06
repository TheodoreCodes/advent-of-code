package main

import (
	"reflect"
	"testing"
)

const testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test_parseInput(t *testing.T) {
	expected := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10_000},
	}

	t.Run("day 01 parseInput", func(t *testing.T) {
		if actual := parseInput(testInput); !reflect.DeepEqual(expected, actual) {
			t.Errorf("day 01 parseInput: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part1(t *testing.T) {
	expected := 24000

	t.Run("day 01 part 1", func(t *testing.T) {
		if actual := part1(testInput); actual != expected {
			t.Errorf("day 01 part 1: expected: %d, actual %d", expected, actual)
		}
	})
}

func Test_part2(t *testing.T) {

}
