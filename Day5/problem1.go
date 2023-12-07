package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Line struct {
	destination int
	source      int
	numRange    int
}

func problem1() {
	start := time.Now()
	fileName := "data.txt"
	matrix, seeds, keymap := parserProblem1(fileName)
	output := math.MaxInt
	for seed := range seeds {
		solved := solveSeed(seed, matrix, keymap)
		if solved < output {
			output = solved
		}
	}

	print("Problem 1: ")
	println(output)
	elapsed := time.Since(start)
	log.Printf("Timer: %s", elapsed)
}

func parserProblem1(input string) (map[string][]*Line, map[int]bool, []string) {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	matrix := make(map[string][]*Line)
	seeds := make(map[int]bool, 0)
	keymap := make([]string, 0)
	currentMap := ""
	first := true
	for scanner.Scan() {
		if first {
			numReg := regexp.MustCompile("[0-9]+")
			num := numReg.FindAllString(scanner.Text(), -1)
			for i := range num {
				seed, _ := strconv.Atoi(num[i])
				seeds[seed] = true

			}
		}
		first = false
		cardRemoved := strings.Split(scanner.Text(), ":")
		lineCount := 0
		for _, strings := range cardRemoved {
			wordReg := regexp.MustCompile("[a-zA-Z]+")
			word := wordReg.FindString(strings)
			numReg := regexp.MustCompile("[0-9]+")
			num := numReg.FindAllString(strings, -1)
			if word != "" {
				currentMap = word
				keymap = append(keymap, currentMap)
				lineCount = 0
				matrix[currentMap] = make([]*Line, 0)
				break
			}
			if num != nil {
				destination, _ := strconv.Atoi(num[0])
				source, _ := strconv.Atoi(num[1])
				rangeNum, _ := strconv.Atoi(num[2])
				currentLine := &Line{
					destination: destination,
					source:      source,
					numRange:    rangeNum,
				}
				matrix[currentMap] = append(matrix[currentMap], currentLine)
				lineCount++
				break
			}
		}
	}
	return matrix, seeds, keymap
}
