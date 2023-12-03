package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getLines(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func testColor(input string, r *regexp.Regexp, max int) bool {
	num, _ := regexp.Compile("\\d+")
	for _, match := range r.FindAllString(input, -1) {
		for _, str := range num.FindAllString(match, -1) {
			n, _ := strconv.Atoi(str)
			if n > max {
				return false
			}
		}
	}
	return true
}

func parseLine(line string) int {
	x := strings.Split(line, ":")
	index, _ := strconv.Atoi(strings.Split(x[0], " ")[1])

	reds, _ := regexp.Compile("\\d+ red")
	greens, _ := regexp.Compile("\\d+ green")
	blues, _ := regexp.Compile("\\d+ blue")

	if testColor(x[1], reds, 12) &&
		testColor(x[1], greens, 13) &&
		testColor(x[1], blues, 14) {
		return index
	}

	return 0
}

func main() {
	//cwd, _ := os.Getwd()
	input, err := os.ReadFile("input/input-full.txt")

	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, line := range getLines(input) {

		sum += parseLine(line)
	}

	fmt.Println(sum, "sum")
}
