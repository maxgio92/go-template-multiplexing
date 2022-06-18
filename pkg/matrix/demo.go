package matrix

import "fmt"

var (
	parts = []Part{
		{Ordinate: 0, Combinations: []Combination{"A", "B"}},                // part
		{Ordinate: 0, Combinations: []Combination{"1", "2", "3", "4", "5"}}, // part
		{Ordinate: 0, Combinations: []Combination{"w", "x", "y", "z"}},      // part
		{Ordinate: 0, Combinations: []Combination{"E", "F", "G", "H"}},      // part
		{Ordinate: 0, Combinations: []Combination{"A", "B"}},                // part
	}
)

func RunDemo() {
	for _, v := range parts {
		fmt.Println(v.Combinations)
	}
	combinations := GetCombinations(parts)
	fmt.Println(combinations)
}
