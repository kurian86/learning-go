package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var total int = 0
	for {
		count, err := f.Read(make([]byte, 1024))
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		total += count
	}
	return total, nil
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
