package day14

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const filename = "day14/aoc14.txt"

func Tasks() {
	template, mapping := getInput()

	polymer := computePolymer(template, mapping)
	freqs := computeFrequencies(polymer)
	chfreqs := make([]int, 0, len(freqs))
	for _, v := range freqs {
		chfreqs = append(chfreqs, v)
	}
	sort.Ints(chfreqs)
	fmt.Printf("Day 14 Task 01: %d\n", chfreqs[len(chfreqs)-1]-chfreqs[0])

	fmt.Printf("Day 14 Task 02: %d\n", doTheCounting(template, mapping))
}

func computeFrequencies(polymer string) map[rune]int {
	freqs := map[rune]int{}

	for _, ch := range polymer {
		freqs[ch] += 1
	}

	return freqs
}

func computePolymer(template string, mapping map[string][2]string) string {
	for i := 0; i < 10; i++ {
		polymer := ""
		for idx := range template {
			if idx == len(template)-1 {
				polymer += string(template[idx])
				continue
			}
			if value, exists := mapping[string(template[idx:idx+2])]; exists {
				polymer += value[0]
			}
		}
		template = polymer
	}
	return template
}

// getInput will read the input file and return:
// - the template string
// - a mapping for the rules
//
// The rules will be encoded as follows:
// NN -> C
// will be represented at:
// "NN" => [2]string{"NC", "CN"}
func getInput() (string, map[string][2]string) {
	bytes, _ := ioutil.ReadFile(filename)

	lines := strings.Split(strings.TrimRight(string(bytes), "\n"), "\n")

	polymer := lines[0]

	mapping := map[string][2]string{}
	for _, line := range lines[2:] {
		f := strings.Fields(line)
		mapping[f[0]] = [2]string{
			string(f[0][0]) + f[2],
			f[2] + string(f[0][1]),
		}
	}

	return polymer, mapping
}
