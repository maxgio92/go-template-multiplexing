package template

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	t "html/template"

	"github.com/maxgio92/go-template-multiplexing/pkg/matrix"
)

// Returns a list of strings of executed templates from a template string
// input, by applying an arbitrary variables inventory with multiple values.
// The expecetd arguments are:

// - templateString: the input string of the template to execute.
// - inventory: the inventory map of the variable data structure to apply the
// template to. The key of the map is the name of the variable, that should
// match related annotation in the template. Each map item is a slice where in each slice item is
// a single variable value.
//
// The result is multiple templates from a single template string and multiple
// arbitrary variable values.
func MultiplexAndExecute(templateString string, inventory map[string][]string) ([]string, error) {
	supportedVariables, err := getSupportedVariables(templateString)
	if err != nil {
		return nil, err
	}

	myTemplateRegex, err := generateTemplateRegex(supportedVariables)
	if err != nil {
		panic(err)
	}
	myTemplatePattern := regexp.MustCompile(myTemplateRegex)

	myTemplateParts, err := cutTemplateString(templateString, closeDelimiter)
	if err != nil {
		panic(err)
	}

	results := []TemplatePart{}

	for _, v := range myTemplateParts {

		// match are the template parts matched against the template regex.
		myTemplatePartMatches := myTemplatePattern.FindStringSubmatch(v)

		// name is the variable data structure to apply the template part to.
		for i, variableName := range myTemplatePattern.SubexpNames() {

			// discard first variable name match and ensure a template part matched.
			if i > 0 && i <= len(myTemplatePartMatches) && myTemplatePartMatches[i] != "" {
				y := len(results)

				results = append(results, TemplatePart{
					TemplateString:  myTemplatePartMatches[i],
					MatchedVariable: variableName,
				})

				results[y].TemplateString = strings.ReplaceAll(results[y].TemplateString, "{{ "+variableName+" }}", "{{ . }}")
				results[y].Template = t.New(fmt.Sprintf("%d", y))
				results[y].Template, err = results[y].Template.Parse(results[y].TemplateString)
				if err != nil {
					return nil, err
				}

				// for each item (variable name) of MatchedVariable
				// Compose one Template and `execute()` it
				for _, value := range inventory[variableName] {
					o := new(bytes.Buffer)
					err = results[y].Template.Execute(o, value)
					if err != nil {
						return nil, err
					}
					results[y].Combinations = append(results[y].Combinations, matrix.Combination(o.String()))
				}
			}
		}
	}

	myParts := results

	matrixColumns := []matrix.Part{}

	for _, part := range myParts {
		matrixColumns = append(matrixColumns, part.Part)
	}

	return matrix.GetCombinations(matrixColumns), nil
}
