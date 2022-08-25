package day07

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const filename = "day07/aoc7.txt"

type IntSlice []int

func abs(k int) int {
	if k < 0 {
		return -k
	} else {
		return k
	}
}

func (vals *IntSlice) minMax() (int, int) {
	var max int = (*vals)[0]
	var min int = (*vals)[0]
	for _, value := range *vals {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func (is *IntSlice) fuelSpend(pos int) (sum int) {
	for _, v := range *is {
		sum += abs(pos - v)
	}
	return sum
}

func (is *IntSlice) fuelSpendIncreasing(pos int) (sum int) {
	for _, v := range *is {
		distance := abs(pos - v)
		sum += (distance * (distance + 1)) / 2
	}
	return sum
}

func Tasks() {
	crabs := getInputs()
	min, max := crabs.minMax()
	minFuelSpend := crabs.fuelSpend(min)
	for i := min; i <= max; i++ {
		fuelSpend := crabs.fuelSpend(i)
		if fuelSpend < minFuelSpend {
			minFuelSpend = fuelSpend
		}
	}
	fmt.Printf("Day 07 Task 01: %d\n", minFuelSpend)

	minFuelSpend = crabs.fuelSpendIncreasing(min)
	for i := min; i <= max; i++ {
		fuelSpend := crabs.fuelSpendIncreasing(i)
		if fuelSpend < minFuelSpend {
			minFuelSpend = fuelSpend
		}
	}
	fmt.Printf("Day 07 Task 02: %d\n", minFuelSpend)
}

func getInputs() IntSlice {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	digits := strings.Split(string(bytes), ",")

	nums := IntSlice{}
	for i := range digits {
		val, err := strconv.Atoi(strings.TrimSpace(digits[i]))
		if err == nil {
			nums = append(nums, val)
		}
	}

	return nums
}
