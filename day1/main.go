package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func calcSum(path string) (int, error) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	scanner := bufio.NewScanner(f)

	total := 0
	var relevant []string
	for scanner.Scan() {
		var sub []string
		for _, r := range scanner.Text() {
			if unicode.IsDigit(r) {
				sub = append(sub, string(r))
			}
		}

		relevant = []string{sub[0], sub[len(sub)-1]}
		joinedValues := strings.Join(relevant[:], "")

		val, err := strconv.Atoi(joinedValues)
		if err != nil {
			panic(err)
		}
		total += val

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total, nil
}

func main() {
	path := "resources/day_1/input-full.txt"
	result, err := calcSum(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
