package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func problem1() {
	file, _ := os.Open("data.txt")
	sum := 0
	scanner := bufio.NewScanner(file)
	lineCount := 0
	var allLines []string
	var allNumbers [][]*Info
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
		allNumbers = append(allNumbers, getLineData(scanner.Text(), lineCount))
		lineCount++
	}

	for i := range allNumbers {
		for j := range allNumbers[i] {
			x := allNumbers[i][j]
			var linesToSearch []string
			if x.lineNr == 0 {
				linesToSearch = append(linesToSearch, allLines[x.lineNr])
				linesToSearch = append(linesToSearch, allLines[x.lineNr+1])
			} else if x.lineNr == len(allNumbers)-1 {
				linesToSearch = append(linesToSearch, allLines[x.lineNr-1])
				linesToSearch = append(linesToSearch, allLines[x.lineNr])
			} else {
				linesToSearch = append(linesToSearch, allLines[x.lineNr-1])
				linesToSearch = append(linesToSearch, allLines[x.lineNr])
				linesToSearch = append(linesToSearch, allLines[x.lineNr+1])
			}
			sum += x.SearchForAnySymbol(linesToSearch)

		}
	}
	print("Problem 1: ")
	println(sum)
	file.Close()
}

func (s *Info) SearchForAnySymbol(lines []string) int {
	rangeOfNumber := len(strconv.Itoa(s.number))
	start := 0
	if s.pos > 0 {
		rangeOfNumber += 1
		start = s.pos - 1
	}
	// check if the range is out of bounds of the line length
	outOfBoundsCheck := (s.pos + rangeOfNumber) > len(lines[0])
	if outOfBoundsCheck {
		rangeOfNumber -= 1
	}
	for x := range lines {
		chars := "+#$*@/=%-&"
		found := strings.ContainsAny(lines[x][start:start+rangeOfNumber+1], chars)
		if found {
			return s.number
		}
	}
	return 0
}
