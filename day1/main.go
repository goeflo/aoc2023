package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Printf("aoc 2023 day 1 ...\n")

	part1("example_part1.txt")
	part1("puzzle.txt")

	part2("example_part2.txt")
	part2("puzzle.txt")
}

func part2(inputfile string) {
	fmt.Printf("part2 input file: %v\n", inputfile)
	f, err := os.OpenFile(inputfile, os.O_RDONLY, os.ModePerm)
	panicIfError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sumOfCalibrationValues := 0
	for scanner.Scan() {
		sumOfCalibrationValues += getCalibrationValuePart2(scanner.Text())
	}

	fmt.Printf("sum of calibration value: %v\n", sumOfCalibrationValues)

}

func getCalibrationValuePart2(line string) int {
	line = replaceWords(line)
	return getCalibrationValue(line)
}

func replaceWords(line string) string {
	line = strings.ReplaceAll(line, "one", "o1e")
	line = strings.ReplaceAll(line, "two", "t2o")
	line = strings.ReplaceAll(line, "three", "t3e")
	line = strings.ReplaceAll(line, "four", "f4r")
	line = strings.ReplaceAll(line, "five", "f5e")
	line = strings.ReplaceAll(line, "six", "s6x")
	line = strings.ReplaceAll(line, "seven", "s7n")
	line = strings.ReplaceAll(line, "eight", "e8t")
	line = strings.ReplaceAll(line, "nine", "n9e")
	return line
}

func part1(inputfile string) {
	fmt.Printf("part1 input file: %v\n", inputfile)
	f, err := os.OpenFile(inputfile, os.O_RDONLY, os.ModePerm)
	panicIfError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sumOfCalibrationValues := 0
	for scanner.Scan() {
		sumOfCalibrationValues += getCalibrationValue(scanner.Text())
	}

	fmt.Printf("sum of calibration value: %v\n", sumOfCalibrationValues)
}

func getCalibrationValue(line string) int {
	digit := regexp.MustCompile(`\d`)
	digits := digit.FindAllString(line, -1)
	calibrationValue := fmt.Sprintf("%v%v", digits[0], digits[len(digits)-1])
	value, err := strconv.Atoi(calibrationValue)
	panicIfError(err)
	return value
}
