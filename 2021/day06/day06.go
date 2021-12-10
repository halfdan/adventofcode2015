package day06

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day06/aoc6.txt"

func fishSum(fish map[int]int) (sum int) {
	for k := range fish {
		sum += fish[k]
	}
	return sum
}

func Tasks() {
	fish := getInputs()

	for i := 0; i < 80; i++ {
		tickDay(i, &fish)
	}
	fmt.Printf("Day 06 Task 01: %d\n", fishSum(fish))

	for i := 80; i < 256; i++ {
		tickDay(i, &fish)
	}

	fmt.Printf("Day 06 Task 02: %d\n", fishSum(fish))
}

func tickDay(day int, fish *map[int]int) {
	(*fish)[(day+7)%9] += (*fish)[day%9]
}

func getInputs() map[int]int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	digits := strings.Split(string(bytes), ",")

	nums := map[int]int{}
	for i := range digits {
		val, err := strconv.Atoi(strings.TrimSpace(digits[i]))
		if err == nil {
			nums[val] += 1
		}
	}

	return nums
}
