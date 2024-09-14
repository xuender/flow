package main

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func main() {
	for num := range seq.Range(-2, -6) {
		fmt.Println(num)
	}
}
