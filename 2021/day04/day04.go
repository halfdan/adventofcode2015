package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/halfdan/adventofcode2021/aocutil"
)

const filename = "day04/aoc4.txt"

func Tasks() {
	data := aocutil.GetInputs(filename)

	numbers, boards := parseInput(data)
	var firstScore, lastScore int

	for i := range numbers {
		for b := range boards {
			if boards[b].hasBingo {
				continue
			}
			boards[b].markDigit(numbers[i])

			if boards[b].isBingo() {
				if firstScore == 0 {
					firstScore = boards[b].sumUnmarked() * numbers[i]
				}
				lastScore = boards[b].sumUnmarked() * numbers[i]
				boards[b].hasBingo = true
			}
		}
	}

	fmt.Printf("Day 04 Task 01: %d\n", firstScore)
	fmt.Printf("Day 04 Task 02: %d\n", lastScore)

}

type Board struct {
	numbers  []int
	matched  []bool
	hasBingo bool
}

func (b *Board) sumUnmarked() (sum int) {
	for i := range b.numbers {
		if !b.matched[i] {
			sum += b.numbers[i]
		}
	}
	return
}

func (b *Board) isBingo() bool {
	// Check rows
	isBingo := true
	for i := 0; i < len(b.matched); i += 5 {
		isBingo = true
		for j := 0; j < 5; j++ {
			isBingo = isBingo && b.matched[i+j]
		}
		if isBingo {
			return true
		}
	}

	// Check columns
	for i := 0; i < 5; i++ {
		isBingo = true
		for j := 0; j < 5; j++ {
			isBingo = isBingo && b.matched[i+j*5]
		}
		if isBingo {
			return true
		}
	}
	return false
}

func (b *Board) markDigit(digit int) {
	for idx := 0; idx < len(b.numbers); idx++ {
		if b.numbers[idx] == digit {
			b.matched[idx] = true
			break
		}
	}
}

func (b Board) String() string {
	result := ""
	for i := range b.numbers {
		if b.matched[i] {
			result += fmt.Sprintf("[% d]", b.numbers[i])
		} else {
			result += fmt.Sprintf(" % d ", b.numbers[i])
		}
		if i+1%5 == 0 {
			result += "\n"
		}
	}
	return result
}

// Returns the numbers to draw as well as all boards
func parseInput(input []string) ([]int, []Board) {
	numbers := strings.Split(input[0], ",")
	nums := make([]int, len(numbers))
	for idx, val := range numbers {
		nums[idx], _ = strconv.Atoi(val)
	}

	boards := make([]Board, 0)

	for idx := 2; idx < len(input); idx += 6 {
		boards = append(boards, newBoard(input[idx:idx+5]))
	}

	return nums, boards
}

func newBoard(input []string) Board {
	if len(input) != 5 {
		panic("A board cannot have dimensions bigger than 5x5")
	}

	numbers := make([]int, 0, 25)
	for _, line := range input {
		digits := strings.Fields(line)
		for _, digit := range digits {
			val, _ := strconv.Atoi(digit)
			numbers = append(numbers, val)
		}
	}
	return Board{numbers: numbers, matched: make([]bool, len(numbers))}
}
