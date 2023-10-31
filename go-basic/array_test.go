package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var number [3]int

	number[0] = 5
	number[1] = 4
	number[2] = 2

	fmt.Println(number[0])

	var values = [3]int{
		90,
		95,
		90,
	}

	fmt.Println(values[0])

}
