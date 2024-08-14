package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	workflows := map[string][]string{}
	parts := []map[rune]int{}
	scanner := bufio.NewScanner(f)
	workflowSection := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			workflowSection = false
			continue
		}
		if workflowSection {
			name, rules := parseWorkflow(line)
			workflows[name] = rules
		} else {
			parts = append(parts, parsePart(line))
		}
	}

	s1 := 0
	for _, part := range parts {
		if isPartAccepted(workflows, part) {
			s1 += partsSum(part)
		}
	}

	fmt.Printf("s1=%d\n", s1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partsSum(part map[rune]int) int {
	sum := 0
	for _, value := range part {
		sum += value
	}
	return sum
}

func isPartAccepted(workflows map[string][]string, part map[rune]int) bool {
	fmt.Println("part: ", part)
	currentWorkflow := "in"
	re := regexp.MustCompile(`([xmas])([><])(\d+):(.*)`)
	for !slices.Contains([]string{"A", "R"}, currentWorkflow) {
		fmt.Println(currentWorkflow)
		rules, _ := workflows[currentWorkflow]
		for _, rule := range rules {
			matches := re.FindStringSubmatch(rule)
			if matches != nil {
				category, comparison, numberStr, nextWorkflow := matches[1], matches[2], matches[3], matches[4]
				ruleNumber, _ := strconv.Atoi(numberStr)
				partNumber, _ := part[[]rune(category)[0]]
				if comparison == "<" {
					if partNumber < ruleNumber {
						currentWorkflow = nextWorkflow
						break
					}
				} else {
					if partNumber > ruleNumber {
						currentWorkflow = nextWorkflow
						break
					}
				}
			} else {
				currentWorkflow = rule
			}
		}
	}
	return currentWorkflow == "A"
}

func parseWorkflow(workflow string) (string, []string) {
	endOfNameIdx := strings.IndexRune(workflow, '{')
	endOfRuleSection := strings.IndexRune(workflow, '}')
	name := workflow[:endOfNameIdx]
	rules := []string{}
	for _, rule := range strings.Split(workflow[endOfNameIdx+1:endOfRuleSection], ",") {
		rules = append(rules, rule)
	}
	return name, rules
}

func parsePart(part string) map[rune]int {
	partMap := map[rune]int{}
	startIdx := strings.IndexRune(part, '{') + 1
	endIdx := strings.IndexRune(part, '}')
	for _, categoryAndValueStr := range strings.Split(part[startIdx:endIdx], ",") {
		categoryAndValue := strings.Split(categoryAndValueStr, "=")
		category, value := categoryAndValue[0], categoryAndValue[1]
		if valueInt, err := strconv.Atoi(value); err == nil {
			partMap[[]rune(category)[0]] = valueInt
		}
	}
	return partMap
}
