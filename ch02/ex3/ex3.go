package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		b      byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	fmt.Println(b+1, smallI+1, bigI+1)
}
