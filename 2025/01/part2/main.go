package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

const NUM_CLICKS = 100

func main() {
  file, err := os.ReadFile("./input.txt")
  if err != nil {
  	fmt.Println("Error al leer archivo ", err)
  	return
  }

  contents := string(file)
	lines := strings.Split(contents, "\n")

	curValue := 50
	password := 0

	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		dir := line[:1]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error al convertir numero ", err)
			return
		}

		if dir == "R" {
			password += (curValue + num) / NUM_CLICKS
			curValue = (curValue + num) % NUM_CLICKS
		} else {
			if curValue - num == 0 {
				password++
			}
			if curValue - num < 0 {
				if curValue != 0 {
					password++
				}
				password += (num - curValue) / NUM_CLICKS
			}
			curValue = (curValue - num) % NUM_CLICKS
		}

	}

	fmt.Println("The password is ", password)
}

