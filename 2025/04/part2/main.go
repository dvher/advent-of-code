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

func updateMatrix(row, col int, matrix []string) {
	aux := []byte(matrix[row])
	aux[col] = byte('x')
	matrix[row] = string(aux)
}

func countReachableRolls(matrix []string) (int, []string) {
	reachableRolls := 0
	newMatrix := []string{}
	newMatrix = append(newMatrix, matrix...)

	for i := range matrix {
		for j := range matrix[i] {
			lengths := [2]int{len(matrix), len(matrix[i])}
			if matrix[i][j] == byte('@') && isReachable(i, j, matrix, lengths) {
				reachableRolls++
				updateMatrix(i, j, newMatrix)
			}
		}
	}

	return reachableRolls, newMatrix
}

func main() {
	lines := readInput("input.txt")
	reachableRolls := 0
	accumulatedRolls := 0

	// Go has no do-while
	for execOnce := true; execOnce; execOnce = (reachableRolls != 0) {
		reachableRolls, lines = countReachableRolls(lines)
		accumulatedRolls += reachableRolls
	}

	fmt.Printf("There are %d reachable rolls\n", accumulatedRolls)
}
