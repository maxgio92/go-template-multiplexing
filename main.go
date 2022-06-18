package main

import (
	"fmt"

	template "github.com/maxgio92/go-template-multiplexing/pkg/template"
)

var (
	myTemplate = `Items are made also of {{ Materials }} and one of the team member is {{ Team }} and one of the cost is {{ Costs }}`

	myInventoryMap = map[string][]string{
		"Materials": {"wool", "iron"},
		"Team":      {"Mario", "Luigi"},
		"Costs":     {"9", "14", "7"},
	}
)

func main() {
	result, err := template.MultiplexAndExecute(myTemplate, myInventoryMap)
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
