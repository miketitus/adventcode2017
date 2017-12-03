package main

import (
	"fmt"
	"math"
)

//const input = 265149
const INPUT = 48
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
	memory, inputX, inputY := initialize(root)
	fmt.Println(inputX, inputY)
	print2D(memory)
	navigate(memory, inputX, inputY)
}

func initialize(root int) ([][]int, int, int) {
	// initialize array
	var inputX, inputY int
	arry := make([][]int, root)
	for i := range arry {
		arry[i] = make([]int, root)
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
			if arry[x-1][y] == 0 {
				direction = UP
			}
		case LEFT:
			y--
			if arry[x+1][y] == 0 {
				direction = DOWN
			}
		case UP:
			x--
			if arry[x][y-1] == 0 {
				direction = LEFT
			}
		case DOWN:
			x++
			if arry[x][y+1] == 0 {
				direction = RIGHT
			}
		}
		arry[x][y] = j
		if INPUT == j {
			inputX = x
			inputY = y
		}
	}
	return arry, inputX, inputY
}

func navigate(memory [][]int, x int, y int) {
	fmt.Println("x,y = ", x, y)
	steps := 0
	var direction int
	// move to center row
	if x == 0 {
		direction = DOWN
	} else if memory[x][y] > memory[x-1][y] {
		direction = UP
	} else {
		direction = DOWN
	}
	for memory[x][y] != 1 {
		if UP == direction {
			x--
		} else {
			x++
		}
		steps++
	}
	fmt.Println("x,y = ", x, y)
	fmt.Println("steps = ", steps)
}

func print2D(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%5d", a[i][j])
		}
		fmt.Printf("\n")
	}
}
