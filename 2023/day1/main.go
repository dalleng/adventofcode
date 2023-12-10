package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
    "unicode"
    "strconv"
)


func main() {
    f, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }
    
    defer f.Close()

    var sum1, sum2 int64;
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        currentLine := scanner.Text()
        //fmt.Println(currentLine)
        currentNum := extractNumber(currentLine)
        currentNum2 := extractNumber2(currentLine)
        sum1 += currentNum
        sum2 += currentNum2
    }
    fmt.Printf("s1: %d\n", sum1)
    fmt.Printf("s2: %d\n", sum2)

    if err := scanner.Err(); err != nil {
        log.Fatal("Invalid input: %s", err)
    }
}


func extractNumber(s string) int64 {
    var first byte;
    var last byte;

    i := 0
    j := len(s) - 1

    for i <= j {
        if first == 0 {
            if unicode.IsDigit(rune(s[i])) {
                first = s[i]
            } else {
                i++
            }
        }

        if last == 0 {
            if unicode.IsDigit(rune(s[j])) {
                last = s[j]
            } else {
                j--
            }
        }

        if first != 0 && last != 0 {
            break
        }
    }


    currentNum, _ := strconv.ParseInt(fmt.Sprintf("%c%c", first, last), 10, 8)
    return currentNum
}

func extractNumber2(s string) int64 {
    digits := []int64{}
    numberLookup := map[string]int64{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    start := 0
    end := 0
    slen := len(s)
    for end < slen {
        if unicode.IsDigit(rune(s[start])) {
            digit, _ := strconv.ParseInt(string(s[start]), 10, 8)
            digits = append(digits, digit)
            start += 1
            end = start
            continue
        }
        length := end - start + 1
        potential := s[start:end+1]
        if val, ok := numberLookup[potential]; ok {
            digits = append(digits, val)
            start = end
            continue
        }
        if length == 5 || end == slen - 1 {
            start += 1
            end = start
            continue
        }
        end += 1
    }
    returnValue, _ := strconv.ParseInt(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]), 10, 8)
    return returnValue
}
