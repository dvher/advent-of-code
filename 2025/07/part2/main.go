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

func findTachyon(line []byte) []int {
	tachyonLoc := make([]int, len(line))
	for i := range line {
		if line[i] == byte('S') {
			tachyonLoc[i]++
		}
	}
	return tachyonLoc
}

func splitTachyons(tachyons []int, divisions []byte) []int {
	newTachyons := []int{}
	newTachyons = append(newTachyons, tachyons...)

	for i, timelines := range tachyons {

		if timelines > 0 && divisions[i] == byte('^') {
			if i == 0 {
				
				newTachyons[i+1] += newTachyons[i]
				newTachyons[i] = 0

			} else if i == len(tachyons) - 1 {

				newTachyons[i-1] += newTachyons[i]
				newTachyons[i] = 0

			} else {

				newTachyons[i+1] += newTachyons[i]
				newTachyons[i-1] += newTachyons[i]
				newTachyons[i] = 0

			}
		}
	}

	return newTachyons
}

func sum(arr []int) int {
	res := 0
	for _, val := range arr {
		res += val
	}
	return res
}

func countSplits(matrix [][]byte) int {
	if len(matrix) < 1 {
		return 0
	}

	tachyonLoc := findTachyon(matrix[0])

	for _, row := range matrix[1:] {
		tachyonLoc = splitTachyons(tachyonLoc, row)
	}

	return sum(tachyonLoc)
}

func main() {
	lines := readInput("input.txt")
	divisions := countSplits(lines)
	fmt.Println("Tachyon splits", divisions, "times")
}
