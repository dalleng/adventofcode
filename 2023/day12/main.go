package main

import (
	"bufio"
	"fmt"
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
	s2 := 0

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

		// Unfold and call count again
		countsUnfolded := make([]int, 0, len(counts)*5)
		springsUnfolded := ""

		for i := 0; i < 5; i++ {
			springsUnfolded += springsAndCounts[0]
			if i < 4 {
				springsUnfolded += "?"
			}
			for _, n := range counts {
				countsUnfolded = append(countsUnfolded, n)
			}
		}

		s2 += count(springsUnfolded, countsUnfolded)
	}
}

func count(springs string, counts []int) int {
	cache := make(map[string]int)

	generateCacheKey := func(springs string, counts []int) string {
		countStr := ""
		for _, n := range counts {
			countStr += fmt.Sprintf("%d", n)
		}
		return fmt.Sprintf("%s%s", springs, countStr)
	}

	var inner func(string, []int) int
	inner = func(springs string, counts []int) int {
		key := generateCacheKey(springs, counts)
		if value, ok := cache[key]; ok {
			return value
		}
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
			result += inner(string([]rune(springs)[1:]), counts)
		}

		if current := []rune(springs)[0]; current == '#' || current == '?' {
			if counts[0] <= len([]rune(springs)) && !strings.Contains(string([]rune(springs)[:counts[0]]), ".") && (len([]rune(springs)) == counts[0] || []rune(springs)[counts[0]] != '#') {
				springsLeft := ""
				if counts[0]+1 < len([]rune(springs)) {
					springsLeft = string([]rune(springs)[counts[0]+1:])
				}
				result += inner(springsLeft, counts[1:])
			}
		}

		cache[key] = result
		return result
	}

	return inner(springs, counts)
}
