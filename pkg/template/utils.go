package template

import (
	"fmt"
	"regexp"
	"strings"
)

func generateTemplateRegex(variables []string) (string, error) {
	if len(variables) < 1 {
		return "", fmt.Errorf("at least one variable is required")
	}

	templateRegex := ``
	for _, v := range variables {
		templateRegex += `(?P<` + v + `>^.*` + openDelimiter + ` ` + v + ` ` + closeDelimiter + `.*$)?`
	}

	return templateRegex, nil
}

func getSupportedVariables(templateString string) ([]string, error) {
	return getVariablesFromTemplateString(templateString)
}

func getVariablesFromTemplateString(templateString string) ([]string, error) {
	rs := openDelimiter + ` (` + variableNameRegex + `) ` + closeDelimiter
	rp := regexp.MustCompile(rs)

	v := []string{}
	ss := rp.FindAllStringSubmatch(templateString, -1)
	if len(ss) < 1 {
		return nil, fmt.Errorf("cannot find supported variables")
	}

	for _, s := range ss {
		if len(s) < 1 {
			return nil, fmt.Errorf("cannot find supported variables")
		}
		v = append(v, s[1])
	}

	return v, nil
}

func cutTemplateString(t string, closeDelimiter string) ([]string, error) {
	var parts []string

	before, after, found := strings.Cut(t, closeDelimiter)
	if !found {
		return nil, fmt.Errorf("cannot cut input template string")
	}

	parts = append(parts, before+closeDelimiter)
	for {
		before, after, found = strings.Cut(after, closeDelimiter)
		if !found {
			break
		}
		parts = append(parts, before+closeDelimiter)
	}

	return parts, nil
}
