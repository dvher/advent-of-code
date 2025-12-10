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

func getCoords(lines []string) [][]int {
	coords := [][]int{}
	for _, line := range lines {
		coordStr := strings.Split(strings.TrimSpace(line), ",")
		coord := []int{}

		for _, c := range coordStr {
			coordInt, err := strconv.Atoi(c)

			if err != nil {
				panic(err)
			}

			coord = append(coord, coordInt)
		}

		coords = append(coords, coord)
	}

	return coords
}

func calculateArea(coordA, coordB []int) int {
	if len(coordA) != len(coordB) {
		panic("Coordinates have different dimensions")
	}

	dimensions := []int{}

	for i := range coordA {
		dimensions = append(dimensions, coordB[i] - coordA[i] + 1)		
	}

	res := 1

	for _, plane := range dimensions {
		res *= plane
	}


	return res
}

func computeAllAreas(coords [][]int) []int {
	areas := []int{}

	for _, coordA := range coords {
		for _, coordB := range coords {
			areas = append(areas, calculateArea(coordA, coordB))
		}
	}

	return areas
}

func getMaxArea(areas []int) int {
	maxArea := 0

	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	lines := readInput("input.txt")

	coords := getCoords(lines)
	areas := computeAllAreas(coords)

	fmt.Println("The maximum area is", getMaxArea(areas))
}
