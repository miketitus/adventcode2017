package main

import (
	"fmt"
	"regexp"
)

var input1 = []string{
	"<abc>",
	"{}",
	"{{{}}}",
	"{{},{}}",
	"{{{},{},{{}}}}",
	"{<{},{},{{}}>}",
	"{<a>,<a>,<a>,<a>}",
	"{{<a>},{<a>},{<a>},{<a>}}",
	"{{<!>},{<!>},{<!>},{<a>}}",
}

var input = []string{
	"{z}",
	"{b},{c}",
	"{{d},{e}}",
	"{{{ghg}}}",
}

var negateRE = regexp.MustCompile("([\\!].)")
var garbageRE = regexp.MustCompile("<[^>]*>")
var contentRE = regexp.MustCompile("[^{}]")
var groupRE = regexp.MustCompile("{(?:{})*}")

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
	fmt.Println("countGroups: ", s)
	groups := 0
	if len(s) > 0 {
		subs := groupRE.FindAllStringSubmatch(s, -1)
		if len(subs) > 0 {
			fmt.Printf("FindAllStringSubmatch %v\n", subs)
			for _, sub := range subs {
				sub2 := sub[0] // strip nesting
				// remove outer {} and check inner string
				end := len(sub2) - 1
				groups += countGroups(sub2[1:end]) + 1
			}
		}
	}
	return groups
}

func main() {
	for _, line := range input {
		fmt.Println("--------------")
		//fmt.Println("orig: ", line)
		line = chopNegates(line)
		//fmt.Println("chopN: ", line)
		line = chopGarbage(line)
		//fmt.Println("chopG: ", line)
		line = chopContent(line)
		//fmt.Println("chopC: ", line)
		groups := countGroups(line)
		fmt.Println("groups =", groups)
	}
}
