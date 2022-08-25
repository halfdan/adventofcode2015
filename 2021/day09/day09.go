package day09

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const filename = "day09/aoc9.txt"

func Tasks() {
	m := getInputs()

	points := m.findLowPoints()
	riskLevel := computeRiskLevel(&m, points)
	fmt.Printf("Day 09 Task 01: %d\n", riskLevel)

	basins := []int{}
	for _, p := range points {
		basins = append(basins, m.basinSize(p))
	}

	sort.Ints(basins)

	fmt.Printf("Day 09 Task 02: %d\n", basins[len(basins)-1]*basins[len(basins)-2]*basins[len(basins)-3])
}

type point struct{ x, y int }

func (m *matrix) findLowPoints() []point {
	points := []point{}
	for i := range *m {
		for j := range (*m)[i] {
			if (*m).isLowPoint(i, j) {
				points = append(points, point{i, j})
			}
		}
	}
	return points
}

func computeRiskLevel(m *matrix, points []point) int {
	var riskSum int

	for _, p := range points {
		riskSum += (*m).valueAt(p.x, p.y)
	}

	return riskSum
}

type matrix [][]int

func (m *matrix) basinSize(p point) int {
	checked := map[point]bool{}

	var f func(point) int

	f = func(p point) int {
		if _, ok := checked[p]; ok {
			return 0
		}

		checked[p] = true
		currentValue := (*m).valueAt(p.x, p.y)

		if currentValue == 9 || currentValue == math.MaxInt {
			return 0
		}

		return 1 + f(point{p.x - 1, p.y}) +
			f(point{p.x, p.y - 1}) +
			f(point{p.x + 1, p.y}) +
			f(point{p.x, p.y + 1})
	}

	return f(p)
}

func (m *matrix) isLowPoint(x, y int) bool {
	val := (*m)[x][y]

	return (val < (*m).valueAt(x-1, y) &&
		val < (*m).valueAt(x, y-1) &&
		val < (*m).valueAt(x+1, y) &&
		val < (*m).valueAt(x, y+1))
}

func (m *matrix) valueAt(x, y int) int {
	if x < 0 || x > len(*m)-1 {
		return math.MaxInt
	}

	if y < 0 || y > len((*m)[x])-1 {
		return math.MaxInt
	}

	return (*m)[x][y]
}

func getInputs() matrix {
	bytes, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(bytes), "\n")

	m := matrix{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		row := []int{}
		for _, digit := range line {
			val, _ := strconv.Atoi(string(digit))
			row = append(row, val)
		}

		m = append(m, row)
	}

	return m
}
