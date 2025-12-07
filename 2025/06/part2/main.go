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
	contents = strings.Trim(contents, "\r\n")

	lines := strings.Split(contents, "\n")

	return lines
}

func separateOperations(line string) []string {
	line = strings.Trim(line, "\n\r")
	re := regexp.MustCompile(`[+*] *`)
	result := re.FindAllString(line, -1)

	for i, res := range result[:len(result)-1] {
		result[i] = res[:len(res)-1]
	}

	return result
}

func getValuesFromLine(line string, operations []string) []string {
	trimmedLine := strings.Trim(line, "\n\r")

	results := []string{}
	lastIdx := 0

  for _, operation := range operations {
		results = append(results, trimmedLine[lastIdx:lastIdx+len(operation)])
		lastIdx += len(operation)+1
	}
	
  return results
}

func getValues(lines []string) (valuesCorrected [][]int, operations []string) {
	lastIdx := len(lines) - 1
	operations = separateOperations(lines[lastIdx])
	values := [][]string{}

	for _, line := range lines[:lastIdx] {
		values = append(values, getValuesFromLine(line, operations))
	}

	valuesCorrected = make([][]int, len(operations))

	for i := range operations {
		valuesCorrected[i] = make([]int, len(values))
		strVals := make([][]byte, len(values))
		for j := range values {
			for k := range values[j][i] {
				strVals[k] = append(strVals[k], values[j][i][k])
			}
		}
		for j, str := range strVals {
			trimmedString := strings.TrimSpace(string(str))

			if len(trimmedString) == 0 {
				continue
			}

			num, err := strconv.Atoi(trimmedString)

			if err != nil {
				panic(err)
			}

			valuesCorrected[i][j] = num
		}
	}

	for i := range operations {
		operations[i] = strings.TrimSpace(operations[i])
	}

	return
}

func performOperation(valuesArray []int, operation string) int {

	if operation == "*" {
		result := 1

		for _, val := range valuesArray {
			if val == 0 {
				continue
			}

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

func handleOperations(valuesArray []int, operation string) int {
	return performOperation(valuesArray, operation)
}

func main() {
	lines := readInput("input.txt")
	values, operations := getValues(lines)
	results := []int{}

	for i := range values {
		results = append(results, handleOperations(values[i], operations[i]))
	}

	fmt.Println("The final result is", performOperation(results, "+"))
}
