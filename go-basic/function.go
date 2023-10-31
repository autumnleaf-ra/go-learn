package main

import "fmt"

func getFullName() (string, string) {
	return "Rama", "Dhan"
}

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

type Filter func(string) string
type Blacklist func(string) bool

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Youre blocked ", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

/* func blacklistAdmin(name string) bool {
	return name == "admin"
}

func blacklistRoot(name string) bool {
	return name == "root"
} */

/* blacklist := func(name string) bool {
	return name == "admin"
}

registerUser("admin", blacklist)
registerUser("eko", blacklist)

registerUser("asekkkkkkkkk", func(name string) bool {
	return name == "root"
}) */
