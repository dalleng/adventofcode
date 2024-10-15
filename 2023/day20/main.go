package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Add inputs to conjunction modules
	for moduleName, destinations := range moduleConfiguration {
		for _, destination := range destinations {
			if inputs, ok := conjunctionState[destination]; ok {
				inputs[moduleName[1:]] = LOW
			}
		}
	}

	highCount := 0
	lowCount := 0
	for i := 0; i < 1000; i++ {
		high, low := countPulses(moduleConfiguration, flipFlopState, conjunctionState)
		highCount += high
		lowCount += low
	}
	fmt.Println("high:", highCount)
	fmt.Println("low:", lowCount)
	fmt.Printf("s1=%d\n", highCount*lowCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func appendToFrontier(frontier []Flow, origin string, destinations []string, pulse Pulse) []Flow {
	for _, destination := range destinations {
		frontier = append(frontier, Flow{origin: origin, destination: destination, pulse: pulse})
	}
	return frontier
}

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
