package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	// solution part 1
	s1 := 1
	for i := 0; i < len(timesAndDistances[0]); i++ {
		time := timesAndDistances[0][i]
		distance := timesAndDistances[1][i]
		speed := calculateSpeed(float64(time), float64(distance))
		s1 *= numberOfWaysToBeat(time, distance, speed)
	}
	fmt.Printf("s1=%d\n", s1)

	// solution part 2
	var b strings.Builder
	for _, n := range timesAndDistances[0] {
		fmt.Fprintf(&b, "%d", n)
	}
	time2, err := strconv.Atoi(b.String())
	if err != nil {
		log.Fatal(err)
	}
	b.Reset()
	for _, n := range timesAndDistances[1] {
		fmt.Fprintf(&b, "%d", n)
	}
	distance2, err := strconv.Atoi(b.String())
	if err != nil {
		log.Fatal(err)
	}
	speed2 := calculateSpeed(float64(time2), float64(distance2))
	fmt.Printf("s2=%d", numberOfWaysToBeat(time2, distance2, speed2))
}

func numberOfWaysToBeat(time int, distanceToBeat int, speedOfCurrentRecord float64) int {
	count := 0
	for i := int(speedOfCurrentRecord) + 1; i <= time/2; i++ {
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
