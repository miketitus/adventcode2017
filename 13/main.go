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
var size int

func main() {
	// build firewall
	size = getFirewallSize(input)
	firewall = make([]layer, size, size)
	for _, line := range input {
		parseLine(line)
	}
	// send packet through firewall
	delay := 0
	var detected bool
	severity := -1
	for true {
		detected, severity = traverseFirewall(delay)
		if !detected {
			break
		}
		delay++
	}
	fmt.Printf("delay = %d, severity = %d\n", delay, severity)
}

func getFirewallSize(f []string) int {
	l := f[len(f)-1]
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

func traverseFirewall(delay int) (bool, int) {
	detected := false
	severity := 0
	for i := 0; i < size; i++ {
		lyr := firewall[i]
		if lyr.depth == i {
			// layer exists (not empty)
			if (i+delay)%lyr.period == 0 {
				// packet detected
				detected = true
				severity += lyr.ranje * i
			}
		}
	}
	return detected, severity
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
