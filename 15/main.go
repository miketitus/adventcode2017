package main

import "fmt"

var inputA uint64 = 65
var factorA uint64 = 16807
var inputB uint64 = 8921
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
		//fmt.Printf("inputA = %x, input B = %x\n", inputA, inputB)
		//fmt.Printf("bitsA = %x, bitsB = %x\n", bitsA, bitsB)
		if bitsA == bitsB {
			matches++
			fmt.Printf("match at %d  %16b  %16b\n", i, bitsA, bitsB)
		}
	}
	fmt.Printf("total matches = %d\n", matches)
}
