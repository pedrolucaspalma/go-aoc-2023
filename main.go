package main

import (
	"flag"
	"fmt"
)

func main() {
	day := getDay()
	fmt.Printf("Displaying solution for day %d \n", day)

}

func getDay() int {
	var dayPtr = flag.Int("d", 1, "n sei oq e isso")
	var day = *dayPtr
	if day < 1 {
		day = 1
	}
	return day
}
