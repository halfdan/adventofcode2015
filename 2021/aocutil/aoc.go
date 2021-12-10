package aocutil

import (
	"io/ioutil"
	"strings"
)

func GetInputs(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), "\n")
}
