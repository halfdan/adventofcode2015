package day08

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const filename = "day08/aoc8.txt"

func Tasks() {
	entries := getInputs()

	sum := 0
	for _, entry := range entries {
		for _, output := range entry.outputs {
			switch len(output) {
			case 2, 4, 3, 7:
				sum += 1
			}
		}
	}

	fmt.Printf("Day 08 Task 01: %v\n", sum)
}

type entry struct {
	signals []string
	outputs []string
}

func getInputs() []entry {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	entries := make([]entry, 0, len(lines))

	for _, line := range lines[:len(lines)-1] {
		elements := strings.Split(line, " | ")
		entries = append(entries, entry{
			signals: strings.Fields(elements[0]),
			outputs: strings.Fields(elements[1]),
		})
	}

	return entries
}
