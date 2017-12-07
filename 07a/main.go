package main

import (
	"fmt"
	"os"
	"strings"
)

var input1 = []string{
	"xhth (11)",
	"havc (22)",
	"ktlj (33)",
	"fwft (44) -> ktlj, cntj, xhth",
	"cntj (55)",
}

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
	parent   *program
	children []*program
}

func (p *program) addChild(c *program) {
	// fmt.Printf("adding child %p to parent %p\n", c, p)
	p.children = append(p.children, c)
}

var nodes = make(map[string]*program)

func main() {
	for _, line := range input {
		parseLine(line)
	}
	printNodes()
	// find root program (no parents)
	for _, p := range nodes {
		if p.parent == nil {
			fmt.Printf("%s is the root node\n", p.name)
		}
	}
}

func parseLine(line string) {
	words := strings.Split(line, " ")
	n := words[0]
	w := words[1][1:3]
	addProgram(n, w, nil, nil)
	if len(words) >= 4 {
		// has children
		parseChildren(words)
	}
}

func parseChildren(words []string) {
	parent, ok := nodes[words[0]]
	if !ok {
		fmt.Fprintf(os.Stderr, "parent %s not found", words[0])
		os.Exit(1)
	}
	for _, child := range words[3:] {
		n := strings.TrimSuffix(child, ",")
		c := addProgram(n, "", parent, nil)
		parent.addChild(c)
	}
}

func addProgram(n string, w string, p *program, c []*program) *program {
	prog, ok := nodes[n]
	if ok {
		// node exists, just update state
		if prog.weight == "" {
			prog.weight = w
		}
		if prog.parent == nil {
			prog.parent = p
		}
		if len(c) > 0 {
			prog.children = append(prog.children, c...)
		}
		/*fmt.Print("found ")
		printNode(prog)*/
	} else {
		// create new node
		prog = &program{name: n, weight: w, parent: p, children: c}
		nodes[n] = prog
		/*fmt.Print("created ")
		printNode(&newProg)*/
	}
	return prog
}

func printNodes() {
	for _, prog := range nodes {
		printNode(prog)
	}
}

func printNode(p *program) {
	prntName := ""
	if p.parent != nil {
		prntName = p.parent.name
	}
	fmt.Printf("%p:%s  (%s)  %p:%v  %v\n", p, p.name, p.weight, p.parent, prntName, p.children)
}
