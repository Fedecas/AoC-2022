package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

func getInput() (res [][]int) {
	const sep string = "\r\n"

	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := strings.TrimSpace(string(bytes))

	lines := strings.Split(content, sep)
	for i := 0; i < len(lines); i++ {
		instr := strings.Split(lines[i], " ")
		value := make([]int, 0)
		if len(instr) > 1 {
			value = append(value, convertToInt(instr[1]))
		}
		res = append(res, value)
	}
	return
}

func enqueueInstructions(instructions [][]int) (res []int) {
	for _, instr := range instructions {
		res = append(res, 0)

		if len(instr) != 0 {
			res = append(res, instr[0])
		}
	}
	return
}

func shouldComputeSignalStrength(cycle int) (res bool) {
	cycleValuesToCompute := []int{20, 60, 100, 140, 180, 220}
	for _, c := range cycleValuesToCompute {
		if cycle == c {
			res = true
		}
	}
	return
}

func sumSignalStrengths(queue []int) (res int) {
	x := 1
	for cycle, v := range queue {
		if shouldComputeSignalStrength(cycle + 1) {
			res += (x * (cycle + 1))
		}
		x += v
	}
	return
}

func drawScreen(queue []int) (res string) {
	x := 1
	for cycle, v := range queue {
		pos := cycle % 40
		if pos == x-1 || pos == x || pos == x+1 {
			res += "#"
		} else {
			res += "."
		}
		x += v
	}

	return
}

func dumpScreen(screen string) {
	nRows := 6
	nColumns := 40
	for i := 0; i < nRows; i++ {
		row := screen[i*nColumns : (i+1)*nColumns]
		fmt.Println(row)
	}
}

func main() {
	instructions := getInput()
	queue := enqueueInstructions(instructions)
	sum := sumSignalStrengths(queue)
	screen := drawScreen(queue)

	fmt.Println("1:", sum)
	fmt.Println("2:")
	dumpScreen(screen)
}
