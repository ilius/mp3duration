package main

import (
	"fmt"
	"os"

	"github.com/ilius/mp3duration"
)

func main() {
	for _, fpath := range os.Args[1:] {
		duration, err := mp3duration.Calculate(fpath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("File %#v, Duration %v\n", fpath, duration)
	}
}
