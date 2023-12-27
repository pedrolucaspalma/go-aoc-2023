package day2

import (
	"fmt"
	"log"
	"testing"
)

func TestDayTwoInitialParam(t *testing.T) {
	val, err := Solution("1", true)
	if err != nil {
		log.Fatal(err)
	}

	if val != 8 {
		formated := fmt.Sprint(val)
		t.Error("Expected 8, got " + formated)
	}
}

func TestDayTwoQuestionParam(t *testing.T) {
	val, err := Solution("2", true)
	if err != nil {
		log.Fatal(err)
	}

	if val != 2685 {
		t.Error("Expected 2685, got " + fmt.Sprint(val))
	}
	fmt.Println(val)
}
