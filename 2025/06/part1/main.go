package main

import (
	"fmt"
	"os"
	"regexp"
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

	lines := strings.Split(contents, "\n")

	return lines
}

func getValues(lines []string) (values [][]string) {
	re := regexp.MustCompile(" +")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = re.ReplaceAllString(line, " ")

		lineValues := strings.Split(line, " ")
		values = append(values, lineValues)
	}

	return
}

func getLengths(matrix [][]string) (int, int) {
	rows := len(matrix)
	cols := 0

	for i := 0; i < min(1, rows); i++ {
		cols = len(matrix[i])
	}

	return rows, cols
}

func getValuesArray(rowLen, colIdx int, matrix [][]string) []int {
	valuesArray := make([]int, rowLen - 1)

	for i := 0; i < rowLen - 1; i++ {
		val, err := strconv.Atoi(matrix[i][colIdx])

		if err != nil {
			panic(err)
		}

		valuesArray[i] = val
	}

	return valuesArray
}

func performOperation(valuesArray []int, operation string) int {

	if operation == "*" {
		result := 1

		for _, val := range valuesArray {
			result *= val
		}

		return result
	}

	result := 0

	for _, val := range valuesArray {
		result += val
	}

	return result
}

func handleOperations(rowLen, colIdx int, matrix [][]string) int {
	valuesArray := getValuesArray(rowLen, colIdx, matrix)

	operation := matrix[rowLen - 1][colIdx]

	return performOperation(valuesArray, operation)
}

func main() {
	lines := readInput("input.txt")
	values := getValues(lines)

	rowLen, colLen := getLengths(values)
	results := []int{}

	for i := range colLen {
		results = append(results, handleOperations(rowLen, i, values))
	}

	fmt.Println("The final result is", performOperation(results, "+"))
}
