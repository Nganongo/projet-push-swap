package main

import (
	"fmt"
	"os"

	"push-swap/src/lecture"
	"push-swap/src/pile"
	"push-swap/src/tri"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	nombres, err := lecture.Lire(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	p := pile.Nouvelles(nombres)
	tri.Trier(p)
	for _, op := range p.Operations {
		fmt.Println(op)
	}
}
