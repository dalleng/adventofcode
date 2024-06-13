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
	s2 := 0

	for _, pattern := range patterns {
		s1 += findColReflection(pattern)
		s2 += findColReflectionWithSmudge(pattern)
		s1 += 100 * findRowReflection(pattern)
		s2 += 100 * findRowReflectionWithSmudge(pattern)
	}

	log.Printf("s1=%d", s1)
	log.Printf("s2=%d", s2)
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

func isEqualWithDiffTolerance(s1 string, s2 string, allowedDiff int) (bool, []int) {
	s1AsRunes := []rune(s1)
	s2AsRunes := []rune(s2)
	diffCount := 0
	diffIndices := []int{}
	for i := 0; i < utf8.RuneCountInString(s1); i++ {
		if s1AsRunes[i] != s2AsRunes[i] {
			diffCount += 1
			diffIndices = append(diffIndices, i)
		}
	}
	equal := diffCount <= allowedDiff
	return equal, diffIndices
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

func findRowReflectionWithSmudge(pattern []string) int {
	reflectionFound := false

	for i := 0; i < len(pattern)-1; i++ {
		isEqualWithTolerance, indices := isEqualWithDiffTolerance(pattern[i], pattern[i+1], 1)
		smudgeFound := len(indices) > 0
		if isEqualWithTolerance {
			reflectionFound = true
			reflectionSize := 0

			if reflectionSizeRight := len(pattern) - (i + 2); i < reflectionSizeRight {
				reflectionSize = i
			} else {
				reflectionSize = reflectionSizeRight
			}

			for j := 1; j <= reflectionSize; j++ {
				diffTolerance := 0
				if !smudgeFound {
					diffTolerance = 1
				}
				isEqualWithTolerance, indices := isEqualWithDiffTolerance(pattern[i-j], pattern[i+1+j], diffTolerance)
				if !smudgeFound {
					smudgeFound = len(indices) > 0
				}
				if !isEqualWithTolerance {
					reflectionFound = false
				}
			}

			if reflectionFound && smudgeFound {
				return i + 1
			}
		}
	}

	return 0
}

func findColReflectionWithSmudge(pattern []string) int {
	transposed := transpose(pattern)
	return findRowReflectionWithSmudge(transposed)
}

func findColReflection(pattern []string) int {
	transposed := transpose(pattern)
	return findRowReflection(transposed)
}
