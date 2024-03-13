package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	seedsList, rangesLists := parseInput("input.txt")
	var locations []int

	for _, seed := range seedsList {
		output := seed
		for _, rangeList := range rangesLists {
			output = getMapping(output, rangeList)
		}
		locations = append(locations, output)
	}

	fmt.Printf("s1=%d", slices.Min(locations))
}

func getMapping(input int, rangeList [][3]int) int {
	slices.SortFunc(rangeList, func(a, b [3]int) int {
		return cmp.Compare(a[0], b[0])
	})
	i, found := slices.BinarySearchFunc(rangeList, input, func(a [3]int, b int) int {
		return cmp.Compare(a[0], b)
	})
	if found {
		return rangeList[i][1]
	} else {
		if i > 0 && input > rangeList[i-1][0] && input < rangeList[i-1][0]+rangeList[i-1][2] {
			return input - rangeList[i-1][0] + rangeList[i-1][1]
		} else {
			return input
		}
	}
}

func parseInput(inputFile string) ([]int, [][][3]int) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	var seedsList []int
	var rangesList [][][3]int

	for scanner.Scan() {
		line := scanner.Text()

		if lineNumber == 0 {
			seedsList = extractSeeds(line)
		}

		if line == "" {
			continue
		}

		re := regexp.MustCompile(`[a-z]+\-to\-[a-z]+ map.*`)
		if re.MatchString(line) {
			rangesList = append(rangesList, extractRange(scanner))
		}

		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return seedsList, rangesList
}

func extractSeeds(line string) []int {
	pattern := regexp.MustCompile(`\d+`)
	match := pattern.FindAllString(line, -1)
	var seedsList []int

	for _, seedNumber := range match {
		n, err := strconv.Atoi(seedNumber)
		if err != nil {
			log.Fatal(err)
		}
		seedsList = append(seedsList, n)
	}

	return seedsList
}

func extractRange(scanner *bufio.Scanner) [][3]int {
	rangeLists := [][3]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		values := strings.Split(line, " ")
		var currentRange [3]int

		if startInputRange, err := strconv.Atoi(values[1]); err != nil {
			log.Fatal(err)
		} else {
			currentRange[0] = startInputRange
		}

		if startOutputRange, err := strconv.Atoi(values[0]); err != nil {
			log.Fatal(err)
		} else {
			currentRange[1] = startOutputRange
		}

		if rangeSize, err := strconv.Atoi(values[2]); err != nil {
			log.Fatal(err)
		} else {
			currentRange[2] = rangeSize
		}

		rangeLists = append(rangeLists, currentRange)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rangeLists
}
