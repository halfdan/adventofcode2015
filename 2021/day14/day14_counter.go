package day14

import "log"

func parseTemplateIntoCounter(polymer string) map[string]int {
	var pairMap = make(map[string]int)

	for idx := 0; idx < len(polymer)-1; idx++ {
		key := string(polymer[idx]) + string(polymer[idx+1])
		pairMap[key]++
	}
	return pairMap
}

func doTheCounting(template string, rules map[string][2]string) int {
	counter := parseTemplateIntoCounter(template)

	for i := 0; i < 40; i++ {
		counter = workCounter(counter, rules)
	}

	return countElements(template, counter)
}

func workCounter(counter map[string]int, rules map[string][2]string) map[string]int {
	newCounter := make(map[string]int)

	for k, v := range counter {
		newKeys, ok := rules[k]
		if !ok {
			log.Fatalf("Could not find key %s in best rules. This should not have happened", k)
		}

		newCounter[newKeys[0]] += v
		newCounter[newKeys[1]] += v
	}

	return newCounter
}

func countElements(template string, counter map[string]int) int {
	first := string(template[0])              // First element in original template
	last := string(template[len(template)-1]) // Last element in original template
	counts := make(map[string]int)
	most := 0
	least := 1<<63 - 1

	for k, v := range counter {
		counts[string(k[0])] += v
		counts[string(k[1])] += v
	}

	for k, v := range counts {
		counts[k] = v / 2

		if k == first || k == last {
			counts[k]++
		}

		if counts[k] > most {
			most = counts[k]
		}

		if counts[k] < least {
			least = counts[k]
		}
	}

	return most - least
}
