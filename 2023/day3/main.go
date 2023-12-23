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

    sum := 0
    for i := 0; i < len(schematic); i++ {
        parts := numberPartsInLine(schematic, i)
        for _, part := range parts {
            sum += part
        }
    }
    fmt.Printf("sum: %d", sum)
}

func numberPartsInLine(schematic []string, row int) []int {
    parts := []int{}
    line := []byte(schematic[row])
    pattern := regexp.MustCompile(`(\d+)`)
    allIndexes := pattern.FindAllSubmatchIndex(line, -1)
    for _, loc := range allIndexes {
        if isPart := isPartNumber(schematic, row, loc); isPart {
            n, _ := strconv.Atoi(schematic[row][loc[0]:loc[1]])
            parts = append(parts, n)
        }
    }
    fmt.Println(parts)
    return parts
}

func isPartNumber(schematic []string, row int, loc []int) bool {
    fmt.Println(schematic[row])

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

    fmt.Printf("startRow: %d endRow: %d startCol: %d endCol: %d\n", startRow, endRow, startCol, endCol)

    for i := startRow; i <= endRow; i++ {
        for j := startCol; j <= endCol; j++ {
            if currentChar := schematic[i][j]; !unicode.IsDigit(rune(currentChar)) && currentChar != '.' {
                return true
            }
        }
    }

    return false
}


