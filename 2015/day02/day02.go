package day02

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const filename = "day02/aoc2.txt"

func Tasks() {
    data := getInputs()

    lines := strings.Split(data, "\n") 
    result := computeWrappingPaper(lines)
    fmt.Printf("Day 2 Task 1: %d\n", result)

    ribbon := computeRibbon(lines)
    fmt.Printf("Day 2 Task 2: %d\n", ribbon)
}

func computeRibbon(lines []string) int {
    var sum int
    var lens []int

    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        vals := strings.Split(line, "x")
        lens = valsToSortedIntSlice(vals)
        a, b, c := lens[0], lens[1], lens[2]

        sum += 2*(a+b) + a*b*c 
    }
    return sum
}

func computeWrappingPaper(lines []string) int {
    var sum int
    var lens []int

    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        vals := strings.Split(line, "x")
        lens = valsToSortedIntSlice(vals)
        a, b, c := lens[0], lens[1], lens[2]

        sum += 2*(a*b + a*c + b*c) + a*b
    }
    return sum
}

func valsToSortedIntSlice(vals []string) []int {
    res := make([]int, 0, len(vals))

    for _, val := range vals {
        num, _ := strconv.Atoi(val)
        res = append(res, num)
    }
    sort.Ints(res)

    return res 
}

func getInputs() string {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
    }

    return string(bytes)
}
