package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func problem1() {
	file, _ := os.Open("data.txt")

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		re := regexp.MustCompile("[1-9]")
		data := re.FindAllString(scanner.Text(), -1)
		number, _ := strconv.Atoi(data[0] + data[len(data)-1])
		sum += number
	}

	print("Problem 1: ")
	println(sum)
}
