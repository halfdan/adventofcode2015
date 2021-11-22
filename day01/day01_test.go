package day01

import (
    "testing"
)

var countTests = []struct {
    input string
    expected int
}{
    {"(())", 0},
    {"()()", 0},
    {"(((", 3},
    {"(()(()(", 3},
    {"))(((((", 3},
    {"())", -1},
    {")))", -3},
    {")())())", -3},
}

func TestCountParenthesis(t *testing.T) {
    for _, test := range countTests {
        ts := CountParenthesis(test.input)
        if ts != test.expected {
            t.Fatalf("CountParenthesis('%s') = %v, want %v",
                test.input, ts, test.expected)
        }
    }
}

func TestCountToBasement(t *testing.T) {
    lvl, _ := CountToBasement(")")
    if lvl != 1 {
        t.Fatal("Not correct")
    }
}
