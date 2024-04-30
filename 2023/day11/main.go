package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

func main() {
	galaxyMap := readInput("input.txt")
	galaxies := findGalaxies(galaxyMap)
	s1 := getSumOfAllShortestPaths(galaxyMap, galaxies, 2)
	fmt.Println(s1)
	s2 := getSumOfAllShortestPaths(galaxyMap, galaxies, 1000000)
	fmt.Println(s2)
}

func getSumOfAllShortestPaths(galaxyMap [][]rune, galaxies [][2]int, expansionScale int) int {
	rowsToExpand, colsToExpand := findRowsAndColsToExpand(galaxyMap)
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			from := galaxies[i]
			to := galaxies[j]
			for row := int(math.Min(float64(from[0]), float64(to[0]))); row < int(math.Max(float64(from[0]), float64(to[0]))); row++ {
				if slices.Contains(rowsToExpand, row) {
					sum += expansionScale
				} else {
					sum += 1
				}
			}
			for col := int(math.Min(float64(from[1]), float64(to[1]))); col < int(math.Max(float64(from[1]), float64(to[1]))); col++ {
				if slices.Contains(colsToExpand, col) {
					sum += expansionScale
				} else {
					sum += 1
				}
			}
		}
	}
	return sum
}

func findGalaxies(galaxyMap [][]rune) [][2]int {
	galaxies := make([][2]int, 0)
	for i := 0; i < len(galaxyMap); i++ {
		for j := 0; j < len(galaxyMap[0]); j++ {
			if galaxyMap[i][j] == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}
	return galaxies
}

func findRowsAndColsToExpand(galaxyMap [][]rune) ([]int, []int) {
	// figure out which rows need to be expanded
	rowsToExpand := make([]int, 0)
	for i := 0; i < len(galaxyMap); i++ {
		rowsToExpand = append(rowsToExpand, i)
		for j := 0; j < len(galaxyMap[0]); j++ {
			if galaxyMap[i][j] == '#' {
				rowsToExpand = rowsToExpand[:len(rowsToExpand)-1]
				break
			}
		}
	}

	// figure out which cols need to be expanded
	colsToExpand := make([]int, 0)
	for j := 0; j < len(galaxyMap[0]); j++ {
		colsToExpand = append(colsToExpand, j)
		for i := 0; i < len(galaxyMap); i++ {
			if galaxyMap[i][j] == '#' {
				colsToExpand = colsToExpand[:len(colsToExpand)-1]
				break
			}
		}
	}
	return rowsToExpand, colsToExpand
}

func readInput(filepath string) [][]rune {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var galaxyMap [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		galaxyMap = append(galaxyMap, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return galaxyMap
}
