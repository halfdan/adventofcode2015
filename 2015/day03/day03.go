package day03

import (
	"fmt"
	"io/ioutil"
)


type Point struct {
    x, y int
}

const filename = "day03/aoc3.txt"

func Tasks() {
    data := getInputs()

    housesVisitedAtLeastOnce := computeSingleVisits(data)
    fmt.Printf("Day 3 Task 1: %d\n", housesVisitedAtLeastOnce)

    housesVisitedWithRobo := computeSingleVisitsWithRoboSanta(data)
    fmt.Printf("Day 3 Task 2: %d\n", housesVisitedWithRobo)
}

func computeSingleVisitsWithRoboSanta(input string) int {
    visits := map[Point]bool{}   
    x, y, x2, y2 := 0, 0, 0, 0

    visits[Point{x, y}] = true

    for _, ch := range input {
        switch ch {
        case '>':
            x += 1
        case '<':
            x -= 1
        case 'v':
            y += 1
        case '^':
            y -= 1
        }

        visits[Point{x, y}] = true
        x, x2 = x2, x
        y, y2 = y2, y
    }

    return len(visits)
}

func computeSingleVisits(input string) int {
    visits := map[Point]bool{}   
    x, y := 0, 0

    visits[Point{x, y}] = true

    for _, ch := range input {
        switch ch {
        case '>':
            x += 1
        case '<':
            x -= 1
        case 'v':
            y += 1
        case '^':
            y -= 1
        }

        visits[Point{x, y}] = true
    }

    return len(visits)
}

func getInputs() string {
    data, _ := ioutil.ReadFile(filename)

    return string(data)
}

