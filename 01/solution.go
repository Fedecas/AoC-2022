package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

const sep = "\r\n"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() (res [][]int) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := string(bytes)
	elves := strings.Split(content, sep+sep)

	n := len(elves)
	for i := 0; i < n; i++ {
		values := strings.Split(elves[i], sep)

		calories := make([]int, 0)
		for j := 0; j < len(values); j++ {
			if values[j] != "" {
				v, err := strconv.Atoi(values[j])
				check(err)
				calories = append(calories, v)
			}
		}
		res = append(res, calories)
	}

	return
}

func sumCalories(elves [][]int) (res []int) {
	for i := 0; i < len(elves); i++ {
		sum := 0
		for j := 0; j < len(elves[i]); j++ {
			sum += elves[i][j]
		}
		res = append(res, sum)
	}

	return
}

func sumMaxN(calories []int, n int) (res int) {
	sort.Ints(calories)

	k := len(calories)
	for i := k - 1; i >= k-n; i-- {
		res += calories[i]
	}

	return
}

func main() {
	elves := getInput()
	totalCalories := sumCalories(elves)

	a := sumMaxN(totalCalories, 1)
	fmt.Println("1:", a)

	b := sumMaxN(totalCalories, 3)
	fmt.Println("2:", b)
}
