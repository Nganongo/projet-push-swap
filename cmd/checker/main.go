package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/src/lecture"
	"push-swap/src/pile"
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
	t := pile.Nouvelles(nombres)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ligne := scanner.Text()
		if !t.Executer(ligne) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}
	if t.EstTriee() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
