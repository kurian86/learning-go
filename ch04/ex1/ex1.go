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

	fmt.Println(result)
}
