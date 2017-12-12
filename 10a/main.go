package main

import (
	"fmt"
	"sort"
)

const size = 5

var list []int
var input = []int{3, 4}
var input1 = []int{
	83, 0, 193, 1, 254, 237, 187, 40, 88, 27, 2, 255, 149, 29, 42, 100,
}

func main() {
	// initialize list
	list = make([]int, size, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	// process input
	position := 0
	skipSize := 0
	for _, length := range input {
		fmt.Printf("---- length %d ----\n", length)
		var sublist []int
		if position+length > len(list) {
			// need to wrap
			end1 := list[position:]
			beg1 := list[:len(list)-position]
			sublist = append(end1, beg1...)
			fmt.Printf("end1:%v, beg1:%v, sub:%v\n", end1, beg1, sublist)
			// sort.Reverse gives wrong results, so reverse manually :(
			k := len(sublist) - 1
			for j := 0; j < k; j++ {
				z := sublist[j]
				sublist[j] = sublist[k]
				sublist[k] = z
				k--
			}
			end2 := sublist[position:]
			beg2 := sublist[:len(sublist)-position]
			fmt.Printf("end2:%v, beg2:%v, sub:%v\n", end2, beg2, sublist)
		} else {
			sublist = list[position:length]
			sort.Sort(sort.Reverse(sort.IntSlice(sublist)))
			remainder := list[(position + length):]
			list = append(sublist, remainder...)
		}
		position += length + skipSize
		skipSize++
		fmt.Printf("pos:%d, ss:%d, sublist:%v\n", position, skipSize, sublist)
		fmt.Printf("list:%v\n", list)
	}
}
