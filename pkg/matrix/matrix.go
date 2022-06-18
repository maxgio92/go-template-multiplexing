package matrix

// (ordinate)
// y
// ^	A	1	Z
// |	B	2	Y
// |		3	X
// ------> x (abscissa)

func GetCombinations(parts []Part) []string {
	combinations := []string{}
	combination_sum := ""
	completed := false

	// For each time the last part has been reached
	// exit from recursion until reaching this:
	for {
		combination_sum = ""

		// Start always from the first part (x=0)
		gotoNextPart(&combinations, &combination_sum, 0, &parts[0], parts, &completed)

		if parts[0].Ordinate == len(parts[0].Combinations) || completed {
			break
		}
	}

	return combinations
}

func gotoNextPart(combinations *[]string, combination_sum *string, abscissa int, part *Part, parts []Part, completed *bool) {

	if abscissa+1 < len(parts) { // Until the last part is reached

		*combination_sum += string(part.Combinations[part.Ordinate])

		// Move forward
		abscissa++
		part = &parts[abscissa]
		gotoNextPart(combinations, combination_sum, abscissa, part, parts, completed)

	} else { // When the last part is reached

		for _, combination := range part.Combinations {
			*combinations = append(*combinations, string(*combination_sum+string(combination)))
		}

		// Move backward
		abscissa--
		part = &parts[abscissa]

		// Store where we gone
		scrollDownPrevPartCombination(part, parts, abscissa, completed)
	}
}

func scrollDownPrevPartCombination(part *Part, parts []Part, abscissa int, completed *bool) {

	if part.Ordinate+1 < len(part.Combinations) {
		part.Ordinate++

	} else {
		part.Ordinate = 0
		abscissa--

		if abscissa >= 0 {
			scrollDownPrevPartCombination(&parts[abscissa], parts, abscissa, completed)
		} else {
			*completed = true
		}
	}
}
