package main

import (
	"fmt"
	"strconv"
	"strings"
)

type layer struct {
	depth  int
	ranje  int
	period int
}

var firewall []layer

func main() {
	size := getFirewallSize(test)
	firewall = make([]layer, size, size)
	for _, line := range test {
		parseLine(line)
	}
	fmt.Println(firewall)
}

func getFirewallSize(f []string) int {
	l := f[len(test)-1]
	w := strings.Split(l, ":")
	s, _ := strconv.Atoi(w[0])
	return s + 1
}

func parseLine(line string) {
	words := strings.Split(line, ":")
	d, _ := strconv.Atoi(words[0])
	r, _ := strconv.Atoi(strings.TrimPrefix(words[1], " "))
	p := r*2 - 2
	l := layer{depth: d, ranje: r, period: p}
	firewall[d] = l
}

var test = []string{
	"0: 3",
	"1: 2",
	"4: 4",
	"6: 4",
}

var input = []string{
	"0: 3",
	"1: 2",
	"2: 4",
	"4: 6",
	"6: 4",
	"8: 6",
	"10: 5",
	"12: 6",
	"14: 8",
	"16: 8",
	"18: 8",
	"20: 6",
	"22: 12",
	"24: 8",
	"26: 8",
	"28: 10",
	"30: 9",
	"32: 12",
	"34: 8",
	"36: 12",
	"38: 12",
	"40: 12",
	"42: 14",
	"44: 14",
	"46: 12",
	"48: 12",
	"50: 12",
	"52: 12",
	"54: 14",
	"56: 12",
	"58: 14",
	"60: 14",
	"62: 14",
	"64: 14",
	"70: 10",
	"72: 14",
	"74: 14",
	"76: 14",
	"78: 14",
	"82: 14",
	"86: 17",
	"88: 18",
	"96: 26",
}
