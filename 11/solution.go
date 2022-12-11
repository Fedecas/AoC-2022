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

func getMonkeyItems(line string) (res []int) {
	items := strings.Split(line, ": ")
	items = strings.Split(items[1], ", ")
	for j := 0; j < len(items); j++ {
		res = append(res, convertToInt(items[j]))
	}
	return
}

func getMonkeyOperation(line string) []int {
	code := -1
	if strings.Index(line, "+") != -1 {
		code = 0
	} else {
		if strings.Index(line, "old") == strings.LastIndex(line, "old") {
			code = 1
		} else {
			code = 2
		}
	}

	num := 0
	if code != 2 {
		aux := strings.Split(line, " ")
		num = convertToInt(aux[len(aux)-1])
	}

	return []int{code, num}
}

func getInputTest(numLine, trueLine, falseLine string) []int {
	aux := strings.Split(numLine, " ")
	num := convertToInt(aux[len(aux)-1])

	aux = strings.Split(trueLine, " ")
	target1 := convertToInt(aux[len(aux)-1])

	aux = strings.Split(falseLine, " ")
	target2 := convertToInt(aux[len(aux)-1])

	return []int{num, target1, target2}
}

func getInput() (res [][][]int) {
	const sep string = "\r\n"

	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := strings.TrimSpace(string(bytes))

	monkeys := strings.Split(content, sep+sep)
	for i := 0; i < len(monkeys); i++ {
		lines := strings.Split(monkeys[i], sep)
		items := getMonkeyItems(lines[1])
		operation := getMonkeyOperation(lines[2])
		test := getInputTest(lines[3], lines[4], lines[5])

		res = append(res, [][]int{items, operation, test})
	}
	return
}

func findTwoMaxValues(values []int) (v1, v2 int) {
	i1 := 0
	v1, v2 = 0, 0
	for i := 0; i < len(values); i++ {
		if values[i] > v1 {
			v1 = values[i]
			i1 = i
		}
	}

	for i := 0; i < len(values); i++ {
		if values[i] > v2 && i != i1 {
			v2 = values[i]
		}
	}

	return v1, v2
}

func nextWorryLevel(n, testProduct int, opArgs []int, worryReduction bool) int {
	op, num := opArgs[0], opArgs[1]
	worryLevel := 0
	switch op {
	case 0:
		worryLevel = n + num
	case 1:
		worryLevel = n * num
	default:
		worryLevel = n * n
	}

	if worryReduction {
		worryLevel %= testProduct
	} else {
		worryLevel /= 3
	}

	return worryLevel
}

func nextTarget(worryLevel int, testArgs []int) int {
	test, m1, m2 := testArgs[0], testArgs[1], testArgs[2]
	target := m2
	if worryLevel%test == 0 {
		target = m1
	}

	return target
}

func mostActiveMonkeysBusiness(monkeys [][][]int, rounds int, worryReduction bool) int {
	itemsPerMonkey := make([][]int, len(monkeys), len(monkeys))
	testProduct := 1
	for i, m := range monkeys {
		itemsPerMonkey[i] = make([]int, len(m[0]), len(m[0]))
		copy(itemsPerMonkey[i], m[0])
		testProduct *= m[2][0]
	}

	inspected := make([]int, len(monkeys), len(monkeys))
	for r := 0; r < rounds; r++ {
		for i, m := range monkeys {
			for j := 0; j < len(itemsPerMonkey[i]); j++ {
				item := itemsPerMonkey[i][j]
				worryLevel := nextWorryLevel(item, testProduct, m[1], worryReduction)
				target := nextTarget(worryLevel, m[2])
				itemsPerMonkey[target] = append(itemsPerMonkey[target], worryLevel)
				inspected[i]++
			}
			itemsPerMonkey[i] = nil
		}
	}

	v1, v2 := findTwoMaxValues(inspected)
	return v1 * v2
}

func main() {
	monkeys := getInput()
	monkeyBusiness := mostActiveMonkeysBusiness(monkeys, 20, false)
	monkeyBusinessReduced := mostActiveMonkeysBusiness(monkeys, 10000, true)

	fmt.Println("1:", monkeyBusiness)
	fmt.Println("2:", monkeyBusinessReduced)
}
