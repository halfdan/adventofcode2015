package day01

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const filename = "day01/aoc1.txt"

func Tasks() {
	measurements := loadMeasurements()

	var count int
	var last int

	last = measurements[0]
	for _, num := range measurements {
		if num > last {
			count++
		}
		last = num
	}
	fmt.Printf("Day 01 Task 01: %d\n", count)

	count = 0
	prev := math.MaxInt
	for i := 0; i < len(measurements)-2; i++ {
		sum := measurements[i] + measurements[i+1] + measurements[i+2]
		if sum > prev {
			count++
		}
		prev = sum
	}
	fmt.Printf("Day 02 Task 02: %d\n", count)
}

func loadMeasurements() []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	nums := []int{}
	for _, val := range lines[:len(lines)-1] {
		num, _ := strconv.Atoi(val)
		nums = append(nums, num)
	}

	return nums
}
