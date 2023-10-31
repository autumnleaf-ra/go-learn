package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	/* months[5] = "apa"
	fmt.Println(slice1)

	Change value months
	*/

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Rama") // membuat slice data baru ketika data sudah penuh
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rama"
	newSlice[1] = "Ramadhan"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
