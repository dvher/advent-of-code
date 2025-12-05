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

func getInputs(lines []string) ([]string) {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[:i]
		}
	}
	panic("Invalid input")
}

func getRange(ranges []string) [][]int {
	rangeMatrix := [][]int{}

	for _, r := range ranges {
		indexes := strings.Split(r, "-")

		intIndexes := make([]int, len(indexes))

		for i := range intIndexes {
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

func rangesIntersect(range1, range2 []int) bool {
	return !(range1[0] > range2[1] || range1[1] < range2[0])
}

func consolidateRanges(range1, range2 []int) []int {
	return []int{min(range1[0], range2[0]), max(range1[1], range2[1])}
}

func combineRanges(rangeMatrix [][]int) [][]int {
	changeMade := false
	lastMatrix := rangeMatrix

	for execOnce := true; execOnce; execOnce = changeMade {
		changeMade = false

		OUTER_MATRIX_LOOP:
		for i := 0; i < len(lastMatrix) - 1; i++ {
			for j := i + 1; j < len(lastMatrix); j++ {
				range1 := lastMatrix[i]
				range2 := lastMatrix[j]
				if rangesIntersect(range1, range2) {
					changeMade = true
					newRange := consolidateRanges(range1, range2)
					toAppend := append(lastMatrix[i+1:j], lastMatrix[j+1:]...)
					lastMatrix = append(lastMatrix[:i], newRange)
					lastMatrix = append(lastMatrix, toAppend...)
					break OUTER_MATRIX_LOOP
				}
			}
		}

	}

	return lastMatrix
}

func countFresh(ranges [][]int) int {
	counter := 0

	for _, r := range ranges {
		counter += (r[1] - r[0] + 1)
	}

	return counter
}

func main() {
	lines := readInput("input.txt")
	ranges := getInputs(lines)
	rangeMatrix := getRange(ranges)

	rangeMatrix = combineRanges(rangeMatrix)
	freshIndexes := countFresh(rangeMatrix)

	fmt.Println("There are", freshIndexes, "fresh ingredients")
}
