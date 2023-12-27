package main

import (
	"aoc-2023/day1"
	"aoc-2023/day2"
	"flag"
	"fmt"
	"log"
)

func main() {
	day := getDay()
	fmt.Printf("Displaying solution for day %d \n", day)

	switch {
	case day == 1:
		// Error handling on this function could be better
		day1.Solution("2", false)
	case day == 2:

		val, err := day2.Solution("2", true)
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
	var dayPtr = flag.Int("d", 1, "n sei oq e isso")
	var day = *dayPtr
	if day < 1 {
		day = 1
	}
	return day
}
