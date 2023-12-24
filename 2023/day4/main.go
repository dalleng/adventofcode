package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "strings"
    "strconv"
    "math"
)


func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(f)

    s := 0
    for scanner.Scan() {
        line := scanner.Text()
        numbers := extractNumbers(line)
        winning, own := numbers[0], numbers[1]
        count := getIntersectionCount(winning, own)
        s += int(math.Pow(2, float64(count-1)))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    log.Print("result: ", s)
}


func getIntersectionCount(winning []int, own []int) int {
    ownMemo := map[int]bool{}
    for _, n := range own {
        ownMemo[n] = true
    }
    count := 0
    for _, winning := range winning {
        if _, ok := ownMemo[winning]; ok {
            count += 1
        }
    }
    return count
} 

func extractNumbers(line string) [2][]int {
    pattern := regexp.MustCompile(`Card\s+\d+: (.*) \| (.*)`)
    match := pattern.FindAllStringSubmatch(line, -1)

    numbers := [2][]int{{}, {}}

    for i := 0; i < 2; i++ {
        for _, n := range strings.Split(match[0][i+1], " ") {
            if n, err := strconv.Atoi(strings.Trim(n, " \t\r\n")); err == nil {
                numbers[i] = append(numbers[i], n)
            }
        }
    }

    return numbers;
}
