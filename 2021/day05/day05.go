package day05

import (
	"fmt"

	"github.com/halfdan/adventofcode2021/aocutil"
)

const filename = "day05/aoc5.txt"

type point struct {
	x, y int
}

type line struct {
	p1, p2 point
}

func (l *line) isDiagonal() bool {
	return l.p1.x != l.p2.x && l.p1.y != l.p2.y
}

func countIntersections(lines []line, includeDiagonal bool) int {
	points := map[point]int{}

	for _, l := range lines {
		addx := 0
		addy := 0
		if l.p1.x > l.p2.x {
			addx = -1
		}
		if l.p1.x < l.p2.x {
			addx = 1
		}
		if l.p1.y > l.p2.y {
			addy = -1
		}
		if l.p1.y < l.p2.y {
			addy = 1
		}

		if !includeDiagonal && l.isDiagonal() {
			continue
		}

		startX, startY, targetX, targetY := l.p1.x, l.p1.y, l.p2.x, l.p2.y
		for startX != targetX || startY != targetY {
			points[point{x: startX, y: startY}] += 1

			startX += addx
			startY += addy
		}
		points[point{x: startX, y: startY}]++
	}

	overlaps := 0
	for _, v := range points {
		if v > 1 {
			overlaps++
		}
	}
	return overlaps
}

func Tasks() {
	data := aocutil.GetInputs(filename)

	var lines []line
	for _, l := range data {
		var (
			x1, x2 int
			y1, y2 int
		)
		fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		lines = append(lines, line{
			p1: point{x: x1, y: y1},
			p2: point{x: x2, y: y2},
		})
	}

	fmt.Printf("Day 05 Task 01: %d\n", countIntersections(lines, false))
	fmt.Printf("Day 05 Task 02: %d\n", countIntersections(lines, true))
}
