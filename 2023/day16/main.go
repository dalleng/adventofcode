package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	RIGHT
	LEFT
)

type Movement struct {
	Position  [2]int
	Direction Direction
}

func main() {
	// f, err := os.Open("example.txt")
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	board := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, []rune(line))
	}
	visited := traversal(board, Movement{Position: [2]int{0, 0}, Direction: RIGHT})
	fmt.Printf("s1=%d\n", len(visited))

}

func getNextDirections(board [][]rune, current Movement) []Direction {
	var transitionRules = map[rune]map[Direction][]Direction{
		'.': {
			UP:    {UP},
			DOWN:  {DOWN},
			LEFT:  {LEFT},
			RIGHT: {RIGHT},
		},
		'/': {
			RIGHT: {UP},
			DOWN:  {LEFT},
			LEFT:  {DOWN},
			UP:    {RIGHT},
		},
		'\\': {
			RIGHT: {DOWN},
			DOWN:  {RIGHT},
			LEFT:  {UP},
			UP:    {LEFT},
		},
		'|': {
			RIGHT: {UP, DOWN},
			LEFT:  {UP, DOWN},
			UP:    {UP},
			DOWN:  {DOWN},
		},
		'-': {
			UP:    {LEFT, RIGHT},
			DOWN:  {LEFT, RIGHT},
			LEFT:  {LEFT},
			RIGHT: {RIGHT},
		},
	}

	nextDirections := []Direction{}
	currentCell := board[current.Position[0]][current.Position[1]]
	if directions, ok := transitionRules[currentCell][current.Direction]; ok {
		nextDirections = directions
	}
	return nextDirections
}

func getNextMovements(board [][]rune, current Movement) []Movement {
	positionDiffForDirection := map[Direction][2]int{
		UP:    {-1, 0},
		DOWN:  {1, 0},
		LEFT:  {0, -1},
		RIGHT: {0, 1},
	}

	nextDirections := getNextDirections(board, current)
	nextMovements := []Movement{}
	for _, direction := range nextDirections {
		positionDiff := positionDiffForDirection[direction]
		nextPosition := [2]int{current.Position[0] + positionDiff[0], current.Position[1] + positionDiff[1]}
		if (nextPosition[0] >= 0 && nextPosition[0] < len(board)) && (nextPosition[1] >= 0 && nextPosition[1] < len(board[0])) {
			nextMovement := Movement{
				Position:  nextPosition,
				Direction: direction,
			}
			nextMovements = append(nextMovements, nextMovement)
		}
	}

	return nextMovements
}

func traversal(board [][]rune, start Movement) map[string]bool {
	frontier := []Movement{start}
	seen := map[Movement]bool{}
	visited := map[string]bool{}

	for len(frontier) > 0 {
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]
		if _, ok := seen[current]; ok {
			continue
		}
		seen[current] = true
		key := fmt.Sprintf("%d,%d", current.Position[0], current.Position[1])
		if _, ok := visited[key]; !ok {
			visited[key] = true
		}
		for _, nextMovement := range getNextMovements(board, current) {
			frontier = append(frontier, nextMovement)
		}
	}

	return visited
}
