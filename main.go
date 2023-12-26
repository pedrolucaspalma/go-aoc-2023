package main

import (
	"aoc-2023/day1"
	"flag"
	"fmt"
)

func main() {
	day := getDay()
	fmt.Printf("Displaying solution for day %d \n", day)

	switch {
	case day == 1:
		day1.Solution("2", false)
	default:
		panic("Invalid date received")
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
