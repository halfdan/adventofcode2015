package day15

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const filename = "day15/aoc15.txt"

var digitMap = map[uint8]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

type coord struct {
	x, y int
}

func findPath(c1, c2 coord, edgeWeights map[coord]int) (int, []coord) {
	h := newHeap()
	visited := make(map[coord]bool)
	h.push(path{value: 0, nodes: []coord{c1}})

	for len(*h.values) > 0 {
		// Find the nearest yet to visit nodes
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]
		if visited[node] {
			continue
		}

		if node == c2 {
			return p.value, p.nodes
		}

		edges := []coord{
			{x: node.x - 1, y: node.y},
			{x: node.x, y: node.y - 1},
			{x: node.x + 1, y: node.y},
			{x: node.x, y: node.y + 1},
		}

		for _, edge := range edges {
			if _, exists := edgeWeights[edge]; exists {
				if !visited[edge] {
					h.push(path{value: p.value + edgeWeights[edge], nodes: append([]coord{}, append(p.nodes, edge)...)})
				}
			}
		}

		visited[node] = true
	}
	return 0, nil
}

func Tasks() {
	grid := getInput()
	grid[coord{x: 0, y: 0}] = 0

	final, _ := findPath(coord{0, 0}, coord{99, 99}, grid)
	fmt.Printf("Day 15 Task 01: %d\n", final)

	largeGrid := computeLargeGrid(grid)
	lFinal, _ := findPath(coord{0, 0}, coord{499, 499}, largeGrid)
	fmt.Printf("Day 15 Task 02: %d\n", lFinal)
}

func newValue(i int) int {
	for i > 9 {
		i -= 9
	}
	return i
}

func computeLargeGrid(grid map[coord]int) map[coord]int {
	lGrid := make(map[coord]int)

	size := int(math.Sqrt(float64(len(grid))))
	for x := 0; x < size*5; x++ {
		for y := 0; y < size*5; y++ {
			var origX, origY int

			origX = x % size
			origY = y % size

			origVal := grid[coord{origX, origY}]

			diffX := int((x - origX) / size)
			diffY := int((y - origY) / size)

			lGrid[coord{x: x, y: y}] = newValue(origVal + diffX + diffY)
		}
	}
	return lGrid
}

func displayGrid(grid map[coord]int) {
	size := int(math.Sqrt(float64(len(grid))))
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			fmt.Print(grid[coord{x: x, y: y}], "\t")
		}
		fmt.Print("\n")
	}
}

func getInput() map[coord]int {
	bytes, _ := ioutil.ReadFile(filename)

	lines := strings.Split(strings.TrimRight(string(bytes), "\n"), "\n")

	grid := map[coord]int{}
	for x := 0; x < len(lines); x++ {
		line := lines[x]
		for y := 0; y < len(line); y++ {
			cost := digitMap[line[y]]
			grid[coord{x, y}] = cost
		}
	}
	return grid
}
