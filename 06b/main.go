package main

import (
    "fmt"
)

type bank int

var banks1 = []bank{0, 2, 7, 0}
var banks = []bank{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}

func main() {
    fmt.Println(banks)
    cycles := 0
    loopStart := 0
    states := []string{}
    states = append(states, fmt.Sprint(banks))
    Outer:
    for {
        // generate new config
        maxBlock := mostBlocks(banks)        
        rebalance(banks, maxBlock)
        cycles++
        // test for previous state
        for i, s := range states {
            if s == fmt.Sprint(banks) {
                loopStart = i
                break Outer
            }
        }
        // store current state
        states = append(states, fmt.Sprint(banks))
        if cycles % 1000 == 0 {
            fmt.Println("cycles =", cycles)
        }
    }
    fmt.Println("cycles =", cycles)
    fmt.Println("loopStart =", loopStart)
    fmt.Println("loopSize =", cycles - loopStart)
    fmt.Println(banks)
}

func rebalance(b []bank, maxBlock int) {
    blocks := b[maxBlock]
    b[maxBlock] = 0
    for i := maxBlock+1 ; blocks > 0; i++ {
        if i == len(b) {
            i = 0
        }
        b[i]++
        blocks--
    }
}

func mostBlocks(b []bank) int {
    maxIndex := 0
    for i:= range b {
        if b[i] > b[maxIndex] {
            maxIndex = i
        }
    }
    return maxIndex
}
