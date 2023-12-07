package main

func main() {
	problem1()
	problem2()
}

func solveSeed(input int, matrix map[string][]*Line, keymap []string) int {
	output := input
	var found bool
	for _, mapType := range keymap {
		slice := matrix[mapType]
		for _, line := range slice {
			output, found = checkRange(line, input)
			if found {
				break
			}
		}
		if found {
			input = output
			found = false
		}
	}
	// if input isnt found within any of the ranges then it must be equal to itself
	return output
}

func checkRange(line *Line, input int) (int, bool) {
	destination := line.destination
	source := line.source
	rangeNum := line.numRange
	sourceRange := source + rangeNum
	// if input within the range for this line return true and return the output
	if input >= source && input <= sourceRange {
		dif := input - source
		output := dif + destination
		return output, true
	}
	return input, false
}
