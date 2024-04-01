package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstGreetings := greetings[:2]
	secondGreetings := greetings[1:4]
	thirdGreetings := greetings[3:]

	fmt.Println(greetings, firstGreetings, secondGreetings, thirdGreetings)
}
