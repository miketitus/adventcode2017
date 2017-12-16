package main

import (
	"fmt"
	"time"
)

var pairs = 16
var inputA uint64 = 65
var factorA uint64 = 16807
var inputB uint64 = 8921
var factorB uint64 = 48271
var divisor uint64 = 2147483647
var mask uint64 = 0x0000FFFF

func main() {
	chA := make(chan uint64)
	chB := make(chan uint64)

	go func() {
		// generator A
		for i := 0; i < pairs; i++ {
			inputA = (inputA * factorA) % divisor
			if inputA%4 == 0 {
				fmt.Println("iA =", i)
				chA <- inputA & mask
			}
		}
		close(chA)
	}()

	go func() {
		// generator B
		for i := 0; i < pairs; i++ {
			inputB = (inputB * factorB) % divisor
			if inputB%8 == 0 {
				fmt.Println("iB =", i)
				chB <- inputB & mask
			}
		}
		close(chB)
	}()

	go func() {
		// judge
		var matches uint32
		for i := 0; i < pairs; i++ {
			bitsA := <-chA
			bitsB := <-chB
			fmt.Printf("bitsA = %d, bitsB = %d\n", bitsA, bitsB)
			if bitsA == bitsB {
				matches++
			}
		}
		fmt.Printf("total matches = %d\n", matches)
	}()

	time.Sleep(8 * time.Second)
}
