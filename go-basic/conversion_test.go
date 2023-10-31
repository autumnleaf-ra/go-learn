package main

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Rama"
	e := name[0]
	eString := string(e)

	fmt.Println(eString)

}
