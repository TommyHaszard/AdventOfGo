package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func problem1() {
	start := time.Now()
	fileName := "data.txt"
	dataSlice := parserProblem1(fileName)
	output := 0
	for _, data := range dataSlice {
		count := algo(data)
		if output != 0 {
			output *= count
		} else {
			output = count
		}
	}

	elapsed := time.Since(start)
	log.Printf("Timer: %s", elapsed)
	print("Problem 1: ")
	println(output)
}

func algo(data *Data) int {
	count := 0
	for holdTime := 1; holdTime < data.time-1; holdTime++ {
		distanceTravelled := holdTime * (data.time - holdTime)
		if distanceTravelled > data.distance {
			count++
		}
	}
	return count
}

func parserProblem1(input string) []*Data {
	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	dataSlice := make([]*Data, 0)
	first := true
	for scanner.Scan() {
		if first {
			numReg := regexp.MustCompile("[0-9]+")
			num := numReg.FindAllString(scanner.Text(), -1)
			for i := range num {
				time, _ := strconv.Atoi(num[i])
				data := &Data{
					time: time,
				}
				dataSlice = append(dataSlice, data)
			}
			first = false
		} else {
			numReg := regexp.MustCompile("[0-9]+")
			num := numReg.FindAllString(scanner.Text(), -1)
			for i := range num {
				distance, _ := strconv.Atoi(num[i])
				dataSlice[i].UpdateDistance(distance)
			}
		}
	}
	return dataSlice
}
