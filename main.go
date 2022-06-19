package main

import (
	"fmt"

	template "github.com/maxgio92/go-template-multiplexing/pkg/template"
)

type MyCustomType string

var (
	myTemplate = `Items are made also of {{ Materials }} and one of the team member is {{ Team }} and one of the cost is {{ Costs }} AGAIN {{ Costs }} AGAGAGAIN {{ Materials }}`

	myInventoryMap = map[string][]interface{}{}
	Materials      = []MyCustomType{"wool", "iron"}
	Team           = []MyCustomType{"Mario", "Luigi"}
	Costs          = []MyCustomType{"9", "14", "7"}
)

func main() {
	for _, v := range Materials {
		myInventoryMap["Materials"] = append(myInventoryMap["Materials"], interface{}(v))
	}
	for _, v := range Team {
		myInventoryMap["Team"] = append(myInventoryMap["Team"], interface{}(v))
	}
	for _, v := range Costs {
		myInventoryMap["Costs"] = append(myInventoryMap["Costs"], interface{}(v))
	}

	result, err := template.MultiplexAndExecute(myTemplate, myInventoryMap)
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
