package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type partNumber struct {
	x      int
	y      int
	x2     int
	number int
}

type symbol struct {
	x int
	y int
}

type line struct {
	pn []partNumber
	sy []symbol
}

var maxX = 0
var maxY = 0

var digitRegex = regexp.MustCompile(`\d+`)
var symbolRegex = regexp.MustCompile(`[^.\d]`) // no a digit and not .

func main() {

	part1("example_part1.txt")
	//part1("puzzle.txt")
}

func part1(filename string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	panicIfError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineCount := 0
	pn := []partNumber{}
	symbls := []symbol{}

	lines := []line{}
	for scanner.Scan() {
		if maxX == 0 {
			maxX = len(scanner.Text())
		}
		fmt.Printf("%v - %v\n", lineCount, scanner.Text())
		pn = getPartNumbers(scanner.Text(), lineCount)
		symbls = getSymbols(scanner.Text(), lineCount)
		lines = append(lines, line{pn: pn, sy: symbls})
		lineCount++
	}
	maxY = lineCount

	for i := 0; i < len(lines); i++ {
		fmt.Printf("%v - %+v / %+v\n", i, lines[i].pn, lines[i].sy)
	}

	partNumbers := getAdjPartNumbers(lines)
	fmt.Printf("part numbers: %v\n", partNumbers)

	sum := 0
	for _, v := range partNumbers {
		sum += v
	}
	fmt.Printf("sum of part numbers: %v\n", sum)
}

func getAdjPartNumbers(lines []line) []int {

	partNumbers := []int{}
	symbols := []symbol{}
	for _, line := range lines {
		symbols = append(symbols, line.sy...)
	}

	for _, line := range lines {
		for _, pn := range line.pn {
			x1 := pn.x - 1
			y1 := pn.y - 1
			x2 := pn.x2 + 1
			y2 := pn.y + 1
			if x1 < 0 {
				x1 = 0
			}
			if y1 < 0 {
				y1 = 0
			}
			if x2 > maxX {
				x2 = maxX
			}
			if y2 > maxY {
				y2 = maxY
			}

			//fmt.Printf("box %v %v / %v %v\n", x1, y1, x2, y2)
			isPartNumber := false
			for _, sy := range symbols {
				//fmt.Printf("check symbol: %+v\n", sy)
				// symbol is inside surounding box
				if !isPartNumber && sy.x >= x1 && sy.y >= y1 && sy.x <= x2 && sy.y <= y2 {
					partNumbers = append(partNumbers, pn.number)
					isPartNumber = true
					break
				}
			}

			fmt.Printf("%v is part number: %v\n", pn, isPartNumber)
		}
	}
	return partNumbers
}

func getSymbols(l string, lc int) []symbol {
	symbols := []symbol{}
	index := symbolRegex.FindAllStringIndex(l, -1)

	if len(index) == 0 {
		return symbols
	}

	for _, idx := range index {
		s := symbol{x: idx[0], y: lc}
		symbols = append(symbols, s)
	}
	return symbols
}

func getPartNumbers(l string, lc int) []partNumber {
	pns := []partNumber{}
	digits := digitRegex.FindAllString(l, -1)
	index := digitRegex.FindAllStringIndex(l, -1)

	if len(digits) == 0 {
		return pns
	}

	for k, v := range digits {
		intValue, err := strconv.Atoi(v)
		panicIfError(err)
		pn := partNumber{
			number: intValue,
			x:      index[k][0],
			x2:     index[k][1] - 1,
			y:      lc,
		}
		pns = append(pns, pn)
	}

	return pns
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
