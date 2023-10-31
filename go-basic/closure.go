package main

import "fmt"

func closure() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	increment()

	fmt.Println(counter)
}
