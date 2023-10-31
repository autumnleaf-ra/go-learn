package main

import (
	"fmt"
	"testing"
)

func TestDataMap(t *testing.T) {

	person := map[string]string{
		"name":    "Rama",
		"address": "Jakarta",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Rama"
	book["ups"] = "Salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
