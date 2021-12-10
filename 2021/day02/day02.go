package day02

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)
const filename = "day02/aoc2.txt"
type Command struct {
    direction string
    steps int
}

func Tasks() {
    cmds := getInputs()

    var pos, depth int

    for _, cmd := range cmds {
        switch cmd.direction {
        case "forward":
            pos += cmd.steps
        case "down":
            depth += cmd.steps
        case "up":
            depth -= cmd.steps
        }
    }

    fmt.Printf("Day 02 Task 01: %d\n", pos*depth)

    var aim int
    pos = 0
    depth = 0

    for _, cmd := range cmds {
        switch cmd.direction {
        case "forward":
            pos += cmd.steps
            depth += aim * cmd.steps
        case "down":
            aim += cmd.steps
        case "up":
            aim -= cmd.steps
        }
    }
    fmt.Printf("Day 02 Task 02: %d\n", pos*depth)
}

func getInputs() []Command {
    bytes, _ := ioutil.ReadFile(filename)

    lines := strings.Split(string(bytes), "\n")
    cmds := []Command{}

    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        instr := strings.Split(line, " ")
        val, _ := strconv.Atoi(instr[1])
        cmds = append(cmds, Command{instr[0], val}) 
    }
    return cmds
}

