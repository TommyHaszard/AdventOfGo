package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func problem1() {
	file, _ := os.Open("data.txt")

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		gameRe := regexp.MustCompile("[0-9]+")
		gameNum, _ := strconv.Atoi(gameRe.FindString(line))
		stripGame := strings.Split(line, ":")
		goesOver := numCheck(stripGame[1])
		if goesOver {
			println(gameNum)
			sum += gameNum
		}
	}
	print("Problem 2: ")
	println(sum)
}

func numCheck(input string) bool {
	sets := strings.Split(input, ";")
	for i := range sets {
		reNum := regexp.MustCompile("[0-9]+")
		colours := strings.Split(sets[i], ",")
		for x := range colours {
			if strings.Contains(colours[x], "red") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				if num > 12 {
					return false
				}
			}
			if strings.Contains(colours[x], "green") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				if num > 13 {
					return false
				}
			}
			if strings.Contains(colours[x], "blue") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				if num > 14 {
					return false
				}
			}
		}
	}
	// }
	return true
}
