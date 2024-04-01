package main

import (
	"fmt"
)

func main() {
	var message = "Hi 👩 and 👨"
	var runes = []rune(message)

	fmt.Println(message[3])
	fmt.Println(string(runes[3]))
}
