package main

import (
	"fmt"
	"math/rand"
)

func main() {
	result := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		result = append(result, rand.Intn(100))
	}

	for _, v := range result {
		fmt.Printf("%d ", v)
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six")
		case v%2 == 0:
			fmt.Println("Two")
		case v%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never mind")
		}
	}
}
