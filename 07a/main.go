package main

import (
	"fmt"
	"strings"
)

var input = []string{
	"pbga (66)",
	"xhth (57)",
	"ebii (61)",
	"havc (66)",
	"ktlj (57)",
	"fwft (72) -> ktlj, cntj, xhth",
	"qoyq (66)",
	"padx (45) -> pbga, havc, qoyq",
	"tknk (41) -> ugml, padx, fwft",
	"jptl (61)",
	"ugml (68) -> gyxo, ebii, jptl",
	"gyxo (61)",
	"cntj (57)",
}

type program struct {
	name     string
	weight   string
	children []program
}

var nodes = make(map[string]program)

func main() {
	for _, line := range input {
		p := parseLine(line)
		fmt.Println(p)
	}
	fmt.Println(nodes)
}

func parseLine(line string) program {
	childs := []program{}
	words := strings.Split(line, " ")
	if len(words) >= 4 {
		// has children
		childs = parseChildren(words)
	}
	n := words[0]
	w := words[1][1:3]
	p := program{name: n, weight: w, children: childs}
	nodes[n] = p
	return p
}

func parseChildren(words []string) []program {
	c := []program{}
	for _, child := range words[3:] {
		n := strings.TrimSuffix(child, ",")
		_, ok := nodes[n]
		if ok {
			fmt.Println("found ", n)
		}
	}
	return c
}
