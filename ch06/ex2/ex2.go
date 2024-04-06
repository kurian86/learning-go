package main

import "fmt"

func UpdateSlice(a []string, b string) {
	a[len(a)-1] = b
	fmt.Println(a)
}

func GrowSlice(a []string, b string) {
	a = append(a, b)
	fmt.Println(a)
}

func main() {
	a := []string{"a", "b", "c", "d", "e"}
	UpdateSlice(a, "g")
	fmt.Println(a)

	b := []string{"a", "b", "c", "d", "e"}
	GrowSlice(b, "f")
	fmt.Println(b)
}
