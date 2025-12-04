package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	ROW = iota
	COL
)

func readInput(filename string) []string {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	contents := string(file)
	contents = strings.Trim(contents, "\r\n")

	lines := strings.Split(contents, "\n")

	return lines
}

func checkCoordOccupied(row, col int, matrix []string, lengths [2]int) bool {
	if row < 0 || col < 0 {
		return false
	}

	if row >= lengths[ROW] || col >= lengths[COL] {
		return false
	}

	if matrix[row][col] == byte('@') {
		return true
	}

	return false
}

func isReachable(row, col int, matrix []string, lengths [2]int) bool {
	coords := []int{-1, 0, 1}

	adjacentRolls := 0

	for _, intRow := range coords {
		for _, intCol := range coords {
			if intRow == 0 && intCol == 0 {
				continue
			}

			if checkCoordOccupied(
				row + intRow,
				col + intCol,
				matrix,
				lengths,
			) {
				adjacentRolls++
			}
		}
	}

	return adjacentRolls < 4
}

func countReachableRolls(matrix []string) int {
	reachableRolls := 0
	for i := range matrix {
		for j := range matrix[i] {
			lengths := [2]int{len(matrix), len(matrix[i])}
			if matrix[i][j] == byte('@') && isReachable(i, j, matrix, lengths) {
				reachableRolls++
			}
		}
	}

	return reachableRolls
}

func main() {
	lines := readInput("input.txt")

	reachableRolls := countReachableRolls(lines)

	fmt.Printf("There are %d reachable rolls\n", reachableRolls)
}
