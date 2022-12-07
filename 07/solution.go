package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func changeDir(currDir, moveDir string) string {
	if moveDir == ".." {
		parentDir := strings.LastIndex(currDir[:len(currDir)-1], "/")
		currDir = currDir[:parentDir+1]
	} else {
		if moveDir == "/" {
			currDir = moveDir
		} else {
			currDir += moveDir + "/"
		}
	}
	return currDir
}

func getInput() (res map[string]int) {
	const sep string = "\r\n"

	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := strings.TrimSpace(string(bytes))

	lines := strings.Split(content, sep)
	currDir := ""
	res = make(map[string]int)
	for i := 0; i < len(lines); i++ {
		prompt := strings.Split(lines[i], " ")

		if prompt[0] == "$" && prompt[1] == "cd" {
			currDir = changeDir(currDir, prompt[2])
		} else if prompt[0] != "$" && prompt[0] != "dir" {
			file, size := prompt[1], convertToInt(prompt[0])
			res[currDir+file] = size
		}
	}
	return
}

func updateDir(currDir string, fileSize int, dirSizes map[string]int) {
	parentDirs := strings.Split(currDir, "/")
	for i := len(parentDirs) - 1; i > 0; i-- {
		dir := strings.Join(parentDirs[:i], "/") + "/"
		dirSizes[dir] += fileSize
	}
}

func getDirSizes(fileSizes map[string]int) (res map[string]int) {
	res = make(map[string]int, 0)
	for file := range fileSizes {
		updateDir(file, fileSizes[file], res)
	}
	return
}

func sortBySizes(dirSizes map[string]int) (res []int) {
	dirs := make([]string, 0, len(dirSizes))
	for d := range dirSizes {
		dirs = append(dirs, d)
	}

	sort.SliceStable(dirs, func(i, j int) bool {
		return dirSizes[dirs[i]] < dirSizes[dirs[j]]
	})

	res = make([]int, 0, len(dirs))
	for _, d := range dirs {
		res = append(res, dirSizes[d])
	}
	return
}

func sumDirSizes(dirSizes []int) (res int) {
	const maxSize int = 100_000

	for i := 0; i < len(dirSizes) && dirSizes[i] <= maxSize; i++ {
		res += dirSizes[i]
	}
	return
}

func findSizeToDelete(dirSizes []int) (res int) {
	const maxRootSize int = 70_000_000
	const minFreeSize int = 30_000_000

	rootSize := dirSizes[len(dirSizes)-1]
	freeSize := maxRootSize - rootSize
	spaceToFree := minFreeSize - freeSize
	i := 0
	for dirSizes[i] < spaceToFree {
		i++
	}
	res = dirSizes[i]
	return
}

func main() {
	fileSizes := getInput()
	dirSizes := getDirSizes(fileSizes)
	sortedDirSizes := sortBySizes(dirSizes)
	sum := sumDirSizes(sortedDirSizes)
	sizeToDelete := findSizeToDelete(sortedDirSizes)

	fmt.Println("1:", sum)
	fmt.Println("2:", sizeToDelete)
}
