package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get input, convert to int64
	fmt.Print("Enter captcha: ")
	reader := bufio.NewReader(os.Stdin)
	captcha, _ := reader.ReadString('\n')
	captcha = strings.TrimSuffix(captcha, "\n")
	i, err := strconv.ParseInt(captcha, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(i)
}
