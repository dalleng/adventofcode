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

type SearchNode struct {
	WorkflowName string
	Ranges       map[rune][2]int
}

func (s SearchNode) String() string {
	rest := ""
	for _, key := range [4]rune{'x', 'm', 'a', 's'} {
		value, _ := s.Ranges[key]
		rest += fmt.Sprintf("%c: [%d %d] ", key, value[0], value[1])
	}
	return fmt.Sprintf("WorkflowName: %s %s", s.WorkflowName, rest)
}

func (s SearchNode) CopyRanges() map[rune][2]int {
	copyMap := make(map[rune][2]int)
	for key, value := range s.Ranges {
		copyMap[key] = value
	}
	return copyMap
}

func main() {
	f, err := os.Open("input.txt")
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s1 := 0
	for _, part := range parts {
		if isPartAccepted(workflows, part) {
			s1 += partsSum(part)
		}
	}

	fmt.Printf("s1=%d\n", s1)

	acceptedRanges := findAcceptedRanges(workflows)
	// printAcceptedRanges(acceptedRanges)
	s2 := 0
	for _, ar := range acceptedRanges {
		innersum := 0
		for _, value := range ar {
			if innersum == 0 {
				innersum = value[1] - value[0] + 1
			} else {
				innersum *= value[1] - value[0] + 1
			}
		}
		s2 += innersum
	}
	fmt.Printf("s2=%d\n", s2)
}

func printAcceptedRanges(acceptedRanges []map[rune][2]int) {
	for _, ar := range acceptedRanges {
		s := ""
		for _, key := range [4]rune{'x', 'm', 'a', 's'} {
			value, _ := ar[key]
			s += fmt.Sprintf("%c: [%d %d] ", key, value[0], value[1])
		}
		fmt.Println(s)
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
	WorkflowName := "in"
	for !slices.Contains([]string{"A", "R"}, WorkflowName) {
		rules, _ := workflows[WorkflowName]
		for _, rule := range rules {
			if rule.hasComparison() {
				partNumber, _ := part[rule.Category]
				if rule.ComparisonOperator == '<' {
					if partNumber < rule.ComparisonValue {
						WorkflowName = rule.Target
						break
					}
				} else {
					if partNumber > rule.ComparisonValue {
						WorkflowName = rule.Target
						break
					}
				}
			} else {
				WorkflowName = rule.Target
			}
		}
	}
	return WorkflowName == "A"
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

func getNextNodes(workflows map[string][]Rule, current SearchNode) []SearchNode {
	nextNodes := []SearchNode{}
	var next SearchNode
	for _, rule := range workflows[current.WorkflowName] {
		if rule.hasComparison() {
			currentRange, _ := current.Ranges[rule.Category]
			currentMin, currentMax := currentRange[0], currentRange[1]
			min, max := currentMin, currentMax

			if rule.ComparisonOperator == '<' {
				if currentMin < rule.ComparisonValue {
					if rule.ComparisonValue-1 < currentMax {
						max = rule.ComparisonValue - 1
					}
					nextRanges := current.CopyRanges()
					nextRanges[rule.Category] = [2]int{currentMin, max}
					next = SearchNode{
						WorkflowName: rule.Target,
						Ranges:       nextRanges,
					}
					nextNodes = append(nextNodes, next)

					if rule.ComparisonValue <= currentMax {
						current.Ranges[rule.Category] = [2]int{rule.ComparisonValue, currentMax}
					} else {
						break
					}
				}

			} else if rule.ComparisonOperator == '>' {
				if currentMax > rule.ComparisonValue {
					if rule.ComparisonValue+1 > currentMin {
						min = rule.ComparisonValue + 1
					}
					nextRanges := current.CopyRanges()
					nextRanges[rule.Category] = [2]int{min, currentMax}
					next = SearchNode{
						WorkflowName: rule.Target,
						Ranges:       nextRanges,
					}

					nextNodes = append(nextNodes, next)
					if rule.ComparisonValue >= currentMin {
						current.Ranges[rule.Category] = [2]int{currentMin, rule.ComparisonValue}
					} else {
						break
					}
				}
			}
		} else {
			next = SearchNode{
				WorkflowName: rule.Target,
				Ranges:       current.Ranges,
			}
			nextNodes = append(nextNodes, next)
		}
	}
	return nextNodes
}

func findAcceptedRanges(workflows map[string][]Rule) []map[rune][2]int {
	acceptedRanges := []map[rune][2]int{}
	initialNode := SearchNode{
		WorkflowName: "in",
		Ranges: map[rune][2]int{
			'x': {1, 4000},
			'm': {1, 4000},
			'a': {1, 4000},
			's': {1, 4000},
		},
	}
	frontier := []SearchNode{initialNode}
	for len(frontier) > 0 {
		current := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]
		for _, node := range getNextNodes(workflows, current) {
			if node.WorkflowName == "A" {
				acceptedRanges = append(acceptedRanges, node.Ranges)
			} else if node.WorkflowName != "R" {
				frontier = append(frontier, node)
			}
		}
	}
	return acceptedRanges
}
