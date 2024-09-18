package main

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func main() {
	fmt.Println("seq.Range(3)")

	for num := range seq.Range(3) {
		fmt.Println(num)
	}

	fmt.Println("\nseq.Range(-2, -6)")

	for num := range seq.Range(-2, -6) {
		fmt.Println(num)
	}

	fmt.Println("\nseq.Range(2, 10, 3)")

	for num := range seq.Range(2, 10, 3) {
		fmt.Println(num)
	}

	fmt.Println("\nseq.Range(3, 7, 0)")

	for num := range seq.Range(3, 7, 0) {
		fmt.Println(num)
	}
}
