package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var checksum, totalTwice, totalThree int
	var twice, three bool
	checksum = 1
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		twice, three = countDuplicates(line)
		if twice {
			totalTwice++
		}
		if three {
			totalThree++
		}
	}
	file.Close()
	checksum *= totalTwice * totalThree
	fmt.Println(checksum)
}

func countDuplicates(line string) (bool, bool) {
	twice := 1
	three := 1
	strArray := strings.Split(line, "")
	sort.Strings(strArray)
	current := 1
	for i := 0; i < len(strArray)-1; i++ {
		if strArray[i] == strArray[i+1] {
			current++
		} else {
			if current == 2 {
				twice++
			} else if current == 3 {
				three++
			}
			current = 1
		}
	}
	if current == 2 {
		twice++
	} else if current == 3 {
		three++
	}
	return twice > 1, three > 1
}
