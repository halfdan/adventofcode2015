package day10

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename = "day10/aoc10.txt"

var pairs = map[int32]int32{
	0x28: 0x29, // (:)
	0x5b: 0x5d, // [:]
	0x7b: 0x7d, // {:}
	0x3c: 0x3e, // <:>
}

var closingScore = map[int32]int{
	0x29: 1,
	0x5d: 2,
	0x7d: 3,
	0x3e: 4,
}

var invalidScore = map[rune]int{
	0x29: 3,     // )
	0x5d: 57,    // ]
	0x7d: 1197,  //
	0x3e: 25137, // >
}

func getInputs() []string {
	bytes, _ := ioutil.ReadFile(filename)

	return strings.Split(strings.TrimRight(string(bytes), "\n"), "\n")
}

func scoreIncomplete(lines []string) int {
	score := 0
	stack := make([]int32, 0, 30)
	scores := []int{}

LineLoop:
	for _, line := range lines {
		stack = stack[:0] // Trim stack
		score = 0
		for _, ch := range line {
			switch ch {
			case '{', '(', '[', '<':
				closing := pairs[ch]
				stack = append(stack, closing)
			case '}', ')', ']', '>':
				if stack[len(stack)-1] != ch {
					continue LineLoop
				}
				stack = stack[:len(stack)-1] // Pop element
			}
		}
		for l := len(stack) - 1; l >= 0; l-- {
			score = score*5 + closingScore[stack[l]]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func totalScoreCorrupted(lines []string) int {
	score := 0
	stack := make([]int32, 0, 30)
	for _, line := range lines {
		stack = stack[:0] // Trim stack
	Loop:
		for _, ch := range line {
			switch ch {
			case '{', '(', '[', '<':
				stack = append(stack, pairs[ch])
			case '}', ')', ']', '>':
				if stack[len(stack)-1] != ch {
					score += invalidScore[ch]
					break Loop
				}
				stack = stack[:len(stack)-1] // Pop element
			}
		}
	}
	return score
}

func Tasks() {
	lines := getInputs()

	fmt.Printf("Day 10 Task 01: %d\n", totalScoreCorrupted(lines))
	fmt.Printf("Day 10 Task 02: %d\n", scoreIncomplete(lines))
}
