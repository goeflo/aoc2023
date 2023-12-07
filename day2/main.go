package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var red uint16
var green uint16
var blue uint16

type bag struct {
	red   int
	green int
	blue  int
}

type games []bag

func main() {

	bag := bag{red: 12, green: 13, blue: 14}

	games := readGames("example_part1.txt")
	part1(bag, games)
	part2(bag, games)

	games = readGames("puzzle.txt")
	part1(bag, games)
	part2(bag, games)

}

func part2(bag bag, games []games) {
	powerSum := 0
	for _, game := range games {
		powerSum += getPower(getMinimumBag(game))
	}
	fmt.Printf("minimum power sum: %v\n", powerSum)
}

func getPower(b bag) int {
	return b.red * b.green * b.blue
}

func getMinimumBag(games games) bag {
	minimumBag := bag{red: 0, green: 0, blue: 0}
	for _, game := range games {
		if game.red > minimumBag.red {
			minimumBag.red = game.red
		}
		if game.green > minimumBag.green {
			minimumBag.green = game.green
		}
		if game.blue > minimumBag.blue {
			minimumBag.blue = game.blue
		}
	}
	return minimumBag
}

func part1(bag bag, games []games) {
	sum := getPossibleGameSum(bag, games)
	fmt.Printf("possible game sum: %v\n", sum)
}

func readGames(filename string) []games {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	panicIfError(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	allGames := []games{}
	for scanner.Scan() {
		allGames = append(allGames, getGame(scanner.Text()))
	}
	return allGames
}

func getPossibleGameSum(bag bag, games []games) int {
	sum := 0
	for k, v := range games {
		possible := true
		for _, v2 := range v {
			if v2.blue > bag.blue || v2.green > bag.green || v2.red > bag.red {
				possible = false
				break
			}
		}
		if possible {
			sum += k + 1
		}
	}
	return sum
}
func getGame(line string) games {
	games := []bag{}
	line = line[strings.Index(line, ":")+1:]
	sets := strings.Split(line, ";")

	for _, set := range sets {
		gameBag := bag{}
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			spaceIdx := strings.Index(cube, " ")
			if strings.HasSuffix(cube[spaceIdx:], "blue") {
				v, err := strconv.Atoi(cube[0:spaceIdx])
				panicIfError(err)
				gameBag.blue = v
			} else if strings.HasSuffix(cube[spaceIdx:], "red") {
				v, err := strconv.Atoi(cube[0:spaceIdx])
				panicIfError(err)
				gameBag.red = v
			} else if strings.HasSuffix(cube[spaceIdx:], "green") {
				v, err := strconv.Atoi(cube[0:spaceIdx])
				panicIfError(err)
				gameBag.green = v
			}
		}
		games = append(games, gameBag)
	}

	return games
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
