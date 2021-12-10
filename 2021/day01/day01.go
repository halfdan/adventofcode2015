package day01


import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

const filename = "day01/aoc1.txt"

func Tasks() {
    data := getInputs()


    var count int
    var last int

    last = data[0]
    for _, num := range data {
        if num > last {
            count++
        }
        last = num
    }
    fmt.Printf("Day 01 Task 01: %d\n", count)

    count = 0
    for idx := 3; idx < len(data); idx++ {
        if data[idx] + data[idx-1] + data[idx-2] > data[idx-1] + data[idx-2] + data[idx-3] {
            count++
        }
    }
    fmt.Printf("Day 02 Task 02: %d\n", count)
}


func getInputs() []int {
    bytes, ok := ioutil.ReadFile(filename)
    if ok != nil {
        fmt.Print(ok)
    }

    lines := strings.Split(string(bytes), "\n")

    nums := []int{}
    for _, val := range lines {
        num, _ := strconv.Atoi(val)
        nums = append(nums, num)
    }

    return nums
}
