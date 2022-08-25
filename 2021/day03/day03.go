package day03

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day03/aoc3.txt"

type BinTree struct {
	one   *BinTree
	zero  *BinTree
	value string
}

func (bt *BinTree) size() int {
	if bt == nil {
		return 0
	}

	// Leadnode doesn't have subnodes
	if bt.one == nil && bt.zero == nil {
		return 1
	}
	return bt.one.size() + bt.zero.size()
}

func (bt BinTree) String() string {
	return fmt.Sprintf("(1=%s, 0=%s, value=%s)", bt.one, bt.zero, bt.value)
}

func (bt *BinTree) insert(input string) {
	// Creates structure based on binary input string
	node := bt

	for _, ch := range input {
		switch ch {
		case '1':
			if node.one == nil {
				node.one = &BinTree{}
			}
			node = node.one
		case '0':
			if node.zero == nil {
				node.zero = &BinTree{}
			}
			node = node.zero
		}
	}
	node.value = input
}

func Tasks() {
	data := getInputs()

	cntr := make([]int, len(data[0]))
	for _, line := range data {
		for idx, ch := range line {
			switch ch {
			case '1':
				cntr[idx] += 1
			case '0':
				cntr[idx] -= 1
			}
		}
	}

	var gamma, epsilon uint64
	for idx, val := range cntr {
		if val > 0 {
			gamma = gamma | 1<<(len(cntr)-1-idx)
		} else {
			epsilon = epsilon | 1<<(len(cntr)-1-idx)
		}
	}
	fmt.Printf("Day 03 Task 01: %d\n", gamma*epsilon)

	numbers := []uint{}

	for _, line := range data {
		val, _ := strconv.ParseUint(line, 2, len(line))
		numbers = append(numbers, uint(val))
	}

	a, b := task2_filter(numbers)
	fmt.Printf("Day 03 Task 02: %d\n", a*b)

	oxygen, co2 := task2(data)
	fmt.Printf("Day 03 Task 02: %d\n", oxygen*co2)
}

func filterNumbers(list []uint, pos int, larger bool) uint {
	if len(list) == 1 {
		return list[0]
	}
	mask := uint(1 << pos)
	ones := []uint{}
	zeros := []uint{}

	for _, item := range list {
		if item&mask >= mask {
			ones = append(ones, item)
		} else {
			zeros = append(zeros, item)
		}
	}

	if (len(ones) >= len(zeros) && larger) || (len(ones) < len(zeros) && !larger) {
		return filterNumbers(ones, pos-1, larger)
	}
	return filterNumbers(zeros, pos-1, larger)
}

func task2_filter(numbers []uint) (uint, uint) {
	oxygen := filterNumbers(numbers, 11, true)
	co2 := filterNumbers(numbers, 11, false)

	return oxygen, co2
}

func task2(data []string) (int64, int64) {
	root := BinTree{}
	for _, line := range data {
		root.insert(line)
	}

	// Oxygen generator value
	node := &root
	for node.one != nil || node.zero != nil {
		if node.one.size() >= node.zero.size() {
			node = node.one
		} else {
			node = node.zero
		}
	}
	oxygen, _ := strconv.ParseInt((*node).value, 2, 16)

	node = &root
	for node.one != nil || node.zero != nil {
		if node.zero.size() == 0 {
			node = node.one
		} else if node.zero.size() <= node.one.size() || node.one.size() == 0 {
			node = node.zero
		} else {
			node = node.one
		}
	}
	co2, _ := strconv.ParseInt((*node).value, 2, 16)
	return oxygen, co2
}

func getInputs() []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(strings.TrimSpace(string(bytes)), "\n")
}
