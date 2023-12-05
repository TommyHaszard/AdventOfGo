package main

import (
	"regexp"
)

func main() {
	//problem1()
	problem2()
}

func getPossible(line string) []string {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)
	possible := make([]string, 0)
	possible = append(possible, numbers...)
	return possible
}
