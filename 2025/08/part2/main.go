package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const PAIRS_TO_CONNECT = 1000
const LARGEST_TO_MULT = 3

func readInput(filename string) []string {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	contents := string(file)
	contents = strings.TrimSpace(contents)

	return strings.Split(contents, "\n")
}

func getCoordMatrix(lines []string) [][]int {
	matrix := [][]int{}

	for _, line := range lines {
		values := strings.Split(line, ",")
		row := make([]int, len(values))

		for i := range values {
			numericValue, err := strconv.Atoi(values[i])

			if err != nil {
				panic(err)
			}

			row[i] = numericValue
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func calculateDistance(coordA, coordB []int) int {
	if len(coordA) != len(coordB) {
		panic("Coordinates length doesn't match")
	}

	if slices.Equal(coordA, coordB) {
		return math.MaxInt
	}

	var result int = 0

	for i := range coordA {
		result += (coordA[i] - coordB[i]) * (coordA[i] - coordB[i])
	}

	return result
}

func getDistanceMatrix(coords [][]int) [][]int {
	matrix := [][]int{}
	calculated := make(map[string]bool)

	for i, coordA := range coords {
		for j, coordB := range coords {
			distance := calculateDistance(coordA, coordB)
			value := []int{i, j, distance}
			key := fmt.Sprintf("%d-%d", j, i)
			key2 := fmt.Sprintf("%d-%d", i, j)
			if _, exists := calculated[key]; !exists {
				matrix = append(matrix, value)
				calculated[key] = true
				calculated[key2] = true
			}
		}
	}
	
	return matrix
}

func cmpDistances(a, b []int) int {
	return a[2] - b[2]
}

func computeAllCircuits(lines []string) [][]int {
	circuits := [][]int{}

	for i := range lines {
		circuits = append(circuits, []int{i})
	}

	return circuits
}

func findInCircuit(v int, circuits [][]int) int {
	for i := range circuits {
		if idx := slices.Index(circuits[i], v); idx != -1 {
			return i
		}
	}

	return -1
}

func delete(matrix [][]int, idx int) [][]int {
	newMatrix := [][]int{}

	newMatrix = append(newMatrix, matrix[:idx]...)
	newMatrix = append(newMatrix, matrix[idx+1:]...)

	return newMatrix
}

func main() {
	lines := readInput("input.txt")

	coords := getCoordMatrix(lines)
	circuits := computeAllCircuits(lines)
	disMatrix := getDistanceMatrix(coords)
	
	slices.SortFunc(disMatrix, cmpDistances)

	for i := range disMatrix {
		idx1 := findInCircuit(disMatrix[i][0], circuits)
		idx2 := findInCircuit(disMatrix[i][1], circuits)

		if len(circuits[idx1]) >= 1 && len(circuits[idx2]) >= 1 && idx1 != idx2 {
			circuits[idx1] = append(circuits[idx1], circuits[idx2]...)
			circuits = delete(circuits, idx2)
		}

		if len(circuits) == 1 {
			coordA := coords[disMatrix[i][0]]
			coordB := coords[disMatrix[i][1]]
			fmt.Println("The result is", coordA[0] * coordB[0])
			break
		}
	}
}
