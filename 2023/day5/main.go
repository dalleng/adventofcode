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
	solutionPart1(seedsList, rangesLists)
	solutionPart2(seedsList, rangesLists)
	// fmt.Println(getRangeMapping([2]int{79, 14}, [][3]int{{98, 50, 2}, {50, 52, 48}}))             // [[81 14]]
	// fmt.Println(getRangeMapping([2]int{79, 14}, [][3]int{{98, 50, 2}, {79, 50, 7}, {86, 57, 7}})) // [[50 7] [57 7]]
	// fmt.Println(getRangeMapping([2]int{79, 14}, [][3]int{{98, 50, 2}}))                           // [[79 14]]
	// fmt.Println(getRangeMapping([2]int{79, 14}, [][3]int{{98, 50, 2}, {79, 50, 7}}))              // [[50 7] [86 7]]
	// fmt.Println(getRangeMapping([2]int{74, 14}, [][3]int{{77, 45, 23}}))                          // [[74 3] [45 11]]
}

func solutionPart1(seedsList []int, rangesLists [][][3]int) {
	var locations []int
	for _, seed := range seedsList {
		output := seed
		for _, rangeList := range rangesLists {
			output = getMapping(output, rangeList)
		}
		locations = append(locations, output)
	}
	fmt.Printf("s1=%d\n", slices.Min(locations))
}

func solutionPart2(seedsList []int, rangesLists [][][3]int) {
	var locationRanges [][2]int
	for i := 0; i < len(seedsList)-1; i += 2 {
		var currentRanges [][2]int
		seed := seedsList[i]
		rangeSize := seedsList[i+1]
		currentRanges = append(currentRanges, [2]int{seed, rangeSize})
		for _, rangeList := range rangesLists {
			var newRanges [][2]int
			for _, currentRange := range currentRanges {
				newRanges = append(newRanges, getRangeMapping(currentRange, rangeList)...)
			}
			currentRanges = newRanges
		}
		locationRanges = append(locationRanges, currentRanges...)
	}
	min := slices.MinFunc(locationRanges, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})
	fmt.Printf("s2=%d\n", min[0])
}

func getRangeMapping(inputRange [2]int, rangeList [][3]int) [][2]int {
	var outputRange [][2]int
	slices.SortFunc(rangeList, func(a, b [3]int) int {
		return cmp.Compare(a[0], b[0])
	})
	rangeStart := inputRange[0]
	rangeSize := inputRange[1]
	for rangeSize > 0 {
		i, found := slices.BinarySearchFunc(rangeList, rangeStart, func(a [3]int, b int) int {
			return cmp.Compare(a[0], b)
		})
		var rangeConsumed int
		if found {
			if rangeSize > rangeList[i][2] {
				rangeConsumed = rangeList[i][2]
			} else {
				rangeConsumed = rangeSize
			}
			outputRange = append(outputRange, [2]int{rangeList[i][1], rangeConsumed})
			rangeSize -= rangeConsumed
			rangeStart += rangeConsumed
		} else {
			if i > 0 && rangeStart < rangeList[i-1][0]+rangeList[i-1][2] {
				rangeConsumed = (rangeList[i-1][0] + rangeList[i-1][2]) - rangeStart
				if rangeConsumed > rangeSize {
					rangeConsumed = rangeSize
				}
				outputRange = append(outputRange, [2]int{rangeStart - rangeList[i-1][0] + rangeList[i-1][1], rangeConsumed})
				rangeSize -= rangeConsumed
				rangeStart += rangeConsumed
			} else if rangeStart+rangeSize-1 >= rangeList[i][0] && rangeStart+rangeSize-1 < rangeList[i][0]+rangeList[i][2] {
				rangeConsumed = rangeList[i][0] - rangeStart
				outputRange = append(outputRange, [2]int{rangeStart, rangeConsumed})
				rangeSize -= rangeConsumed
				rangeStart += rangeConsumed
			} else {
				outputRange = append(outputRange, [2]int{rangeStart, rangeSize})
				rangeStart += rangeSize
				rangeSize = 0
			}
		}
	}
	return outputRange
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
