package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	startingPos, pipesMap := getMapAndStartingPos("input.txt")
	printMap(pipesMap)
	log.Println("startingPos: ", startingPos)
	maxDistance := getFarthestDistance(startingPos, pipesMap)
	log.Println("maxDistance: ", maxDistance)
}

func printMap(pipesMap [][]rune) {
	for _, line := range pipesMap {
		fmt.Println(string(line))
	}
}

func getFarthestDistance(startingPos [2]int, pipesMap [][]rune) int {
	maxDistance := 0
	frontier := [][3]int{{startingPos[0], startingPos[1], 0}}
	seen := map[[2]int]bool{
		startingPos: true,
	}

	for len(frontier) > 0 {
		current := frontier[0]
		// remove element from queue
		frontier = frontier[1:]

		distance := current[2]
		if distance > maxDistance {
			maxDistance = distance
		}

		// get next elements and append to frontier
		for _, nextPos := range getNextPositions(current, pipesMap) {
			log.Println("nextPos: ", nextPos)
			if _, ok := seen[[2]int(nextPos[:2])]; ok {
				log.Println("position has already been seen")
				continue
			}
			frontier = append(frontier, nextPos)
		}
		position := [2]int(current[:2])
		seen[position] = true
	}
	return maxDistance
}

func getNextPositions(currentPosAndDistance [3]int, pipesMap [][]rune) [][3]int {
	var nextPositions [][3]int
	row, col := currentPosAndDistance[0], currentPosAndDistance[1]
	distance := currentPosAndDistance[2]
	currentPipe := pipesMap[row][col]
	allowedMoves := map[rune]map[string][]rune{
		'|': {
			"up":   []rune{'|', '7', 'F'},
			"down": []rune{'|', 'L', 'J'},
		},
		'-': {
			"right": []rune{'-', 'J', '7'},
			"left":  []rune{'-', 'L', 'F'},
		},
		'L': {
			"up":    []rune{'|', '7', 'F'},
			"right": []rune{'-', 'J', '7'},
		},
		'J': {
			"left": []rune{'-', 'F', 'L'},
			"up":   []rune{'|', '7', 'F'},
		},
		'7': {
			"left": []rune{'-', 'L', 'F'},
			"down": []rune{'|', 'J', 'L'},
		},
		'F': {
			"down":  []rune{'|', 'L', 'J'},
			"right": []rune{'-', 'J', '7'},
		},
		'S': {
			"up":    []rune{'|', '7', 'F'},
			"down":  []rune{'|', 'L', 'J'},
			"right": []rune{'-', 'J', '7'},
			"left":  []rune{'-', 'L', 'F'},
		},
	}

	if row-1 >= 0 {
		pipeAbove := pipesMap[row-1][col]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["up"]; ok && slices.Contains(validNextPipes, pipeAbove) {
				nextPositions = append(nextPositions, [3]int{row - 1, col, distance + 1})
			}
		}
	}

	// go down
	if row+1 < len(pipesMap) {
		pipeBelow := pipesMap[row+1][col]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["down"]; ok && slices.Contains(validNextPipes, pipeBelow) {
				nextPositions = append(nextPositions, [3]int{row + 1, col, distance + 1})
			}
		}
	}

	// go right
	if col+1 < len(pipesMap[0]) {
		pipeRight := pipesMap[row][col+1]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["right"]; ok && slices.Contains(validNextPipes, pipeRight) {
				nextPositions = append(nextPositions, [3]int{row, col + 1, distance + 1})
			}
		}
	}

	// go left
	if col-1 >= 0 {
		pipeLeft := pipesMap[row][col-1]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["left"]; ok && slices.Contains(validNextPipes, pipeLeft) {
				nextPositions = append(nextPositions, [3]int{row, col - 1, distance + 1})
			}
		}
	}

	return nextPositions
}

func getMapAndStartingPos(inputFile string) ([2]int, [][]rune) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var pipesMap [][]rune
	var startingPos [2]int
	scanner := bufio.NewScanner(f)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		var pipesRow []rune
		for col, char := range line {
			pipesRow = append(pipesRow, char)
			if char == 'S' {
				startingPos = [2]int{row, col}
			}
		}
		pipesMap = append(pipesMap, pipesRow)
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return startingPos, pipesMap
}
