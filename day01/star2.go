package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	var sums []int = []int{}
	sums = append(sums, 0)
	found := false
	for !found {
		file, _ := os.Open("input.txt")
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			number, _ := strconv.Atoi(line)
			sum += number
			if Contains(sums, sum) {
				fmt.Println(sum)
				found = true
				break
			} else {
				sums = append(sums, sum)
			}
		}
		file.Close()
	}
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
