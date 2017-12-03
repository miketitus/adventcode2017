package main

import (
	"fmt"
	"math"
)

const INPUT = 25
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
	//fmt.Println("x,y = ", x, y)
	sum := 0
	xSet := []int{x}
	ySet := []int{y}
	if x != 0 {
		xSet = append(xSet, x-1)
	}
	if x != root-1 {
		xSet = append(xSet, x+1)
	}
	if y != 0 {
		ySet = append(ySet, y-1)
	}
	if y != root-1 {
		ySet = append(ySet, y+1)
	}
	//fmt.Println("xSet = ", xSet)
	//fmt.Println("ySet = ", ySet)
	for i := range xSet {
		for j := range ySet {
			xVal := xSet[i]
			yVal := ySet[j]
			sum += a[xVal][yVal]
			//fmt.Println("xVal, yVal, sum = ", xVal, yVal, sum)
		}
	}
	if sum == 0 {
		sum = 1
	}
	//fmt.Println("-------------")
	return sum
}

func print2D(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Printf("\n")
	}
}
