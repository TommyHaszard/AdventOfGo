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

func problem2() {
	start := time.Now()
	fileName := "data.txt"
	matrix, seeds, keymap := parserProblem2(fileName)
	output := math.MaxInt
	for seed := range seeds {
		solved := solveSeed(seed, matrix, keymap)
		if solved < output {
			output = solved
		}
	}

	print("Problem 2: ")
	println(output)
	elapsed := time.Since(start)
	log.Printf("Timer: %s", elapsed)
}

func parserProblem2(input string) (map[string][]*Line, map[int]bool, []string) {
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
			nums := numReg.FindAllString(scanner.Text(), -1)
			seed := 0
			for x, num := range nums {
				if x%2 == 0 {
					seed, _ = strconv.Atoi(num)
				} else {
					ranger, _ := strconv.Atoi(num)
					for i := seed; i < seed+ranger; i++ {
						seeds[i] = true
					}
				}
			}
			first = false
			continue
		}
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
