package main

import (
	"fmt"
	"sync"
)

var pairs = 5000000
var inputA uint64 = 634
var factorA uint64 = 16807
var inputB uint64 = 301
var factorB uint64 = 48271
var divisor uint64 = 2147483647
var mask uint64 = 0x0000FFFF

func main() {
	chA := make(chan uint64)
	chB := make(chan uint64)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// generator A
		for true {
			inputA = (inputA * factorA) % divisor
			if inputA%4 == 0 {
				chA <- inputA & mask
			}
		}
	}()

	go func() {
		// generator B
		for true {
			inputB = (inputB * factorB) % divisor
			if inputB%8 == 0 {
				chB <- inputB & mask
			}
		}
	}()

	go func() {
		// judge
		var matches uint32
		for i := 0; i < pairs; i++ {
			bitsA := <-chA
			bitsB := <-chB
			if bitsA == bitsB {
				matches++
			}
		}
		fmt.Printf("total matches = %d\n", matches)
		wg.Done()
	}()

	wg.Wait()
}
