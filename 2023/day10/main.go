package main

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
)

type PathStep struct {
	Position [2]int
	Path     [][2]int
	Distance int
}

func main() {
	startingPos, pipesMap := getMapAndStartingPos("input.txt")
	maxDistance := getFarthestDistance(startingPos, pipesMap)
	log.Println("s1: ", maxDistance)

	coordinates := getLoopCoordinates(startingPos, pipesMap)
	// remove the ending point which will be the same as the starting point
	coordinates = coordinates[:len(coordinates)-1]

	// Sort coordinates in the loop
	slices.SortFunc(coordinates, func(a, b [2]int) int {
		cmpFirst := cmp.Compare(a[0], b[0])
		if cmpFirst == 0 {
			return cmp.Compare(a[1], b[1])
		}
		return cmpFirst
	})

	startingPosPipe := findPipeForStartingPos(startingPos, coordinates)
	pipesMap[startingPos[0]][startingPos[1]] = startingPosPipe

	enclosed := countEnclosedTiles(coordinates, pipesMap)
	log.Println("s2: ", enclosed)
}

func findPipeForStartingPos(startingPos [2]int, loopCoordinates [][2]int) rune {
	pos := 0
	for i, coord := range loopCoordinates {
		if coord == startingPos {
			pos = i
			break
		}
	}

	next := loopCoordinates[(pos+1)%len(loopCoordinates)]

	var previous [2]int
	if (pos - 1) < 0 {
		previous = loopCoordinates[len(loopCoordinates)-1]
	} else {
		previous = loopCoordinates[(pos-1)%len(loopCoordinates)]
	}

	row, col := startingPos[0], startingPos[1]

	if positions := [][2]int{{row - 1, col}, {row + 1, col}}; slices.Contains(positions, previous) && slices.Contains(positions, next) {
		return '|'
	} else if positions := [][2]int{{row, col - 1}, {row, col + 1}}; slices.Contains(positions, previous) && slices.Contains(positions, next) {
		return '-'
	} else if positions := [][2]int{{row - 1, col}, {row, col + 1}}; slices.Contains(positions, previous) && slices.Contains(positions, next) {
		return 'L'
	} else if positions := [][2]int{{row - 1, col}, {row, col - 1}}; slices.Contains(positions, previous) && slices.Contains(positions, next) {
		return 'J'
	} else if positions := [][2]int{{row - 1, col}, {row, col + 1}}; slices.Contains(positions, previous) && slices.Contains(positions, next) {
		return 'F'
	}

	return '7'
}

func countEnclosedTiles(loopCoordinates [][2]int, pipesMap [][]rune) int {
	rowsSeen := map[int]bool{}
	loopCoordinatesCache := map[[2]int]bool{}
	enclosedTiles := 0

	for _, coord := range loopCoordinates {
		loopCoordinatesCache[coord] = true
	}

	// Start from coordinates in the loop and loop through the whole row
	for _, coord := range loopCoordinates {
		row, col := coord[0], coord[1]

		if _, ok := rowsSeen[row]; ok {
			continue
		}

		enclosed := false
		rowCount := 0
		j := col
		colLen := len(pipesMap[0])

		for j < colLen {
			_, coordInLoop := loopCoordinatesCache[[2]int{row, j}]

			if enclosed && !coordInLoop {
				rowCount++
			}

			currentTile := pipesMap[row][j]

			/* We are inside the loop if either:
			1. we see a '|'
			2. we see a 'F', followed by 0 or more '-' and finally a 'j'
			3. we see a 'L' followed by 0 or more '-' and finally a '7'

			We use the same logic to check when we are getting out of the loop
			*/
			if coordInLoop {
				if currentTile == '|' {
					enclosed = !enclosed
				} else if slices.Contains([]rune{'L', 'F'}, currentTile) {
					startTile := currentTile
					j++
					currentTile = pipesMap[row][j]
					for currentTile == '-' && j < colLen {
						j++
						currentTile = pipesMap[row][j]
					}
					if startTile == 'L' && currentTile == '7' {
						enclosed = !enclosed
					} else if startTile == 'F' && currentTile == 'J' {
						enclosed = !enclosed
					}
				}
			}

			if !enclosed {
				enclosedTiles += rowCount
				rowCount = 0
			}

			j++
		}

		rowsSeen[row] = true
	}

	return enclosedTiles
}

func printMap(pipesMap [][]rune) {
	for _, line := range pipesMap {
		log.Println(string(line))
	}
}

func getLoopCoordinates(startingPos [2]int, pipesMap [][]rune) [][2]int {
	/*
		DFS traversal to find all nodes that make up the loop
	*/
	frontier := []PathStep{
		{
			Position: [2]int{startingPos[0], startingPos[1]},
			Path:     [][2]int{{startingPos[0], startingPos[1]}},
			Distance: 0,
		},
	}
	seen := map[[2]int]bool{
		startingPos: true,
	}

	for len(frontier) > 0 {
		last := len(frontier) - 1
		current := frontier[last]
		// remove element from the stack
		frontier = frontier[:last]

		// get next elements and append to frontier
		for _, nextStep := range getNextPositions(current, pipesMap) {
			row, col := nextStep.Position[0], nextStep.Position[1]
			// nextStep.Distance > 2 avoids going straight back to the starting position
			if pipesMap[row][col] == 'S' && nextStep.Distance > 2 {
				return nextStep.Path
			}
			if _, ok := seen[nextStep.Position]; ok {
				continue
			}
			frontier = append(frontier, nextStep)
		}
		seen[current.Position] = true
	}

	return [][2]int{}
}

func getFarthestDistance(startingPos [2]int, pipesMap [][]rune) int {
	/*
		BFS traversal to find the farthest node
	*/
	maxDistance := 0
	frontier := []PathStep{
		{
			Position: [2]int{startingPos[0], startingPos[1]},
			Path:     [][2]int{{startingPos[0], startingPos[1]}},
			Distance: 0,
		},
	}
	seen := map[[2]int]bool{
		startingPos: true,
	}

	for len(frontier) > 0 {
		current := frontier[0]
		// remove element from queue
		frontier = frontier[1:]

		distance := current.Distance
		if distance > maxDistance {
			maxDistance = distance
		}

		// get next elements and append to frontier
		for _, nextPos := range getNextPositions(current, pipesMap) {
			if _, ok := seen[nextPos.Position]; ok {
				continue
			}
			frontier = append(frontier, nextPos)
		}
		seen[current.Position] = true
	}
	return maxDistance
}

func buildNextStep(nextPosition [2]int, path [][2]int, distance int) PathStep {
	lenPath := len(path)
	nextPath := make([][2]int, lenPath+1)
	copy(nextPath, path)
	nextPath[lenPath] = nextPosition
	return PathStep{
		Position: nextPosition,
		Path:     nextPath,
		Distance: distance,
	}
}

func getNextPositions(step PathStep, pipesMap [][]rune) []PathStep {
	var nextSteps []PathStep
	row, col := step.Position[0], step.Position[1]
	distance := step.Distance
	currentPipe := pipesMap[row][col]
	allowedMoves := map[rune]map[string][]rune{
		'|': {
			"up":   []rune{'|', '7', 'F', 'S'},
			"down": []rune{'|', 'L', 'J', 'S'},
		},
		'-': {
			"right": []rune{'-', 'J', '7', 'S'},
			"left":  []rune{'-', 'L', 'F', 'S'},
		},
		'L': {
			"up":    []rune{'|', '7', 'F', 'S'},
			"right": []rune{'-', 'J', '7', 'S'},
		},
		'J': {
			"left": []rune{'-', 'F', 'L', 'S'},
			"up":   []rune{'|', '7', 'F', 'S'},
		},
		'7': {
			"left": []rune{'-', 'L', 'F', 'S'},
			"down": []rune{'|', 'J', 'L', 'S'},
		},
		'F': {
			"down":  []rune{'|', 'L', 'J', 'S'},
			"right": []rune{'-', 'J', '7', 'S'},
		},
		'S': {
			"up":    []rune{'|', '7', 'F'},
			"down":  []rune{'|', 'L', 'J'},
			"right": []rune{'-', 'J', '7'},
			"left":  []rune{'-', 'L', 'F'},
		},
	}

	// go up
	if row-1 >= 0 {
		pipeAbove := pipesMap[row-1][col]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["up"]; ok && slices.Contains(validNextPipes, pipeAbove) {
				nextPosition := [2]int{row - 1, col}
				nextStep := buildNextStep(nextPosition, step.Path, distance+1)
				nextSteps = append(nextSteps, nextStep)
			}
		}
	}

	// go down
	if row+1 < len(pipesMap) {
		pipeBelow := pipesMap[row+1][col]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["down"]; ok && slices.Contains(validNextPipes, pipeBelow) {
				nextPosition := [2]int{row + 1, col}
				nextStep := buildNextStep(nextPosition, step.Path, distance+1)
				nextSteps = append(nextSteps, nextStep)
			}
		}
	}

	// go right
	if col+1 < len(pipesMap[0]) {
		pipeRight := pipesMap[row][col+1]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["right"]; ok && slices.Contains(validNextPipes, pipeRight) {
				nextPosition := [2]int{row, col + 1}
				nextStep := buildNextStep(nextPosition, step.Path, distance+1)
				nextSteps = append(nextSteps, nextStep)
			}
		}
	}

	// go left
	if col-1 >= 0 {
		pipeLeft := pipesMap[row][col-1]
		if allowed, ok := allowedMoves[currentPipe]; ok {
			if validNextPipes, ok := allowed["left"]; ok && slices.Contains(validNextPipes, pipeLeft) {
				nextPosition := [2]int{row, col - 1}
				nextStep := buildNextStep(nextPosition, step.Path, distance+1)
				nextSteps = append(nextSteps, nextStep)
			}
		}
	}

	return nextSteps
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
