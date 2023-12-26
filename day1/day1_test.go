package day1

import (
	"fmt"
	"testing"
)

func TestDayOne(t *testing.T) {
	val := Solution("1", true)
	if val != 142 {
		formated := fmt.Sprint(val)
		t.Error("Expected 42, got " + formated)
	}
}

func TestDayOneRealInput(t *testing.T) {
	val := Solution("2", true)
	if val != 57346 {
		formated := fmt.Sprint(val)
		t.Error("Expected 57346, got " + formated)
	}
}
