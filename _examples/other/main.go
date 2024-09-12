package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{}

	for num := range 100 {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}

	slice = slice[5 : 5+4]
	slices.Reverse(slice)

	for _, num := range slice {
		fmt.Println(num)
	}
}
