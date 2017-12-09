package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = []string{
	"b inc 5 if a > 1",
	"a inc 1 if b < 5",
	"c dec -10 if a >= 1",
	"c inc -20 if c == 10",
}

type instruction struct {
	register  string
	offset    int
	condReg   string
	condOper  string
	condValue int
}

func (i *instruction) execute() {
	// evaluate conditional
	isTrue := false
	switch i.condOper {
	case "<":
		isTrue = registers[i.condReg] < i.condValue
	case "<=":
		isTrue = registers[i.condReg] <= i.condValue
	case ">":
		isTrue = registers[i.condReg] > i.condValue
	case ">=":
		isTrue = registers[i.condReg] >= i.condValue
	case "==":
		isTrue = registers[i.condReg] == i.condValue
	case "!=":
		isTrue = registers[i.condReg] != i.condValue
	default:
		fmt.Fprintf(os.Stderr, "unknown conditional '%s'\n", i.condOper)
		os.Exit(1)
	}
	if isTrue {
		registers[i.register] += i.offset
	} else {
		fmt.Printf("%s %s %d FALSE\n", i.condReg, i.condOper, i.condValue)
	}
}

var registers = make(map[string]int)

func main() {
	instructions := parseInput()
	//fmt.Println(instructions)
	for _, i := range instructions {
		i.execute()
	}
	//fmt.Println(registers)
	findLargestReg()
}

func parseInput() []instruction {
	lines := []instruction{}
	for _, line := range input {
		words := strings.Split(line, " ")
		r := words[0]
		createRegister(r)
		o, _ := strconv.ParseInt(words[2], 10, 0)
		if words[1] == "dec" {
			// negate offset
			o = -o
		}
		cReg := words[4]
		cOper := words[5]
		cValue, _ := strconv.ParseInt(words[6], 10, 0)
		i := instruction{register: r, offset: int(o), condReg: cReg, condOper: cOper, condValue: int(cValue)}
		lines = append(lines, i)
	}
	return lines
}

func createRegister(register string) {
	// create register if needed -- it may already exist
	_, ok := registers[register]
	if !ok {
		registers[register] = 0
	}
}

func findLargestReg() {
	maxKey := ""
	maxValue := 0
	for k, v := range registers {
		if maxValue < v {
			maxKey = k
			maxValue = v
		}
	}
	fmt.Printf("Largest value is %d in %s\n", maxValue, maxKey)
}
