package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput() (res [][]int) {
	const sep string = "\r\n"

	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := strings.TrimSpace(string(bytes))

	lines := strings.Split(content, sep)
	n := len(lines[0])
	res = make([][]int, n, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			res[i][j] = int(lines[i][j])
		}
	}
	return
}

func treesVisibleUp(grid [][]int, i, j int) (res int) {
	cont := true
	for k := i - 1; cont && k >= 0; k-- {
		res++
		cont = grid[k][j] < grid[i][j]
	}
	return
}

func treesVisibleDown(grid [][]int, i, j int) (res int) {
	cont := true
	for k := i + 1; cont && k < len(grid); k++ {
		res++
		cont = grid[k][j] < grid[i][j]
	}
	return
}

func treesVisibleRight(grid [][]int, i, j int) (res int) {
	cont := true
	for k := j + 1; cont && k < len(grid[i]); k++ {
		res++
		cont = grid[i][k] < grid[i][j]
	}
	return
}

func treesVisibleLeft(grid [][]int, i, j int) (res int) {
	cont := true
	for k := j - 1; cont && k >= 0; k-- {
		res++
		cont = grid[i][k] < grid[i][j]
	}
	return
}

func isVisible(grid [][]int, i, j int) bool {
	n := len(grid) - 1

	fromUp := treesVisibleUp(grid, i, j) == i
	if fromUp && i > 0 {
		fromUp = grid[i][j] > grid[0][j]
	}

	fromDown := treesVisibleDown(grid, i, j) == n-i
	if fromDown && i < n {
		fromDown = grid[i][j] > grid[n][j]
	}

	fromLeft := treesVisibleLeft(grid, i, j) == j
	if fromLeft && j > 0 {
		fromLeft = grid[i][j] > grid[i][0]
	}

	fromRight := treesVisibleRight(grid, i, j) == n-j
	if fromRight && j < n {
		fromRight = grid[i][j] > grid[i][n]
	}

	return fromUp || fromDown || fromLeft || fromRight
}

func countVisibleTrees(grid [][]int) (res int) {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isVisible(grid, i, j) {
				res++
			}
		}
	}
	return
}

func findHighestScore(grid [][]int) int {
	n := len(grid)
	highest := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			up := treesVisibleUp(grid, i, j)
			down := treesVisibleDown(grid, i, j)
			left := treesVisibleLeft(grid, i, j)
			right := treesVisibleRight(grid, i, j)
			score := up * down * left * right

			if score > highest {
				highest = score
			}
		}
	}
	return highest
}

func main() {
	grid := getInput()
	visibles := countVisibleTrees(grid)
	bestScore := findHighestScore(grid)

	fmt.Println("1:", visibles)
	fmt.Println("2:", bestScore)
}
