package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	platform := [][]rune{}
	lineno := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := []rune{}
		for _, chr := range []rune(line) {
			lineSlice = append(lineSlice, chr)
		}
		platform = append(platform, lineSlice)
		lineno++
	}

	for _, runes := range platform {
		fmt.Println(string(runes))
	}

	tilt(platform)
	fmt.Println()

	for _, runes := range platform {
		fmt.Println(string(runes))
	}

	fmt.Printf("s1=%d\n", calculateLoad(platform))
}

func tilt(platform [][]rune) {
	for j := 0; j < len(platform[0]); j++ {
		freePosition := -1
		for i := 0; i < len(platform); i++ {
			current := platform[i][j]
			if current == '.' && freePosition == -1 {
				freePosition = i
			}
			// take into account # that blocks sliding
			if current == '#' {
				freePosition = -1
			}
			if current == 'O' && freePosition != -1 {
				platform[freePosition][j] = 'O'
				platform[i][j] = '.'
				freePosition++
			}
		}
	}
}

func calculateLoad(platform [][]rune) int {
	total := 0
	for i := 0; i < len(platform); i++ {
		rowCounter := 0
		for j := 0; j < len(platform[0]); j++ {
			if platform[i][j] == 'O' {
				rowCounter++
			}
		}
		total += rowCounter * (len(platform) - i)
	}
	return total
}
