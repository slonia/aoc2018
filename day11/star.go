package main

import "fmt"

const serial = 9798

func main() {
	maxPower := -1000
	var maxX, maxY int
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
		for j := range grid[i] {
			grid[i][j] = powerFor(i, j)
		}
	}
	for i := 0; i <= 297; i++ {
		for j := 0; j <= 297; j++ {
			power := grid[i][j] + grid[i][j+1] + grid[i][j+2] +
				grid[i+1][j] + grid[i+1][j+1] + grid[i+1][j+2] +
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if power > maxPower {
				maxPower = power
				maxX = i
				maxY = j
			}
		}
	}
	fmt.Println(maxPower, maxX, maxY)
}

func powerFor(x, y int) int {
	rackId := x + 10
	level := rackId * y
	level += serial
	level *= rackId
	if level > 99 {
		level = (level - (level/1000)*1000 - (level % 100)) / 100
	} else {
		level = 0
	}
	return level - 5
}

// 1234
// 1234 % 100 = 34
// 1234 / 1000 = 1

// (1234 - 1*1000 - 34) / 100
