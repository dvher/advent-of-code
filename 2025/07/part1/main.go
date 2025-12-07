package main

import (
	"bytes"
	"fmt"
	"os"
)

func readInput(filename string) [][]byte {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	file = bytes.TrimSpace(file)
	
	return bytes.Split(file, []byte("\n"))
}

func findTachyon(line []byte) []bool {
	tachyonLoc := make([]bool, len(line))
	for i := range line {
		if line[i] == byte('S') {
			tachyonLoc[i] = true
		}
	}
	return tachyonLoc
}

func splitTachyons(tachyons []bool, divisions []byte, divisionsCounter *int) []bool {
	newTachyons := make([]bool, len(tachyons))

	for i, hasTachyon := range tachyons {
		if !hasTachyon {
			continue
		}

		if divisions[i] == byte('^') {
			*divisionsCounter++
			if i == 0 {
				newTachyons[i+1] = true
			} else if i == len(tachyons) - 1 {
				newTachyons[i-1] = true
			} else {
				newTachyons[i+1] = true
				newTachyons[i-1] = true
			}
		} else {
			newTachyons[i] = hasTachyon
		}
	}

	return newTachyons
}

func countSplits(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}

	divisions := 0

	tachyonLoc := findTachyon(matrix[0])

	for _, row := range matrix[1:] {
		tachyonLoc = splitTachyons(tachyonLoc, row, &divisions)
	}

	return divisions
}

func main() {
	lines := readInput("input.txt")
	divisions := countSplits(lines)
	fmt.Println("Tachyon splits", divisions, "times")
}
