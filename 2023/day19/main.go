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

type Rule struct {
	Category           rune
	ComparisonOperator rune
	ComparisonValue    int
	Target             string
}

func (r Rule) hasComparison() bool {
	return r.Category != 0 && r.ComparisonOperator != 0 && r.ComparisonValue != 0
}

func (r Rule) String() string {
	if !r.hasComparison() {
		return r.Target
	}
	return fmt.Sprintf("%c%c%d:%s", r.Category, r.ComparisonOperator, r.ComparisonValue, r.Target)
}

func main() {
	f, err := os.Open("input.txt")
	// f, err := os.Open("example.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	workflows := map[string][]Rule{}
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

	// fmt.Println(workflows)
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

func isPartAccepted(workflows map[string][]Rule, part map[rune]int) bool {
	fmt.Println("part: ", part)
	currentWorkflow := "in"
	for !slices.Contains([]string{"A", "R"}, currentWorkflow) {
		fmt.Println("currentWorkflow: ", currentWorkflow)
		rules, _ := workflows[currentWorkflow]
		for _, rule := range rules {
			if rule.hasComparison() {
				partNumber, _ := part[rule.Category]
				if rule.ComparisonOperator == '<' {
					if partNumber < rule.ComparisonValue {
						currentWorkflow = rule.Target
						break
					}
				} else {
					if partNumber > rule.ComparisonValue {
						currentWorkflow = rule.Target
						break
					}
				}
			} else {
				currentWorkflow = rule.Target
			}
		}
	}
	return currentWorkflow == "A"
}

func parseRule(rawRule string) Rule {
	re := regexp.MustCompile(`([xmas])([><])(\d+):(.*)`)
	matches := re.FindStringSubmatch(rawRule)
	if matches != nil {
		category, comparison, numberStr, nextWorkflow := matches[1], matches[2], matches[3], matches[4]
		value, _ := strconv.Atoi(numberStr)
		return Rule{
			Category:           rune(category[0]),
			ComparisonOperator: rune(comparison[0]),
			ComparisonValue:    value,
			Target:             nextWorkflow,
		}
	}
	return Rule{
		Target: rawRule,
	}
}

func parseWorkflow(workflow string) (string, []Rule) {
	endOfNameIdx := strings.IndexRune(workflow, '{')
	endOfRuleSection := strings.IndexRune(workflow, '}')
	name := workflow[:endOfNameIdx]
	rules := []Rule{}
	for _, rule := range strings.Split(workflow[endOfNameIdx+1:endOfRuleSection], ",") {
		rules = append(rules, parseRule(rule))
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
