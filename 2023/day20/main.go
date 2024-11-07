package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Pulse int

const (
	LOW Pulse = iota
	HIGH
)

type Flow struct {
	origin      string
	destination string
	pulse       Pulse
}

func main() {
	flipFlopState := map[string]bool{}
	conjunctionState := map[string]map[string]Pulse{}
	moduleConfiguration := map[string][]string{}

	f, err := os.Open("input.txt")
	// f, err := os.Open("example2.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		originAndDestinations := strings.Split(line, " -> ")
		origin := originAndDestinations[0]

		if origin[0] == '%' {
			if _, ok := flipFlopState[origin[1:]]; !ok {
				flipFlopState[origin[1:]] = false
			}
		} else if origin[0] == '&' {
			if _, ok := conjunctionState[origin[1:]]; !ok {
				conjunctionState[origin[1:]] = map[string]Pulse{}
			}
		}

		destinations := []string{}
		for _, destination := range strings.Split(originAndDestinations[1], ",") {
			destinations = append(destinations, strings.Trim(destination, " "))
		}
		moduleConfiguration[origin] = destinations
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Add inputs to conjunction modules
	for moduleName, destinations := range moduleConfiguration {
		for _, destination := range destinations {
			if inputs, ok := conjunctionState[destination]; ok {
				inputs[moduleName[1:]] = LOW
			}
		}
	}

	// highCount := 0
	// lowCount := 0
	// for i := 0; i < 1000; i++ {
	// 	high, low := countPulses(moduleConfiguration, flipFlopState, conjunctionState)
	// 	highCount += high
	// 	lowCount += low
	// }
	// fmt.Println("high:", highCount)
	// fmt.Println("low:", lowCount)
	// fmt.Printf("s1=%d\n", highCount*lowCount)

	// Solution part 2

	// Step 1. Find the module that inputs into rx
	var conjunctionToRx string
	for module, destinations := range moduleConfiguration {
		// fmt.Println(module)
		if slices.Contains(destinations, "rx") {
			// We assume rx has only one input and it is a conjunction
			if _, ok := conjunctionState[module[1:]]; ok {
				conjunctionToRx = module[1:]
				break
			}
		}
	}
	fmt.Println(conjunctionToRx)

	// Step 2.
	// Initialize maps to track if a high pulse has been seen for each module
	// that is an input to `conjunctionToRx`` and to count the button presses
	// needed for each module to output a high pulse.
	seenHighPulse := map[string]bool{}
	buttonPresses := map[string]int{}
	if conjunctionToRx != "" {
		// find all the modules that are inputs to rx
		if inputs, ok := conjunctionState[conjunctionToRx]; ok {
			for input := range inputs {
				seenHighPulse[input] = false
				buttonPresses[input] = 0
			}
		}
	}
	fmt.Println(seenHighPulse)

	// THIS DOES NOT WORK, LOOK INTO THE VALUES OF conjunctionState inside countPulses2
	for {
		for module, seen := range seenHighPulse {
			if !seen {
				buttonPresses[module] += 1
			}
		}
		countPulses2(moduleConfiguration, flipFlopState, conjunctionState, seenHighPulse)
		allHigh := true
		for _, seen := range seenHighPulse {
			if !seen {
				allHigh = false
				break
			}
		}
		if allHigh {
			break
		}
	}

	counts := []int{}
	for _, count := range buttonPresses {
		counts = append(counts, count)
	}
	fmt.Println(lcm(counts[0], counts[1], counts[2:]...))
}

func gcd(a, b int) int {
	for b > 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a int, b int, integers ...int) int {
	lcm := a * b / gcd(a, b)
	for _, i := range integers {
		lcm = lcm * i / gcd(lcm, i)
	}
	return lcm
}

func appendToFrontier(frontier []Flow, origin string, destinations []string, pulse Pulse) []Flow {
	for _, destination := range destinations {
		frontier = append(frontier, Flow{origin: origin, destination: destination, pulse: pulse})
	}
	return frontier
}

// Function used for the solution of the 1st part
func countPulses(moduleConfiguration map[string][]string, flipFlopState map[string]bool, conjunctionState map[string]map[string]Pulse) (int, int) {
	frontier := []Flow{{origin: "button", destination: "broadcaster", pulse: LOW}}
	highCounter := 0
	lowCounter := 0

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if current.pulse == LOW {
			lowCounter++
		} else {
			highCounter++
		}

		if destinations, ok := moduleConfiguration[current.destination]; ok {
			frontier = appendToFrontier(frontier, current.destination, destinations, current.pulse)
		} else if destinations, ok := moduleConfiguration["%"+current.destination]; ok {
			isIgnored, nextState, output := flipFlop(current.pulse, flipFlopState[current.destination])
			if isIgnored {
				continue
			}
			flipFlopState[current.destination] = nextState
			frontier = appendToFrontier(frontier, current.destination, destinations, output)
		} else if destinations, ok := moduleConfiguration["&"+current.destination]; ok {
			conjunctionState[current.destination][current.origin] = current.pulse
			output := conjunction(conjunctionState[current.destination])
			frontier = appendToFrontier(frontier, current.destination, destinations, output)
		}
	}

	return highCounter, lowCounter
}

func countPulses2(moduleConfiguration map[string][]string, flipFlopState map[string]bool, conjunctionState map[string]map[string]Pulse, seenHighPulse map[string]bool) bool {
	frontier := []Flow{{origin: "button", destination: "broadcaster", pulse: LOW}}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if seen, ok := seenHighPulse[current.origin]; ok {
			if !seen && current.pulse == HIGH {
				seenHighPulse[current.origin] = true
			}
		}

		if destinations, ok := moduleConfiguration[current.destination]; ok {
			frontier = appendToFrontier(frontier, current.destination, destinations, current.pulse)
		} else if destinations, ok := moduleConfiguration["%"+current.destination]; ok {
			isIgnored, nextState, output := flipFlop(current.pulse, flipFlopState[current.destination])
			if isIgnored {
				continue
			}
			flipFlopState[current.destination] = nextState
			frontier = appendToFrontier(frontier, current.destination, destinations, output)
		} else if destinations, ok := moduleConfiguration["&"+current.destination]; ok {
			conjunctionState[current.destination][current.origin] = current.pulse
			output := conjunction(conjunctionState[current.destination])
			frontier = appendToFrontier(frontier, current.destination, destinations, output)
		}
	}

	return false
}

func flipFlop(input Pulse, state bool) (isIgnored bool, nextState bool, output Pulse) {
	/*
		Flip-flop
			- can be on/off
			- initially off
			- if receives high pulse, ignores
			- if receives low pulse, toggles on/off
				-if it was off, turns on and sends high
				- if it was on, turns off and sends low
	*/
	nextState = state
	isIgnored = true
	output = input

	if input == LOW {
		isIgnored = false
		nextState = !state
		if !state {
			output = HIGH
		}
	}

	return
}

func conjunction(state map[string]Pulse) Pulse {
	/*
		Conjunction
			- remebers the type of the most recent pulse from every input
			- initially low for each input
			- when a pulse is received, first updates memory.
				- if it remembers high for all inputs, sends low
				- else, sends high pulse
	*/
	allHigh := true
	for _, memory := range state {
		if memory == LOW {
			allHigh = false
			break
		}
	}
	if allHigh {
		return LOW
	}
	return HIGH
}
