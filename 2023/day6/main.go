package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var timesAndDistances [2][]int
	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		timesAndDistances[lineNumber] = extractNumbers(line)
		lineNumber += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	s1 := 1
	for i := 0; i < len(timesAndDistances[0]); i++ {
		time := timesAndDistances[0][i]
		distance := timesAndDistances[1][i]
		speed := calculateSpeed(float64(time), float64(distance))
		s1 *= numberOfWaysToBeat(time, distance, int(speed)+1)
	}
	fmt.Printf("s1=%d", s1)
}

func numberOfWaysToBeat(time int, distanceToBeat int, speedOfCurrentRecord int) int {
	count := 0
	for i := speedOfCurrentRecord; i <= time/2; i++ {
		currentDistance := i * (time - i)
		if currentDistance > distanceToBeat {
			if time-i != i {
				count += 2
			} else {
				count += 1
			}
		} else {
			break
		}
	}
	return count
}

func extractNumbers(line string) []int {
	var numbers []int
	re := regexp.MustCompile(`\d+`)
	for _, s := range re.FindAllString(line, -1) {
		if n, err := strconv.Atoi(s); err != nil {
			log.Fatal(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func calculateSpeed(time, distance float64) float64 {
	return (time - math.Sqrt(math.Pow(time, 2)-4*distance)) / 2
}
