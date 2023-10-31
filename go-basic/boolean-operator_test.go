package main

import (
	"fmt"
	"testing"
)

func TestBooleanOperator(t *testing.T) {
	var ujian = 85
	var absensi = 81

	// var lulusUjian = ujian > 80
	// var lulusAbsensi = absensi > 80

	// lulus := lulusAbsensi && lulusUjian
	// fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 87)

}
