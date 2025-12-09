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

func intersection(arr1, arr2 []int) []int {
	intersect := []int{}

	// Go doesn't have sets
	set := make(map[int]bool)

	for _, value := range arr1 {
		set[value] = true
	}

	for _, value := range arr2 {
		if _, found := set[value]; found {
			intersect = append(intersect, value)
		}
	}

	return intersect
}

func merge(arr1, arr2 []int) []int {
	// Go doesn't have sets
	set := make(map[int]bool)
	finalSet := []int{}

	for _, val := range arr1 {
		set[val] = true
	}

	for _, val := range arr2 {
		set[val] = true
	}

	for key := range set {
		finalSet = append(finalSet, key)
	}

	return finalSet
}

func mergeCircuits(circuits [][]int) [][]int {

	if len(circuits) < 2 {
		return circuits
	}

	merged := false

	for execOnce := true; execOnce; execOnce = merged {
		merged = false
    madeChanges := make(map[int]bool)
  
    for i := range circuits {
      if madeChanges[i] {
      	continue
      }
      for j := i+1; j < len(circuits); j++ {
      	if madeChanges[j] {
      		continue
      	}
      	c1 := circuits[i]
      	c2 := circuits[j]
      	intersect := intersection(c1, c2)
      	if len(intersect) > 0 {
      		madeChanges[j] = true
      		merged = true
      		mergedCircuits := merge(c1, c2)
      		circuits[i] = mergedCircuits
      	}
      }
    }

    newCircuits := [][]int{}

    for i, circ := range circuits {
    	if !madeChanges[i] {
    		newCircuits = append(newCircuits, circ)
    	}
    }

    circuits = newCircuits
  }

  return circuits
}

func cmpDistances(a, b []int) int {
	return a[2] - b[2]
}

func findNShortestPaths(disMatrix [][]int) [][]int {
	circuits := [][]int{}

	slices.SortFunc(disMatrix, cmpDistances)

	for i := range PAIRS_TO_CONNECT {
		circuits = append(circuits, disMatrix[i][0:2])
	}

	return circuits
}

func cmpLengths(a, b []int) int {
	return len(b) - len(a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func multiplyNLargest(circuits [][]int) int {
	res := 1
	for i := range min(LARGEST_TO_MULT, len(circuits)) {
		res *= len(circuits[i])
	}

	return res
}

func main() {
	lines := readInput("input.txt")

	coords := getCoordMatrix(lines)
	disMatrix := getDistanceMatrix(coords)
	circuits := findNShortestPaths(disMatrix)
	circuits = mergeCircuits(circuits)

	slices.SortFunc(circuits, cmpLengths)

	fmt.Println("The result is", multiplyNLargest(circuits))
}
