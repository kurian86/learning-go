package main

import (
	"fmt"
	"math"
)

func main() {
	i := 20
	var f float64 = float64(i)

	fmt.Println(i, f)

	const value = 50

	var i2 int = value
	var f2 float64 = value

	fmt.Println(i2, f2)

	var (
		b      byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	fmt.Println(b + 1, smallI + 1, bigI + 1)
}
