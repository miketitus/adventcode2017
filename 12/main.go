package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type program struct {
	id       int
	rawpipes []string
	pipes    []*program
	coolkid  bool
}

func (p *program) String() string {
	buf := bytes.NewBufferString(strconv.Itoa(p.id))
	buf.WriteRune(':')
	buf.WriteString(strconv.FormatBool(p.coolkid))
	return buf.String()
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
	}
	// must iterate map in key order
	keys := make([]string, 0)
	for k := range programs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		prog := programs[k]
		findCoolKids(prog)
		fmt.Printf("%d : %v\n", prog.id, prog.coolkid)
	}
}

func parseLine(line string) {
	words := strings.Split(line, " ")
	id, _ := strconv.Atoi(words[0])
	pipes := words[2:]
	for i, pipe := range pipes {
		pipes[i] = strings.TrimSuffix(pipe, ",")
	}
	p := program{id: id, rawpipes: pipes}
	programs[words[0]] = &p
}

func parsePipes(p *program) {
	for _, remoteID := range p.rawpipes {
		p.pipes = append(p.pipes, programs[remoteID])
	}
}

func findCoolKids(p *program) {
	if p.id == 0 {
		p.coolkid = true
	}
	if p.coolkid {
		// transmit coolness to connections
		for _, remote := range p.pipes {
			if p.id < remote.id {
				// process links in ascending order to avoid dupes
				remote.coolkid = true
			}
		}
	}
}
