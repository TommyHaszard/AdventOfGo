package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func problem2() {
	start := time.Now()
	fileName := "data.txt"
	data := parserProblem2(fileName)
	output := algo(data)

	elapsed := time.Since(start)
	log.Printf("Timer: %s", elapsed)
	print("Problem 2: ")
	println(output)
}

func parserProblem2(input string) *Data {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	first := true
	time := 0
	distance := 0
	for scanner.Scan() {
		if first {
			colonRemoved := strings.Split(scanner.Text(), ":")
			whiteSpaceRemoved := strings.Replace(colonRemoved[1], " ", "", -1)
			time, _ = strconv.Atoi(whiteSpaceRemoved)
			first = false
		} else {
			colonRemoved := strings.Split(scanner.Text(), ":")
			whiteSpaceRemoved := strings.Replace(colonRemoved[1], " ", "", -1)
			distance, _ = strconv.Atoi(whiteSpaceRemoved)
		}
	}
	return &Data{
		distance: distance,
		time:     time,
	}
}
