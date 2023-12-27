package main

import (
	"aoc-2023/day1"
	"aoc-2023/day2"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	day := getDay()
	fmt.Printf("Displaying solution for day %d \n", day)

	switch {
	case day == 1:
		// Error handling on this function could be better
		day1.Solution("2", false)
	case day == 2:

		val, err := day2.Solution("2", false)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Day2 solution \n")
		fmt.Printf("Sum: %d\n", val)
	default:
		log.Fatal("Invalid date received")
	}
}

func getDay() int {
	var dayString = os.Args[1]

	day, err := strconv.ParseInt(dayString, 10, 64)
	if err != nil {
		log.Fatal("Invalid date received")
	}

	fmt.Printf("day variable: %d \n", day)
	if day < 1 {
		day = 1
	}
	return int(day)
}
