package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func getPathLength(navigation map[string][2]string, movements string, origin string, isFinalPosition func(string) bool) int {
	steps := 0
	currentPosition := origin
	var direction int
	for !isFinalPosition(currentPosition) {
		for _, mov := range movements {
			if mov == 'L' {
				direction = 0
			} else {
				direction = 1
			}
			currentPosition = navigation[currentPosition][direction]
			steps += 1
		}
	}
	return steps
}

func gcd(a, b int) int {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a int, b int, integers ...int) int {
	lcm := a * b / gcd(a, b)
	for _, i := range integers {
		lcm = lcm * i / gcd(lcm, i)
	}
	return lcm
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var movements string
	navigation := make(map[string][2]string)
	lineno := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lineno == 0 {
			movements = line
			lineno++
			continue
		}

		if line == "" {
			lineno++
			continue
		}

		re := regexp.MustCompile(`([A-Z1-9]{3}) = \(([A-Z1-9]{3})\, ([A-Z1-9]{3})\)`)
		submatches := re.FindStringSubmatch(line)

		navigation[submatches[1]] = [2]string{submatches[2], submatches[3]}
		lineno++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("s1=%d", getPathLength(navigation, movements, "AAA", func(position string) bool {
		return position == "ZZZ"
	}))

	var startingPositions []string
	for key := range navigation {
		if key[2] == 'A' {
			startingPositions = append(startingPositions, key)
		}
	}
	var pathLengths []int
	for _, pos := range startingPositions {
		pathLengths = append(pathLengths, getPathLength(navigation, movements, pos, func(s string) bool {
			return s[2] == 'Z'
		}))
	}
	fmt.Println(pathLengths)
	fmt.Printf("s2=%d", lcm(pathLengths[0], pathLengths[1], pathLengths[2:]...))
}
