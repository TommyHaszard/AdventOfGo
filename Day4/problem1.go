package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func problem1() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cardRemoved := strings.Split(scanner.Text(), ":")
		line := strings.Split(cardRemoved[1], "|")
		winners := getWinners(line[0])
		sum += returnPoints(winners, line[1])
	}

	print("Problem 1: ")
	println(sum)
}

func getWinners(line string) map[string]int {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)
	winners := make(map[string]int)
	for _, x := range numbers {
		winners[x] = 0
	}
	return winners
}

func returnPoints(line map[string]int, possible string) int {
	data := getPossible(possible)
	count := 0
	for _, possible := range data {
		_, ok := line[possible]
		if ok {
			if count > 1 {
				count = count * 2
			} else {
				count += 1
			}
		}
	}
	return count
}
