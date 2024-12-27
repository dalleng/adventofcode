package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Step struct {
	Position  [2]int
	StepCount int
}

func main() {
	/*
		Points to consider:

		1. The Manhattan distance is always going to be the shortest path in a grid with only vertical and horizontal moves.

		Assumptions for Part 2

		 1. The grid is square.
		 2. Borders are free.
		 3. The horizontal and vertical lines that go through the center are free.
		 4. S, the start, is always in the center of the grid.
		 5. The total numbers of steps 26501365 = 202,300 * 131 + 65.

		 Explanation of the geometric solution for Part 2

		- https://advent-of-code.xavd.id/writeups/2023/day/21/
	*/

	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	start := [2]int{}
	lineCount := 0
	garden := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		garden = append(garden, line)
		lineCount++
		if index := strings.Index(line, "S"); index != -1 {
			// lineCount-1 for 0-based indexing
			start[0], start[1] = lineCount-1, index
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("s1=%d\n", len(tilesReachedInSteps(garden, Step{Position: start, StepCount: 0}, 64)))

	steps := 26501365
	n := steps / len(garden[0])

	reachedOdd := tilesReachedInSteps(garden, Step{Position: start, StepCount: 0}, steps)
	reachedEven := tilesReachedInSteps(garden, Step{Position: start, StepCount: 1}, steps)
	nReachedOdd := len(reachedOdd)
	nReachedEven := len(reachedEven)

	nReachedOddCorner := 0
	for _, distance := range reachedOdd {
		if distance > len(garden[0])/2 {
			nReachedOddCorner++
		}
	}

	nReachedEvenCorner := 0
	for _, distance := range reachedEven {
		if distance > len(garden[0])/2 {
			nReachedEvenCorner++
		}
	}

	s2 := (n+1)*(n+1)*nReachedOdd + n*n*nReachedEven + n*nReachedEvenCorner - (n+1)*nReachedOddCorner
	fmt.Printf("s2=%d\n", s2)
}

func getNextPositions(garden []string, start [2]int) [][2]int {
	nextPositions := [][2]int{}
	positionDiff := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, diff := range positionDiff {
		// out of bounds check
		nextPos := [2]int{start[0] + diff[0], start[1] + diff[1]}
		if nextPos[0] < 0 || nextPos[0] >= len(garden) {
			continue
		}
		if nextPos[1] < 0 || nextPos[1] >= len(garden[0]) {
			continue
		}
		// rock check
		if garden[nextPos[0]][nextPos[1]] == '#' {
			continue
		}
		nextPositions = append(nextPositions, nextPos)
	}
	return nextPositions
}

func tilesReachedInSteps(garden []string, start Step, maxSteps int) map[[2]int]int {
	/*
		Returns a map of reachable tile positions and the minimum steps needed to reach them within maxSteps.

		The function uses BFS (Breadth-First Search) to:
		1. Track positions by steps taken in ascending order
		2. Stop when maxSteps is exceeded or all positions are explored
		3. Count a tile as reachable if either:
		   - It takes exactly maxSteps to reach it
		   - It can be revisited in maxSteps (when maxSteps minus steps_to_reach is even)

		Parameters:
		- garden: Grid of plots ('.') and rocks ('#')
		- start: Starting position [row, col]
		- maxSteps: Maximum steps allowed
		- isInfiniteGarden: Whether garden repeats infinitely

		Returns:
		- Map of positions to minimum steps needed to reach them
	*/

	seen := map[[2]int]bool{}
	reached := map[[2]int]int{}
	frontier := []Step{start}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if current.StepCount > maxSteps {
			break
		}

		if current.StepCount == maxSteps || (maxSteps-current.StepCount)%2 == 0 {
			if _, ok := reached[current.Position]; !ok {
				reached[current.Position] = current.StepCount
			}
		}

		for _, nextPos := range getNextPositions(garden, current.Position) {
			if _, ok := seen[nextPos]; !ok {
				frontier = append(frontier, Step{Position: nextPos, StepCount: current.StepCount + 1})
				seen[nextPos] = true
			}
		}
	}

	return reached
}
