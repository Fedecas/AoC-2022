package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const sep = "\r\n"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertToInt(s string) (v int) {
	v, err := strconv.Atoi(s)
	check(err)
	return
}

func reverseStack(stack []string) (res []string) {
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return
}

func copyStacks(stacks [][]string) (res [][]string) {
	for i := 0; i < len(stacks); i++ {
		aux := make([]string, len(stacks[i]))
		copy(aux, stacks[i])
		res = append(res, aux)
	}
	return
}

func getInput() (stacks [][]string, movements [][]int) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := string(bytes)
	parts := strings.Split(content, sep+sep)

	// make stacks array
	stacksInput := strings.Split(parts[0], sep)
	stacksNum := strings.Split(stacksInput[len(stacksInput)-1], "   ")
	stacks = make([][]string, len(stacksNum))
	for i := len(stacksInput) - 2; i >= 0; i-- {
		aux := strings.ReplaceAll(stacksInput[i], "    ", " ")
		crates := strings.Split(aux, " ")
		for j := 0; j < len(stacks); j++ {
			if crates[j] != "" {
				crate := strings.Trim(crates[j], "[]")
				stacks[j] = append(stacks[j], crate)
			}
		}
	}

	// make movements array
	movementsInput := strings.Split(parts[1], sep)
	for i := 0; i < len(movementsInput); i++ {
		movs := strings.Split(movementsInput[i], " ")
		if len(movs) == 6 {
			args := []int{convertToInt(movs[1]), convertToInt(movs[3]), convertToInt(movs[5])}
			movements = append(movements, args)
		}
	}

	return
}

func moveCrates(stacks [][]string, movement []int, oneAtATime bool) {
	count, from, to := movement[0], movement[1]-1, movement[2]-1
	moving := stacks[from][len(stacks[from])-count:]
	if oneAtATime {
		moving = reverseStack(moving)
	}

	stacks[from] = stacks[from][:len(stacks[from])-count]
	stacks[to] = append(stacks[to], moving...)
}

func getTopCratesAfterMove(stacks [][]string, movements [][]int, oneAtATime bool) (res string) {
	stacksCopy := copyStacks(stacks)
	for i := 0; i < len(movements); i++ {
		moveCrates(stacksCopy, movements[i], oneAtATime)
	}

	for i := 0; i < len(stacksCopy); i++ {
		res += stacksCopy[i][len(stacksCopy[i])-1]
	}

	return
}

func main() {
	stacks, movements := getInput()
	crates1 := getTopCratesAfterMove(stacks, movements, true)
	crates2 := getTopCratesAfterMove(stacks, movements, false)

	fmt.Println("1:", crates1)
	fmt.Println("2:", crates2)
}
