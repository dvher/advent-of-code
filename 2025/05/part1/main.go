package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	contents := string(file)
	contents = strings.TrimSpace(contents)

	return strings.Split(contents, "\n")
}

func getInputs(lines []string) ([]string, []string) {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[:i], lines[i+1:]
		}
	}
	panic("Invalid input")
}

func getRange(ranges []string) [][]int {
	rangeMatrix := [][]int{}

	for _, r := range ranges {
		indexes := strings.Split(r, "-")

		intIndexes := make([]int, len(indexes))

		for i, _ := range intIndexes {
			idx, err := strconv.Atoi(indexes[i])

			if err != nil {
				panic(err)
			}

			intIndexes[i] = idx
		}

		rangeMatrix = append(rangeMatrix, intIndexes)
	}

	return rangeMatrix
}

func countFresh(rangeMatrix [][]int, ingredients []string) int {
	freshIngredients := 0

	for _, ingredient := range ingredients {
		idx, err := strconv.Atoi(ingredient)

		if err != nil {
			panic(err)
		}

		for _, ranges := range rangeMatrix {
			if idx >= ranges[0] && idx <= ranges[1] {
				freshIngredients++
				break
			}
		}
		
	}

	return freshIngredients
}

func main() {
	lines := readInput("input.txt")
	ranges, ingredients := getInputs(lines)
	rangeMatrix := getRange(ranges)

	freshIngredients := countFresh(rangeMatrix, ingredients)

	fmt.Println("There are", freshIngredients, "fresh ingredients")
}
