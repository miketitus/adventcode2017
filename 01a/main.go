package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter captcha: ")
	captcha, _ := reader.ReadString('\n')
	captcha = strings.TrimSuffix(captcha, "\n")
	digits := []byte(captcha)
	// validate input, convert from ascii to integers
	for i := 0; i < len(digits); i++ {
		if digits[i] < 48 || digits[i] > 57 {
			fmt.Println("Invalid input")
			os.Exit(1)
		} else {
			digits[i] -= 48
		}
	}
	// compute captcha sum
	sum := 0
	for i := 0; i < len(digits); i++ {
		nextI := i + 1
		if nextI == len(digits) {
			nextI = 0
		}
		if digits[i] == digits[nextI] {
			sum += int(digits[nextI])
		}
	}
	fmt.Println(sum)
}
