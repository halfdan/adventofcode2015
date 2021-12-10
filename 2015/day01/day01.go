package day01

import (
	"fmt"
	"io/ioutil"
    "errors"
)

const filename = "day01/aoc1.txt"

func Tasks() {
    data := getInputs()
    result := CountParenthesis(data)

    fmt.Printf("Day 1 Task 1: %d\n", result)  
    result2, ok := CountToBasement(data)
    if ok != nil {
        fmt.Println(ok)
    }
    fmt.Printf("Day 1 Task 2: %d\n", result2)  
}


func getInputs() string {
    bytes, ok := ioutil.ReadFile(filename)
    if ok != nil {
        fmt.Print(ok)
    }

    return string(bytes)
}

func CountParenthesis(input string) int {
    var count int
    for _, char := range input {
        if char == '(' { 
            count++ 
        } else if char == ')' { 
            count-- 
        }
    }
    return count
}

func CountToBasement(input string) (int, error) {
    var level int
    for idx, char := range input {
        if char == '(' {
            level++
        } else if char == ')' {
            level--
        }
        if level < 0 {
            return idx + 1, nil
        }
    }
    return 0, errors.New("elevator never goes to basement") 
}
