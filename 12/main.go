package main

import (
	"fmt"
	"strings"
)

type program struct {
	id       string
	rawpipes []string
	pipes    []*program
	coolkid  bool
}

var input = []string{
	"0 <-> 2",
	"1 <-> 1",
	"2 <-> 0, 3, 4",
	"3 <-> 2, 4",
	"4 <-> 2, 3, 6",
	"5 <-> 6",
	"6 <-> 4, 5",
}
var programs = make(map[string]*program)

func main() {
	for _, line := range input {
		parseLine(line)
	}
	for _, prog := range programs {
		parsePipes(prog)
		fmt.Println(prog)
	}
}

func parseLine(line string) {
	words := strings.Split(line, " ")
	id := words[0]
	pipes := words[2:]
	for i, pipe := range pipes {
		pipes[i] = strings.TrimSuffix(pipe, ",")
	}
	p := program{id: id, rawpipes: pipes}
	programs[id] = &p
}

func parsePipes(p *program) {
	for _, remoteID := range p.rawpipes {
		p.pipes = append(p.pipes, programs[remoteID])
	}
}
