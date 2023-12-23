package main

import (
    "fmt"
    "os"
    "log"
    "bufio"
    "regexp"
    "unicode"
    "strconv"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    schematic := []string{}

    for scanner.Scan() {
        line := scanner.Text()
        schematic = append(schematic, line)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    gears := map[[2]int][]int{}
    sum := 0
    for i := 0; i < len(schematic); i++ {
        parts := numberPartsInLine(schematic, i, gears)
        for _, part := range parts {
            sum += part
        }
    }

    sum2 := 0
    for _, value := range gears {
        if len(value) == 2 {
            sum2 += value[0] * value[1]
        }
    }

    fmt.Println("sum:", sum)
    fmt.Println("sum2:", sum2)
}

func numberPartsInLine(schematic []string, row int, gears map[[2]int][]int) []int {
    parts := []int{}
    line := []byte(schematic[row])
    pattern := regexp.MustCompile(`(\d+)`)
    allIndexes := pattern.FindAllSubmatchIndex(line, -1)
    for _, loc := range allIndexes {
        if isPart := isPartNumber(schematic, row, loc, gears); isPart {
            n, _ := strconv.Atoi(schematic[row][loc[0]:loc[1]])
            parts = append(parts, n)
        }
    }
    return parts
}

func isPartNumber(schematic []string, row int, loc []int, gears map[[2]int][]int) bool {
    foundPartNumber := false

    startCol, endCol := loc[0], loc[1]

    startRow := row
    endRow := row

    if startRow - 1 >= 0 {
        startRow -= 1
    }

    if endRow + 1 < len(schematic) {
        endRow += 1
    }

    if startCol - 1 >= 0 {
        startCol -= 1
    }

    if endCol > len(schematic[0]) - 1 {
        endCol -= 1
    }


    for i := startRow; i <= endRow; i++ {
        for j := startCol; j <= endCol; j++ {
            currentChar := schematic[i][j]
            if !unicode.IsDigit(rune(currentChar)) && currentChar != '.' {
                foundPartNumber = true
            }
            if currentChar == '*' {
                n, _ := strconv.Atoi(schematic[row][loc[0]:loc[1]])
                key := [2]int{i, j}
                if value, ok := gears[key]; ok {
                    gears[key] = append(value, n)
                } else {
                    gears[key] = []int{n}
                }
            }
        }
    }

    return foundPartNumber
}


