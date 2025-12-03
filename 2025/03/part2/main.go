package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const BATTERIES_TO_TURN = 12

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

func getMax(index, limit int, str string) (string, int) {
	maxVal := byte('0')
	maxIndex := -1
	for i := index; i < limit; i++ {
		if str[i] > maxVal {
			maxVal = str[i]
			maxIndex = i
		}
	}

	return string(maxVal), maxIndex
}

func getMaxJoltage(bank string) int {
	bankLen := len(bank)
	lastIndex := -1
	lastJoltage := ""
	curJoltage := ""

	for i := range BATTERIES_TO_TURN {
		lastJoltage, lastIndex = getMax(lastIndex + 1, bankLen - BATTERIES_TO_TURN + i + 1, bank)
		curJoltage += lastJoltage
	}

	joltageValue, err := strconv.Atoi(curJoltage)

	if err != nil {
		panic(err)
	}

	return joltageValue
}

func main() {
	lines := readInput("input.txt")

	total := 0
	for _, line := range lines {
		total += getMaxJoltage(line)
	}

	fmt.Println("Total joltage:", total)

}
