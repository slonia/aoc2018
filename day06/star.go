package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	id int
	x  int
	y  int
}

var points []*point
var maxX, maxY, minX, minY int

func main() {
	points = []*point{}
	minX = 100
	minY = 100
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		addPoint(line, i)
		i++
	}
	file.Close()
	sort.Slice(points, func(i, j int) bool {
		if points[i].x < points[j].x {
			return true
		} else if points[i].x == points[j].x {
			return points[i].y < points[j].y
		} else {
			return false
		}
	})
	board := make([][]int, maxX-minX)
	for x := minX; x < maxX; x++ {
		board[x-minX] = make([]int, maxY-minY)
		for y := minY; y < maxY; y++ {
			lastDistance := 10000
			for _, el := range points {
				dist := positiveDistance(x, el.x) + positiveDistance(y, el.y)
				if dist < lastDistance {
					board[x-minX][y-minY] = el.id
					lastDistance = dist
				}
			}
		}
	}
	fmt.Println(board)
}

func addPoint(line string, id int) {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1][1:])
	points = append(points, &point{id, x, y})
	if x > maxX {
		maxX = x
	}
	if y > maxY {
		maxY = y
	}
	if x < minX {
		minX = x
	}
	if y < minY {
		minY = y
	}
}

func positiveDistance(x1, x2 int) int {
	dist := x1 - x2
	if dist < 0 {
		return -dist
	} else {
		return dist
	}
}
