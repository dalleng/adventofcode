package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Box struct {
	// cache to access elements in the list faster than with linear search
	cache map[string]*list.Element
	// holds the values in the inserted order
	lenses *list.List
}

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := 0
	boxes := map[int]Box{}

	for _, value := range strings.Split(line, ",") {
		s1 += computeHash(value)

		if lastIdx := len(value) - 1; value[lastIdx] == '-' {
			label := value[:lastIdx]
			boxNumber := computeHash(label)
			if box, ok := boxes[boxNumber]; ok {
				if e, ok := box.cache[label]; ok {
					box.lenses.Remove(e)
					delete(box.cache, label)
				}
			}
		} else {
			labelAndNumber := strings.Split(value, "=")
			label := labelAndNumber[0]
			boxNumber := computeHash(label)
			box, boxExists := boxes[boxNumber]
			if !boxExists {
				box = Box{
					cache:  make(map[string]*list.Element),
					lenses: list.New(),
				}
				boxes[boxNumber] = box
			}
			if e, ok := box.cache[label]; !ok {
				box.cache[label] = box.lenses.PushBack(value)
			} else {
				e.Value = value
			}
		}
	}

	s2 := 0
	for boxNumber, box := range boxes {
		elementIndex := 0
		for e := box.lenses.Front(); e != nil; e = e.Next() {
			labelAndNumber := strings.Split(e.Value.(string), "=")
			number, _ := strconv.Atoi(labelAndNumber[1])
			s2 += (boxNumber + 1) * (elementIndex + 1) * number
			elementIndex++
		}
	}

	fmt.Printf("s1=%d\n", s1)
	fmt.Printf("s2=%d\n", s2)
}

func computeHash(value string) int {
	hash := 0
	for _, asciiValue := range []byte(value) {
		hash += int(asciiValue)
		hash *= 17
		hash = hash % 256
	}
	return hash
}
