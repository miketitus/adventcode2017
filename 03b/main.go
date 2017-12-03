package main

import (
	"fmt"
	"math"
)

const INPUT = 24
const (
	INIT  = iota
	RIGHT = iota
	LEFT  = iota
	UP    = iota
	DOWN  = iota
)

func main() {
	// calculate size of array square to hold data
	root := int(math.Ceil(math.Sqrt(float64(INPUT))))
	fmt.Println("root = ", root)
	// initialize array
	memory := make([][]int, root)
	for i := range memory {
		memory[i] = make([]int, root)
	}
	direction := INIT
	x := int(root / 2)
	y := int(root / 2)
	fmt.Println("x,y = ", x, y)
	for j := 1; j <= root*root; j++ {
		switch direction {
		case INIT:
			direction = RIGHT
		case RIGHT:
			y++
			if memory[x-1][y] == 0 {
				direction = UP
			}
		case LEFT:
			y--
			if memory[x+1][y] == 0 {
				direction = DOWN
			}
		case UP:
			x--
			if memory[x][y-1] == 0 {
				direction = LEFT
			}
		case DOWN:
			x++
			if memory[x][y+1] == 0 {
				direction = RIGHT
			}
		}
		memory[x][y] = adjacentSum(memory, x, y, root)
	}
	print2D(memory)
}

func adjacentSum(a [][]int, x int, y int, root int) int {
	xSet := []int{x}
	ySet := []int{y}
	if 0 != x {
		xSet = append(xSet, x-1)
	}
	if root != x {
		xSet = append(xSet, x+1)
	}
	if 0 != y {
		ySet = append(ySet, y-1)
	}
	if root != y {
		ySet = append(ySet, y+1)
	}
	fmt.Println("xSet = ", xSet)
	fmt.Println("ySet = ", ySet)
	return 1
}

func print2D(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Printf("\n")
	}
}
