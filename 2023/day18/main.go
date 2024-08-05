package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	Direction rune
	Steps     int
}

func (m Move) String() string {
	return fmt.Sprintf("{Direction: %c, Steps: %d}", m.Direction, m.Steps)
}

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	digPlan := []Move{}
	digPlan2 := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if steps, err := strconv.Atoi(lineSplit[1]); err == nil {
			digPlan = append(digPlan, Move{
				Direction: []rune(lineSplit[0])[0],
				Steps:     steps,
			})
		} else {
			log.Fatal(err)
		}
		digPlan2 = append(digPlan2, parseColor(lineSplit[2]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := getArea(digPlan)
	fmt.Printf("s1=%d\n", s1)

	s2 := getArea(digPlan2)
	fmt.Printf("s1=%d\n", s2)
}

func parseColor(color string) Move {
	stepsHex := []rune(color)[2 : len(color)-2]

	// Add 3 zeros to the left to get the hex string to be decoded into 4 bytes
	decoded, err := hex.DecodeString("000" + string(stepsHex))

	if err != nil {
		log.Fatal(err)
	}

	directionRune := []rune(color)[len(color)-2]
	directionMap := map[rune]rune{'0': 'R', '1': 'D', '2': 'L', '3': 'U'}
	direction, _ := directionMap[directionRune]

	steps := binary.BigEndian.Uint32(decoded)
	if err != nil {
		log.Fatal(err)
	}

	return Move{
		Direction: direction,
		Steps:     int(steps),
	}
}

func getArea(moves []Move) int {
	vertices := getVertices(moves)
	perimeter := getPerimeter(moves)

	area := 0
	for i := 0; i < len(vertices)-1; i++ {
		x1, y1 := vertices[i][0], vertices[i][1]
		x2, y2 := vertices[i+1][0], vertices[i+1][1]
		area += (x1*y2 - x2*y1)
	}

	area /= 2
	if area < 0 {
		area *= -1
	}

	return area + perimeter/2 + 1
}

func getPerimeter(digPlan []Move) int {
	perimeter := 0
	for _, move := range digPlan {
		perimeter += move.Steps
	}
	return perimeter
}

func getVertices(digPlan []Move) [][2]int {
	current := [2]int{0, 0}
	coords := [][2]int{current}
	directionMap := map[rune][2]int{
		'U': {-1, 0},
		'D': {1, 0},
		'L': {0, -1},
		'R': {0, 1},
	}
	for _, move := range digPlan {
		positionDiff, _ := directionMap[move.Direction]
		current[0], current[1] = current[0]+positionDiff[0]*move.Steps, current[1]+positionDiff[1]*move.Steps
		coords = append(coords, current)
	}
	return coords
}
