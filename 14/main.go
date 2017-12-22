package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/miketitus/adventcode2017/aoc"
)

var key = "hfdlxzhv"
var size = 128

func main() {
	hamming := 0
	for i := 0; i < size; i++ {
		hashInput := key + "-" + strconv.Itoa(i)
		kh := aoc.KnotHash(hashInput)
		for _, r := range kh {
			bin, err := strconv.ParseUint(string(r), 16, 8)
			if err != nil {
				fmt.Printf("Parse error: %v\n", err)
				os.Exit(1)

			}
			hamming += hammingDistance(uint8(bin))
		}
	}
	fmt.Printf("distance = %d\n", hamming)
}

func hammingDistance(x uint8) int {
	return int(table[x&0xF])
}

// from Hacker's Delight, about p. 70
var table = [16]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
}
