package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var input = []string{
	"{}",
	"{{{}}}",
	"{{},{}}",
	"{{{},{},{{}}}}",
	"{<a>,<a>,<a>,<a>}",
	"{{<ab>},{<ab>},{<ab>},{<ab>}}",
	"{{<!!>},{<!!>},{<!!>},{<!!>}}",
	"{{<a!>},{<a!>},{<a!>},{<ab>}}",
}

var input1 = []string{
	"{z}",
	"{b},{c}",
	"{{d},{e}}",
	"{{{ghg}}}",
}

var negateRE = regexp.MustCompile("([\\!].)")
var garbageRE = regexp.MustCompile("<[^>]*>")
var contentRE = regexp.MustCompile("[^{}]")

func chopNegates(s string) string {
	return negateRE.ReplaceAllString(s, "")
}

func chopGarbage(s string) string {
	return garbageRE.ReplaceAllString(s, "")
}

func chopContent(s string) string {
	return contentRE.ReplaceAllString(s, "")
}

func countGroups(s string) int {
	// build stack of counted groups
	st := stack.New()
	for len(s) > 0 {
		st.Push(strings.Count(s, "{}"))
		s = strings.Replace(s, "{}", "", -1)
	}
	// iterate stack and calculate score
	level := 0
	score := 0
	for st.Len() > 0 {
		level++
		score += st.Pop().(int) * level
	}
	fmt.Printf("level = %d, score = %d\n", level, score)
	return score
}

func main() {
	totalScore := 0
	for _, line := range input {
		line = chopNegates(line)
		line = chopGarbage(line)
		line = chopContent(line)
		totalScore += countGroups(line)
	}
	fmt.Println("totalScore =", totalScore)
}
