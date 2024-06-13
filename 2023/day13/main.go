package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// f, err := os.Open("input_example.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	patterns := [][]string{{}}
	currentPatternNo := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, "\n")
		if line != "" {
			patterns[currentPatternNo] = append(patterns[currentPatternNo], line)
		} else {
			patterns = append(patterns, []string{})
			currentPatternNo++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := 0
	for _, pattern := range patterns {
		s1 += findColReflection(pattern)
		s1 += 100 * findRowReflection(pattern)
	}
	log.Printf("s1=%d", s1)
}

func transpose(s []string) []string {
	transposed := []string{}
	for col := 0; col < utf8.RuneCountInString(s[0]); col++ {
		line := ""
		for row := 0; row < len(s); row++ {
			line += string([]rune(s[row])[col])
		}
		transposed = append(transposed, line)
	}
	return transposed
}

func findRowReflection(pattern []string) int {
	reflectionFound := false

	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == pattern[i+1] {
			reflectionFound = true
			reflectionSize := 0
			if reflectionSizeRight := len(pattern) - (i + 2); i < reflectionSizeRight {
				reflectionSize = i
			} else {
				reflectionSize = reflectionSizeRight
			}

			for j := 1; j <= reflectionSize; j++ {
				if pattern[i-j] != pattern[i+1+j] {
					reflectionFound = false
				}
			}

			if reflectionFound {
				return i + 1
			}
		}
	}

	return 0
}

func findColReflection(pattern []string) int {
	transposed := transpose(pattern)
	return findRowReflection(transposed)
}
