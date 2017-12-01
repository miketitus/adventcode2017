package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter captcha: ")
	captcha, _ := reader.ReadString('\n')
	fmt.Println(captcha)
	digits := []byte(captcha)
	// validate input, convert from ascii to integers
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] < 48 || digits[i] > 57 {
			// error
		} else {
			digits[i] -= 48
		}
	}
	fmt.Println(digits)
}
