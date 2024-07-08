package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
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

	copy := copyRuneSlice(platform)
	tiltVertical(copy, "north")
	fmt.Printf("s1=%d\n", calculateLoad(copy))

	cycleSize := 0
	cycleCount := 0
	cache := map[uint64]int{}
	for true {
		key := getCacheKey(platform)
		if value, ok := cache[key]; ok {
			cycleSize = cycleCount - value
			break
		} else {
			cache[key] = cycleCount
		}
		cycle(platform)
		cycleCount++
	}

	remaining := (1000_000_000 - cycleCount) % cycleSize
	if remaining > 0 {
		for i := 0; i < remaining; i++ {
			cycle(platform)
		}
	}

	fmt.Printf("s2=%d\n", calculateLoad(platform))
}

func printPlatform(platform [][]rune) {
	for _, runes := range platform {
		fmt.Println(string(runes))
	}
}

func getCacheKey(platform [][]rune) uint64 {
	h := fnv.New64a()
	for _, row := range platform {
		for _, r := range row {
			// Convert rune to bytes and write to the hash function
			var buf [4]byte
			n := copy(buf[:], string(r))
			h.Write(buf[:n])
		}
	}
	return h.Sum64()
}

func copyRuneSlice(original [][]rune) [][]rune {
	copied := make([][]rune, len(original))

	for i, innerSlice := range original {
		copiedInnerSlice := make([]rune, len(innerSlice))
		copy(copiedInnerSlice, innerSlice)
		copied[i] = copiedInnerSlice
	}

	return copied
}

func cycle(platform [][]rune) {
	for _, direction := range []string{"north", "west", "south", "east"} {
		if direction == "north" || direction == "south" {
			tiltVertical(platform, direction)
		}
		if direction == "east" || direction == "west" {
			tiltHorizontal(platform, direction)
		}
	}
}

func tiltVertical(platform [][]rune, direction string) {
	for j := 0; j < len(platform[0]); j++ {
		freePosition := -1
		for i := 0; i < len(platform); i++ {
			var row int
			if direction == "north" {
				row = i
			} else {
				row = len(platform) - 1 - i
			}
			current := platform[row][j]
			if current == '.' && freePosition == -1 {
				freePosition = row
			}
			// take into account # that blocks sliding
			if current == '#' {
				freePosition = -1
			}
			if current == 'O' && freePosition != -1 {
				platform[freePosition][j] = 'O'
				platform[row][j] = '.'
				if direction == "north" {
					freePosition++
				} else {
					freePosition--
				}
			}
		}
	}
}

func tiltHorizontal(platform [][]rune, direction string) {
	for i := 0; i < len(platform); i++ {
		freePosition := -1
		for j := 0; j < len(platform[0]); j++ {
			var col int
			if direction == "west" {
				col = j
			} else {
				col = len(platform[0]) - 1 - j
			}
			current := platform[i][col]
			if current == '.' && freePosition == -1 {
				freePosition = col
			}
			// take into account # that blocks sliding
			if current == '#' {
				freePosition = -1
			}
			if current == 'O' && freePosition != -1 {
				platform[i][freePosition] = 'O'
				platform[i][col] = '.'
				if direction == "west" {
					freePosition++
				} else {
					freePosition--
				}
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
