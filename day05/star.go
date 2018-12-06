package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(len(line))
	somethingDone := true
	for somethingDone {
		somethingDone = false
		for i := 0; i < len(line)-1; i++ {
			diff := int(line[i]) - int(line[i+1])
			if diff == 32 || diff == -32 {
				line = line[0:i] + line[i+2:]
				somethingDone = true
				break
			}
		}
	}
	fmt.Println(len(line))
}
