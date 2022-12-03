package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const sep = "\r\n"
const types = ".abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
		if len(lines[i]) > 1 {
			res = append(res, lines[i])
		}
	}
	return
}

func getPriority(itemType byte) int {
	return strings.IndexByte(types, itemType)
}

func sumPriorities(rucksacks []string) (res int) {
	for i := 0; i < len(rucksacks); i++ {
		compSize := len(rucksacks[i]) / 2
		comp1, comp2 := rucksacks[i][:compSize], rucksacks[i][compSize:]
		itemTypeIndex := strings.IndexAny(comp1, comp2)
		res += getPriority(comp1[itemTypeIndex])
	}
	return
}

func getBadge(e1 string, e2 string, e3 string) (res byte) {
	isBadge := false
	start := 0
	itemTypeIndex := -1
	for !isBadge {
		start += itemTypeIndex + 1
		itemTypeIndex = strings.IndexAny(e1[start:], e2)
		isBadge = strings.IndexByte(e3, e1[start+itemTypeIndex]) != -1
	}
	res = e1[start+itemTypeIndex]
	return
}

func sumBadgesPriorities(rucksacks []string) (res int) {
	for i := 0; i < len(rucksacks); i += 3 {
		elf1, elf2, elf3 := rucksacks[i], rucksacks[i+1], rucksacks[i+2]
		badge := getBadge(elf1, elf2, elf3)
		res += getPriority(badge)
	}
	return
}

func main() {
	rucksacks := getInput()
	total := sumPriorities(rucksacks)
	totalBadges := sumBadgesPriorities(rucksacks)

	fmt.Println("1:", total)
	fmt.Println("2:", totalBadges)
}
