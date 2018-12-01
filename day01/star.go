package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		sum += number
	}
	file.Close()
	fmt.Println(sum)
}
