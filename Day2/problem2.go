package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func problem2() {
	file, _ := os.Open("data.txt")

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stripGame := strings.Split(line, ":")
		power := sumPower(stripGame[1])
		sum += power
	}
	print("Problem 2: ")
	println(sum)
}

func sumPower(input string) int {
	sets := strings.Split(input, ";")
	greenMax := 0
	blueMax := 0
	redMax := 0

	for i := range sets {
		reNum := regexp.MustCompile("[0-9]+")
		colours := strings.Split(sets[i], ",")

		for x := range colours {
			if strings.Contains(colours[x], "red") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				redMax = int(math.Round(math.Max(float64(redMax), float64(num))))
			}
			if strings.Contains(colours[x], "green") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				greenMax = int(math.Round(math.Max(float64(greenMax), float64(num))))
			}
			if strings.Contains(colours[x], "blue") {
				num, _ := strconv.Atoi(reNum.FindString(colours[x]))
				blueMax = int(math.Round(math.Max(float64(blueMax), float64(num))))
			}
		}
	}

	power := greenMax * blueMax * redMax
	return power
}
