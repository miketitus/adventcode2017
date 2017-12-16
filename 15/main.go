package main

import "fmt"

var inputA uint64 = 634
var factorA uint64 = 16807
var inputB uint64 = 301
var factorB uint64 = 48271
var divisor uint64 = 2147483647
var mask uint64 = 0x0000FFFF

func main() {
	var matches uint32
	for i := 0; i < 40000000; i++ {
		inputA = (inputA * factorA) % divisor
		inputB = (inputB * factorB) % divisor
		bitsA := inputA & mask
		bitsB := inputB & mask
		if bitsA == bitsB {
			matches++
		}
	}
	fmt.Printf("total matches = %d\n", matches)
}
