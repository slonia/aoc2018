package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	file.Close()
	strArray := strings.Split(line, " ")
	numPlayers, _ := strconv.Atoi(strArray[0])
	last, _ := strconv.Atoi(strArray[6])
	fmt.Println(numPlayers, last)
	players := map[int]int{}
	for i := 0; i < numPlayers; i++ {
		players[i] = 0
	}
	fmt.Println(players)
	marbles := make([]int, 1)
	fmt.Println(marbles)
	// curIndex := 0
	curPlayer := 0
	// direction := 1
	left := 0
	right := 0
	for i := 0; i < last; i++ {
		if left == right {
			marbles = append(marbles, i+1)
			right++
		} else {
			marbles = append(marbles[:(left+1)], append([]int{i + 1}, marbles[(left+1):]...)...)
			left++
			right++
		}
		fmt.Println(marbles)
		curPlayer++
		if curPlayer > numPlayers {
			curPlayer = 0
		}
		fmt.Println(curPlayer)
	}
}
