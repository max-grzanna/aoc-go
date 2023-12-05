package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func parseLine(input string) [][]int {
	//separate the input in a slice of two slices
	parts := strings.Split(input, "|")
	l := extractDigitsFromPart(parts[0])
	r := extractDigitsFromPart(parts[1])
	leftRightNested := [][]int{
		l,
		r,
	}

	return leftRightNested
}

func extractDigitsFromPart(part string) []int {
	var digits []int

	// Split the part by whitespace
	substrings := strings.Fields(part)

	// Extract digits and convert them into integers
	for _, substr := range substrings {
		digit, err := strconv.Atoi(substr)
		if err == nil {
			digits = append(digits, digit)
		}
	}

	return digits
}

func getCommons(arr1, arr2 []int) []int {
	var out []int
	bucket := map[int]bool{}
	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}
	return out
}

func calcPoints(numMatches int) int {
	if numMatches == 0 {
		return 0
	} else {
		points := 1
		for i := 1; i < numMatches; i++ {
			points *= 2
		}
		return points
	}
}

func main() {
	cwd, _ := os.Getwd()
	input, err := os.ReadFile(cwd + "/input/input-full.txt")

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
}
