package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	left  string
	right string
}

var network map[string]node

func main() {

	part1("puzzle.txt")
}

var lrLine string

func part1(puzzle string) {
	f, err := os.OpenFile(puzzle, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	network = make(map[string]node)

	scanner := bufio.NewScanner(f)
	currentElement := "AAA"
	lineCounter := 0
	for scanner.Scan() {

		if lineCounter == 0 {
			lrLine = scanner.Text()
		} else {
			line := scanner.Text()

			// skip emtpy line in puzzle
			if line == "" {
				continue
			}
			splits := strings.Split(line, " = ")
			nodes := strings.Split(splits[1], ", ")

			// save start element
			//if currentElement == "" {
			//	currentElement = splits[0]
			//}

			network[splits[0]] = node{
				left:  strings.Replace(nodes[0], "(", "", -1),
				right: strings.Replace(nodes[1], ")", "", -1),
			}

		}

		lineCounter++
	}

	fmt.Printf("direction: %v\n", lrLine)
	fmt.Printf("network: %v\n", network)

	stepCounter := 0
	lrPointer := 0

	for stepCounter = 0; ; stepCounter++ {
		nextElement := getElement(network[currentElement], string(lrLine[lrPointer]))
		lrPointer++
		if lrPointer >= len(lrLine) {
			lrPointer = 0
		}

		if nextElement == "ZZZ" {
			break
		}
		currentElement = nextElement
	}

	fmt.Printf("steps: %v\n", stepCounter+1)
}

func getElement(n node, lr string) string {
	if lr == "L" {
		return n.left
	} else if lr == "R" {
		return n.right
	} else {
		panic("unknown lr string")
	}
}
