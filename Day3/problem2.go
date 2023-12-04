package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func problem2() {
	file, _ := os.Open("data.txt")
	defer file.Close()
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

	for gearIndex, line := range allLines {
		gearPos := getGearData(line)

		// returns slice of slice but each second slice value is only of index 0
		for _, pos := range gearPos {
			var found []*Info
			for _, line := range allNumbers {
				for _, info := range line {
					if info.SearchForGear(gearIndex, pos[0], len(allLines)) {
						found = append(found, info)
					}
				}
			}
			if len(found) == 2 {
				sum += (found[0].number * found[1].number)
			}
		}

	}
	print("Problem 2: ")
	println(sum)

}
func (s *Info) SearchForGear(index int, gearPos int, maxLines int) bool {
	rangeOfNumber := len(strconv.Itoa(s.number))
	end := s.pos + rangeOfNumber
	for i := index - 1; i <= index+1; i++ {
		if i < 0 || i > maxLines || s.lineNr != i {
			continue
		}
		// checks if position of gear is within the search radius around the number
		if s.pos-1 <= gearPos && gearPos <= end {
			return true
		}
	}
	return false
}

func getGearData(line string) [][]int {
	re := regexp.MustCompile("[*]")
	gears := re.FindAllStringIndex(line, -1)
	return gears
}
