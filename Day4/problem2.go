package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func problem2() {
	file, _ := os.Open("data.txt")
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	cardMap, data := initProblem(scanner)
	currentCard := 1
	for _, line := range data {
		cardRemoved := strings.Split(line, ":")
		line := strings.Split(cardRemoved[1], "|")
		val := cardMap[currentCard]
		wins := getLineWins(line[0], line[1])
		for i := 0; i <= val; i++ {
			addWinners(currentCard, cardMap, wins)
		}
		currentCard++
	}
	// add the total number of OG cards - 1
	sum += currentCard - 1
	for i := range cardMap {
		sum += cardMap[i]
	}

	print("Problem 2: ")
	println(sum)

}

func initProblem(scanner *bufio.Scanner) (map[int]int, []string) {
	cardMap := make(map[int]int)
	data := make([]string, 0)
	i := 1
	for scanner.Scan() {
		cardMap[i] = 0
		data = append(data, scanner.Text())
	}
	return cardMap, data
}

func addWinners(currentLine int, cardMap map[int]int, winnersCount int) {
	lineRange := winnersCount + currentLine
	for i := currentLine + 1; i <= lineRange; i++ {
		cardMap[i] += 1
	}
}

func getLineWins(winners string, possible string) int {
	re := regexp.MustCompile("[0-9]+")
	winNum := re.FindAllString(winners, -1)
	checkNum := getPossible(possible)
	count := 0
	for _, win := range winNum {
		for _, possible := range checkNum {
			if win == possible {
				count += 1
				break
			}
		}
	}
	return count
}
