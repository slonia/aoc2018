package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type guard struct {
	id        int
	total     int
	lastStart int
	lastEnd   int
	minutes   map[int]int
}

var currentGuard int
var guards map[int]*guard

func main() {
	guards = map[int]*guard{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	file.Close()
	sort.Strings(lines)
	for _, line := range lines {
		process(line)
	}
	maxId := 0
	maxTime := 0
	for _, el := range guards {
		if el.total > maxTime {
			maxTime = el.total
			maxId = el.id
		}
	}
	maxMinutesId := 0
	maxMinutesVal := 0
	for k, v := range guards[maxId].minutes {
		if v > maxMinutesVal {
			maxMinutesVal = v
			maxMinutesId = k
		}
	}
	fmt.Println(maxId, maxMinutesId, maxId*maxMinutesId)
}

func process(line string) {
	time := line[:17]
	action := line[19:]
	if strings.HasPrefix(action, "Guard") {
		processGuard(action)
	} else if strings.HasPrefix(action, "falls") {
		setStart(time)
	} else {
		setEnd(time)
	}
}

func processGuard(action string) {
	strArray := strings.Split(action, " ")
	currentGuard, _ = strconv.Atoi(strArray[1][1:])
	_, ok := guards[currentGuard]
	if !ok {
		guards[currentGuard] = &guard{currentGuard, 0, 0, 0, map[int]int{}}
	}
}

func setStart(time string) {
	minutes := extractTime(time)
	guards[currentGuard].lastStart = minutes
}

func setEnd(time string) {
	minutes := extractTime(time)
	guards[currentGuard].lastEnd = minutes
	guards[currentGuard].total += guards[currentGuard].lastEnd - guards[currentGuard].lastStart
	for i := guards[currentGuard].lastStart; i < guards[currentGuard].lastEnd; i++ {
		_, ok := guards[currentGuard].minutes[i]
		if ok {
			guards[currentGuard].minutes[i]++
		} else {
			guards[currentGuard].minutes[i] = 1
		}
	}
}

func extractTime(time string) int {
	array := strings.Split(time, ":")
	minutes, _ := strconv.Atoi(array[1][:2])
	return minutes
}
