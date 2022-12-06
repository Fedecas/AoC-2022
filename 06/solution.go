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

func getInput() (res string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)
	content := string(bytes)
	res = strings.TrimSpace(content)
	return
}

func isMarker(data string) (res bool) {
	res = true
	for i := 0; i < len(data); i++ {
		if strings.Count(data, string(data[i])) > 1 {
			res = false
		}
	}
	return
}

func findStartMarker(data string, length int) (res int) {
	found := false
	start := 0
	for !found {
		found = isMarker(data[start : start+length])
		start++
	}
	res = start + length - 1
	return
}

func main() {
	dataStream := getInput()
	markerStart := findStartMarker(dataStream, 4)
	markerMessage := findStartMarker(dataStream, 14)

	fmt.Println("1:", markerStart)
	fmt.Println("2:", markerMessage)
}
