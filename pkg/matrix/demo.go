package matrix

import "fmt"

var (
	columns = []Column{
		{OrdinateIndex: 0, Points: []string{"A", "B"}},                // part
		{OrdinateIndex: 0, Points: []string{"1", "2", "3", "4", "5"}}, // part
		{OrdinateIndex: 0, Points: []string{"w", "x", "y", "z"}},      // part
		{OrdinateIndex: 0, Points: []string{"E", "F", "G", "H"}},      // part
		{OrdinateIndex: 0, Points: []string{"A", "B"}},                // part
	}
)

func RunDemo() {
	for _, v := range columns {
		fmt.Println(v.Points)
	}
	combinations, err := GetColumnOrderedCombinationRows(columns)
	if err != nil {
		panic(err)
	}
	fmt.Println(combinations)
}
