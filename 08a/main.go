package main

import (
	"fmt"
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
	condition string
}

var registers []int

func main() {
	registers = make([]int, len(input))
	instructions := parseInput()
	fmt.Println(instructions)
}

func parseInput() []instruction {
	lines := []instruction{}
	for _, line := range input {
		words := strings.Split(line, " ")
		r := words[0]
		o, _ := strconv.ParseInt(words[2], 10, 0)
		if words[1] == "dec" {
			// negate offset
			o = -o
		}
		c := strings.Join(words[4:], " ")
		i := instruction{register: r, offset: int(o), condition: c}
		lines = append(lines, i)
	}
	return lines
}
