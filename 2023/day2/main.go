package main

import (
    "os"
    "log"
    "bufio"
    "regexp"
    "strings"
    "strconv"
)


func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    sum := 0
    sumPower := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        gid, game := parseGame(line)
        if isGamePossible(game) {
            sum += gid
        }
        minSet := getMinSet(game)
        sumPower += minSet["red"] * minSet["green"] * minSet["blue"]
    }

    if err := scanner.Err(); err != nil {
        log.Fatal("Invalid input: %s", err)
    }

    log.Printf("result1=%d", sum)
    log.Printf("result2=%d", sumPower)
}

func parseGame(line string) (int, []map[string]int) {
    game := []map[string]int{}
    
    pattern := regexp.MustCompile(`Game (\d+): (.*)`)
    match := pattern.FindStringSubmatch(line)

    gid, _ := strconv.Atoi(match[1])
    rest := match[2]

    for _, gameset := range strings.Split(rest, ";") {
        gameset = strings.TrimSpace(gameset)
        d := map[string]int{}

        for _, cube := range strings.Split(gameset, ",") {
            cube = strings.TrimSpace(cube)
            num_and_color := strings.Split(cube, " ")
            num, _ := strconv.Atoi(num_and_color[0])
            color := num_and_color[1]
            d[color] = num
        }
        game = append(game, d)
    }

    return gid, game
}

func isGamePossible(game []map[string]int) bool {
    maxCubes := map[string]int{
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    for _, cubes := range game {
        for color, num := range cubes {
            if maxCubes[color] < num {
                return false
            }
        }
    }
    return true
}

func getMinSet(game []map[string]int) map[string]int {
    maxCubes := map[string]int{
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    for _, gameset := range game {
        for color, count := range gameset {
            if count > maxCubes[color] {
                maxCubes[color] = count
            }
        }
    }
    return maxCubes
}
