package main

import "fmt"

type bank int

var banks = []bank{0, 2, 7, 0}
var banks1 = []bank{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}

func main() {
    fmt.Println(banks)
    cycles := 0
    states := [][]bank{}
    for {
        // store current config
        states = append(states, banks)
        // generate new config
        max := mostBlocks(banks)
        fmt.Println("max = ", max)
        // test for previous state
        cycles++
        if cycles > 0 {
            break
        }
    }
    fmt.Println(cycles)
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
