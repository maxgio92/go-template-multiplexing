package matrix

import "fmt"

// (ordinate)
// y
// ^
// |		4
// |		3	Z
// |	B	2	Y
// |	A	1	X
// ---------------> x (abscissa)

// A + 1 + X
// A + 1 + Y
// A + 1 + Z
// A + 2 + X
// ...
// B + 4 + Z
func GetColumnOrderedCombinationRows(columns []Column) ([]string, error) {
	rows := []string{}
	row := ""
	completed := false

	// For each time the last column has been reached
	// exit from recursion until reaching this:
	for {
		row = ""

		// Start always from the first column (x=0)
		err := gotoNextColumn(&rows, &row, 0, &columns[0], columns, &completed)
		if err != nil {
			return nil, err
		}

		if columns[0].OrdinateIndex == len(columns[0].Points) || completed {
			break
		}
	}

	return rows, nil
}

func gotoNextColumn(points *[]string, row *string, abscissaIndex int, column *Column, columns []Column, completed *bool) error {

	if abscissaIndex+1 < len(columns) { // Until the last column is reached

		sp, ok := column.Points[column.OrdinateIndex].(string)
		if !ok {
			return fmt.Errorf("a point of the matrix is not of supported types. Supported types are (string)")
		}
		*row += sp

		// Move forward
		abscissaIndex++
		column = &columns[abscissaIndex]
		gotoNextColumn(points, row, abscissaIndex, column, columns, completed)

	} else { // When the last column is reached

		for _, point := range column.Points {
			sp, ok := point.(string)
			if !ok {
				return fmt.Errorf("a point of the matrix is not of supported types. Supported types are (string)")
			}
			*points = append(*points, string(*row+sp))
		}

		// Move backward
		abscissaIndex--
		column = &columns[abscissaIndex]

		// Store where we gone
		scrollDownPrevColumnPoint(column, columns, abscissaIndex, completed)
	}
	return nil
}

func scrollDownPrevColumnPoint(column *Column, columns []Column, abscissaIndex int, completed *bool) {

	if column.OrdinateIndex+1 < len(column.Points) {
		column.OrdinateIndex++

	} else {
		column.OrdinateIndex = 0
		abscissaIndex--

		if abscissaIndex >= 0 {
			scrollDownPrevColumnPoint(&columns[abscissaIndex], columns, abscissaIndex, completed)
		} else {
			*completed = true
		}
	}
}
