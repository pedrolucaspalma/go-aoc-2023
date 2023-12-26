package day1

import (
	"fmt"
	"os"
	"strconv"
)

func Solution(inputToSelect string, isTest bool) int {
	var txt = getInput(inputToSelect, isTest)

	var sum int
	var curr string
	for _, c := range txt {
		if string(c) == "\n" {
			firstChar := string(curr[0])
			secondChar := string(curr[len(curr)-1])
			curr = firstChar + secondChar
			rowNumber, err := strconv.ParseInt(curr, 10, 64)
			if err != nil {
				panic(err)
			}
			sum += int(rowNumber)
			curr = ""

			continue
		} else if _, err := strconv.ParseInt(string(c), 10, 64); err == nil {
			curr += string(c)
		}

	}

	fmt.Printf("Sum: %d \n", sum)
	return sum
}

func getInput(input string, isTest bool) string {
	var basePath string

	if isTest == true {
		basePath = "./"
	} else {
		basePath = "./day1/"
	}

	var filePath = basePath + "input-" + input + ".md"

	text, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return string(text)
}
