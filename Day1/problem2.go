package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var wordMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func problem2() {
	// this is gross gotta redo it, but tried to get it done with regex
	file, _ := os.Open("data.txt")

	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		replaced := findReplace(scanner.Text())
		re := regexp.MustCompile("[1-9]")
		data := re.FindAllString(replaced, -1)
		number, _ := strconv.Atoi(data[0] + data[len(data)-1])
		sum += number
	}

	print("Problem 2: ")
	println(sum)
}

func findReplace(input string) string {
	for k, v := range wordMap {
		input = strings.Replace(input, k, v, -1)
	}
	return input
}
