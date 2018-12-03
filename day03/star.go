package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var startX, startY, sizeX, sizeY int
	set := map[int]map[int]int{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		startX, startY, sizeX, sizeY = extract(line)
		for i := 0; i < sizeX; i++ {
			for j := 0; j < sizeY; j++ {
				_, ok := set[startX+i]
				if !ok {
					set[startX+i] = map[int]int{}
				}
				_, ok = set[startX+i][startY+j]
				if ok {
					set[startX+i][startY+j]++
				} else {
					set[startX+i][startY+j] = 1
				}
			}
		}
	}
	file.Close()
	answer := 0
	for _, outerV := range set {
		for _, v := range outerV {
			if v > 1 {
				answer++
			}
		}
	}
	fmt.Println(answer)
}

func extract(line string) (int, int, int, int) {
	strArray := strings.Split(line, " ")
	startPoints := strings.Split(strArray[2], ",")
	startX, _ := strconv.Atoi(startPoints[0])
	startY, _ := strconv.Atoi(strings.TrimSuffix(startPoints[1], ":"))
	size := strings.Split(strArray[3], "x")
	sizeX, _ := strconv.Atoi(size[0])
	sizeY, _ := strconv.Atoi(size[1])
	return startX, startY, sizeX, sizeY
}
