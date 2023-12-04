package main

import (
	"log"
	"regexp"
	"strconv"
)

type Info struct {
	lineNr       int
	pos          int
	lineNumCount int
	number       int
}

func main() {
	//problem1()
	problem2()
}

func getLineData(line string, lineCount int) []*Info {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllStringIndex(line, -1)
	var data []*Info
	for i, x := range numbers {
		// pass in i as we need to know how many numbers in the line later on
		number := CreateInfo(line, lineCount, x, i)
		//fmt.Printf("Line Number: %d - Pos: %d, Number: %d\n", number.lineNr, number.pos, number.number)
		data = append(data, number)
	}
	return data
}

func CreateInfo(line string, lineCount int, indices []int, numberCounter int) *Info {
	number, err := strconv.Atoi(line[indices[0]:indices[1]])
	if err != nil {
		log.Fatal(err)
	}
	return &Info{
		lineNr:       lineCount,
		pos:          indices[0],
		number:       number,
		lineNumCount: numberCounter,
	}
}
