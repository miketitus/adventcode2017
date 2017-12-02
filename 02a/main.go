package main

import "fmt"

var data = [][]uint16{
	{5, 1, 9, 5},
	{7, 5, 3},
	{2, 4, 6, 8},
}

func main() {
	sum := uint16(0)
	for i := 0; i < len(data); i++ {
		rowMin := data[i][0]
		rowMax := data[i][0]
		for j := 0; j < len(data[i]); j++ {
			if rowMin > data[i][j] {
				rowMin = data[i][j]
			} else if rowMax < data[i][j] {
				rowMax = data[i][j]
			}
		}
		sum += rowMax - rowMin
	}
	fmt.Println(sum)
}
