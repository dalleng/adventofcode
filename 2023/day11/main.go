package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	galaxyMap := readInput("input.txt")
	// printMap(galaxyMap)
	galaxyMap = addExpansions(galaxyMap)
	// printMap(galaxyMap)

	sum := 0
	sum2 := 0
	galaxies := findGalaxies(galaxyMap)
	fmt.Println(galaxies)
	fmt.Println(len(galaxies))

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum2++
			from := galaxies[i]
			to := galaxies[j]
			log.Printf("Distance between (%d, %d) and (%d, %d) is %d", from[0], from[1], to[0], to[1], int(math.Abs(float64(to[0]-from[0]))+math.Abs(float64(to[1]-from[1]))))
			sum += int(math.Abs(float64(to[0]-from[0])) + math.Abs(float64(to[1]-from[1])))
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
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

func printMap(galaxyMap [][]rune) {
	for _, line := range galaxyMap {
		log.Println(string(line))
	}
}

func addExpansions(galaxyMap [][]rune) [][]rune {
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

	slices.Reverse(rowsToExpand)
	for _, row := range rowsToExpand {
		log.Printf("The row to expand is: %d\n", row)
		galaxyMap = slices.Insert(galaxyMap, row, []rune(strings.Repeat(".", len(galaxyMap[0]))))
	}

	slices.Reverse(colsToExpand)
	for _, col := range colsToExpand {
		log.Printf("The col to expand is: %d\n", col)
		for i := 0; i < len(galaxyMap); i++ {
			galaxyMap[i] = slices.Insert(galaxyMap[i], col, '.')
		}
	}

	return galaxyMap
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
