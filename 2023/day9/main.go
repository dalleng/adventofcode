package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
  sequences := readSequencesFromFile("input.txt")
  s1 := 0
  s2 := 0
  for _, seq := range sequences {
    s1 += findNextElement(seq)
    s2 += findPreviousElement(seq)
  }
  fmt.Println("s1:", s1)
  fmt.Println("s2:", s2)
}

func readSequencesFromFile(filename string) [][]int {
  f, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  var sequences [][]int
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    line := scanner.Text()
    numbers := make([]int, 0)
    for _, value := range strings.Split(line, " ") {
      if n, err := strconv.Atoi(value); err != nil {
        log.Fatal(err)
      } else {
        numbers = append(numbers, n)
      }
    }
    sequences = append(sequences, numbers)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return sequences
}

func findNextElement(sequence []int) int {
  areAllEqual := areAllElementsEqual(sequence)

  if areAllEqual {
    return sequence[0]
  }

  var lastElement int
  lastElements := []int{sequence[len(sequence)-1]}
  currentSequence := sequence

  for !areAllEqual {
    currentSequence, lastElement = getDiffSequence(currentSequence)
    lastElements = append(lastElements, lastElement)
    areAllEqual = areAllElementsEqual(currentSequence)
  }

  lenLastElements := len(lastElements)
  last := lastElements[lenLastElements-1]
  for i := lenLastElements - 2; i >= 0; i-- {
    last += lastElements[i]
  }

  return last
}


func findPreviousElement(sequence []int) int {
  areAllEqual := areAllElementsEqual(sequence)

  if areAllEqual {
    return sequence[0]
  }

  firstElements := []int{sequence[0]}
  currentSequence := sequence

  for !areAllEqual {
    currentSequence, _ = getDiffSequence(currentSequence)
    firstElements = append(firstElements, currentSequence[0])
    areAllEqual = areAllElementsEqual(currentSequence)
  }

  lenFirstElements := len(firstElements)
  first := firstElements[lenFirstElements-1]
  for i := lenFirstElements - 2; i >= 0; i-- {
    first = firstElements[i] - first
  }

  return first
}

func areAllElementsEqual(sequence []int) bool {
  for i := 0; i < len(sequence) - 1; i++ {
    if sequence[i] != sequence[i+1] {
      return false
    }
  }
  return true
}

func getDiffSequence(sequence []int) ([]int, int) {
  var diffSeq []int
  var diff int
  for i := 0; i < len(sequence) - 1; i++ {
    diff = sequence[i+1] - sequence[i]
    diffSeq = append(diffSeq, diff)
  }
  return diffSeq, diff
}
