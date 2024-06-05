package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	s1 := 0

	for scanner.Scan() {
		line := scanner.Text()
		springsAndCounts := strings.Split(line, " ")
		counts := make([]int, 0)
		for _, n := range strings.Split(springsAndCounts[1], ",") {
			if value, err := strconv.Atoi(n); err != nil {
				log.Fatal(err)
			} else {
				counts = append(counts, value)
			}
		}
		s1 += count(springsAndCounts[0], counts)
	}

	log.Printf("s1: %d", s1)
}

func count(springs string, counts []int) int {
	if springs == "" {
		if len(counts) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(counts) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		} else {
			return 1
		}
	}

	result := 0

	if current := []rune(springs)[0]; current == '.' || current == '?' {
		result += count(string([]rune(springs)[1:]), counts)
	}

	if current := []rune(springs)[0]; current == '#' || current == '?' {
		if counts[0] <= len([]rune(springs)) && !strings.Contains(string([]rune(springs)[:counts[0]]), ".") && (len([]rune(springs)) == counts[0] || []rune(springs)[counts[0]] != '#') {
			springsLeft := ""
			if counts[0]+1 < len([]rune(springs)) {
				springsLeft = string([]rune(springs)[counts[0]+1:])
			}
			result += count(springsLeft, counts[1:])
		}
	}

	return result
}
