package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lines = [][]int{}
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, countDuplicates(line))
	}
	file.Close()
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			diff := findDiff(lines[i], lines[j])
			if diff == 1 {
				for k := 0; k < len(lines[i]); k++ {
					fmt.Printf("%c", rune(lines[i][k]))
				}
				fmt.Println()
				for k := 0; k < len(lines[j]); k++ {
					fmt.Printf("%c", rune(lines[j][k]))
				}
				fmt.Println()
				fmt.Println(diff)
				break
			}

		}
	}
}

func countDuplicates(line string) []int {
	strArray := strings.Split(line, "")
	nums := []int{}
	for _, el := range strArray {
		nums = append(nums, int(el[0]))
	}
	return nums
}

func findDiff(arr1, arr2 []int) int {
	diff := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			diff++
		}
	}
	return diff
}
