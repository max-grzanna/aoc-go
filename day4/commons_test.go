package main

import (
	"log"
	"os"
	"testing"
)

func TestCommons(t *testing.T) {
	cwd, _ := os.Getwd()
	input, err := os.ReadFile(cwd + "/input/input-small.txt")

	if err != nil {
		log.Fatal(err)
	}

	points := 0
	result := 0
	for _, line := range getLines(input) {
		parsed := parseLine(line)
		result = len(getCommons(parsed[0], parsed[1]))
		points = points + calcPoints(result)
	}

	if points != 13 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", points, "13")
	}
}
