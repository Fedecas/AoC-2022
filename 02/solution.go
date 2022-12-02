package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const sep = "\r\n"

var pointsFirstStrat = map[string]int{
	"AX": 4, // draw + rock
	"AY": 8, // win + paper
	"AZ": 3, // lose + scissors
	"BX": 1, // lose + rock
	"BY": 5, // draw + paper
	"BZ": 9, // win + scissors
	"CX": 7, // win + rock
	"CY": 2, // lose + paper
	"CZ": 6, // draw + scissors
}

var pointsSecondStrat = map[string]int{
	"AX": 3, // lose + scissors
	"AY": 4, // draw + rock
	"AZ": 8, // win + paper
	"BX": 1, // lose + rock
	"BY": 5, // draw + paper
	"BZ": 9, // win + scissors
	"CX": 2, // lose + paper
	"CY": 6, // draw + scissors
	"CZ": 7, // win + rock
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() (res []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := string(bytes)

	lines := strings.Split(content, sep)
	for i := 0; i < len(lines); i++ {
		round := strings.Split(lines[i], " ")
		if len(round) > 1 {
			res = append(res, round[0]+round[1])
		}
	}
	return
}

func firstStrat(rounds []string) (res int) {
	for i := 0; i < len(rounds); i++ {
		res += pointsFirstStrat[rounds[i]]
	}
	return
}

func secondStrat(rounds []string) (res int) {
	for i := 0; i < len(rounds); i++ {
		res += pointsSecondStrat[rounds[i]]
	}
	return
}

func main() {
	rounds := getInput()

	points1 := firstStrat(rounds)
	points2 := secondStrat(rounds)

	fmt.Println("1:", points1)
	fmt.Println("2:", points2)
}
