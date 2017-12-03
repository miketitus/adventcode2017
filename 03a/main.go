package main

import (
	"fmt"
	"math"
)

//const input = 265149
const input = 49
const (
	INIT  = iota
	RIGHT = iota
	LEFT  = iota
	UP    = iota
	DOWN  = iota
)

func main() {
	// calculate size of array square to hold data
	root := int(math.Ceil(math.Sqrt(float64(input))))
	fmt.Println("root = ", root)
	// initialize memory array
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
		memory[x][y] = j
	}
	print2D(memory)
}

func print2D(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Printf("\n")
	}
}
