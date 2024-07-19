package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
	STILL
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case STILL:
		return "STILL"
	default:
		return "UNKNOWN"
	}
}

type Position [2]int

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p[0], p[1])
}

type Movement struct {
	Position                Position
	Direction               Direction
	StepsInCurrentDirection int
}
type QueueElement struct {
	Cost     int
	Movement Movement
}

type SearchQueue []QueueElement

func (q QueueElement) String() string {
	return fmt.Sprintf(
		"{Position: %s, Direction: %s, Cost: %d, StepsInCurrentDirection: %d}",
		q.Movement.Position, q.Movement.Direction, q.Cost, q.Movement.StepsInCurrentDirection,
	)
}

func (sq SearchQueue) Len() int {
	return len(sq)
}

func (sq SearchQueue) Less(i, j int) bool {
	return sq[i].Cost < sq[j].Cost
}

func (sq SearchQueue) Swap(i, j int) {
	sq[i], sq[j] = sq[j], sq[i]
}

func (sq *SearchQueue) Push(x any) {
	*sq = append(*sq, x.(QueueElement))
}

func (sq *SearchQueue) Pop() any {
	old := *sq
	n := len(old)
	x := old[n-1]
	*sq = old[0 : n-1]
	return x
}

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	_, cost := search(grid, Position{0, 0}, Position{len(grid) - 1, len(grid[0]) - 1})
	fmt.Printf("s1=%d\n", cost)
}

func getPositionDiffForDirection(d Direction) Position {
	positionDiffs := map[Direction]Position{
		UP:    {-1, 0},
		DOWN:  {1, 0},
		LEFT:  {0, -1},
		RIGHT: {0, 1},
	}
	if diff, ok := positionDiffs[d]; ok {
		return diff
	}
	return Position{0, 0}
}

func getOppositeDirection(d Direction) Direction {
	opposites := map[Direction]Direction{
		UP:    DOWN,
		DOWN:  UP,
		LEFT:  RIGHT,
		RIGHT: LEFT,
	}
	if opposite, ok := opposites[d]; ok {
		return opposite
	}
	return d
}

func isPositionOutOfBounds(board [][]rune, p Position) bool {
	row, col := p[0], p[1]
	if (row >= 0 && row < len(board)) && (col >= 0 && col < len(board[row])) {
		return false
	}
	return true
}

func getNextElements(board [][]rune, current QueueElement) []QueueElement {
	nextElements := []QueueElement{}
	for _, direction := range []Direction{UP, DOWN, LEFT, RIGHT} {
		if direction == getOppositeDirection(current.Movement.Direction) {
			continue
		}

		diff := getPositionDiffForDirection(direction)
		newPosition := Position{current.Movement.Position[0] + diff[0], current.Movement.Position[1] + diff[1]}
		if isPositionOutOfBounds(board, newPosition) {
			continue
		}

		nextCost, _ := strconv.Atoi(string(board[newPosition[0]][newPosition[1]]))
		newStepCounter := 1

		if direction == current.Movement.Direction {
			newStepCounter = current.Movement.StepsInCurrentDirection + 1
		}

		if newStepCounter > 3 {
			continue
		}

		nextElements = append(nextElements, QueueElement{
			Cost: current.Cost + nextCost,
			Movement: Movement{
				Position:                newPosition,
				Direction:               direction,
				StepsInCurrentDirection: newStepCounter,
			},
		})
	}
	return nextElements
}

func search(board [][]rune, origin Position, destination Position) (Position, int) {
	frontier := &SearchQueue{}
	seen := map[Movement]bool{}
	heap.Init(frontier)
	initial := QueueElement{
		Cost: 0,
		Movement: Movement{
			Position:                origin,
			StepsInCurrentDirection: 0,
			Direction:               STILL,
		},
	}
	heap.Push(frontier, initial)
	seen[initial.Movement] = true

	for frontier.Len() > 0 {
		current := heap.Pop(frontier)

		if current.(QueueElement).Movement.Position == destination {
			return current.(QueueElement).Movement.Position, current.(QueueElement).Cost
		}

		for _, nextElement := range getNextElements(board, current.(QueueElement)) {
			if _, ok := seen[nextElement.Movement]; !ok {
				heap.Push(frontier, nextElement)
				seen[nextElement.Movement] = true
			}
		}
	}

	return initial.Movement.Position, initial.Cost
}
