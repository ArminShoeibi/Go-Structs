package main

import (
	"fmt"
)

type user struct {
	userName  string
	firstName string
	lastName  string
}

func main() {
	var user1 user
	user1.userName = "ArminShoeibi1"
	user1.firstName = "Armin1"
	user1.lastName = "Shoeibi1"

	fmt.Printf("%+v", user1)
	fmt.Println()
	fmt.Println(user1)

	var user2 user = user{
		userName:  "ArminShoeibi2",
		firstName: "Armin2",
		lastName:  "Shoeibi2",
	}
	fmt.Printf("%+v", user2)
	fmt.Println()
	fmt.Println(user2)

	user3 := user{
		"ArminShoeibi3",
		"Armin3",
		"Shoeibi3",
	}
	fmt.Printf("%+v", user3)
	fmt.Println()
	fmt.Println(user3)

	var user4 = user{
		"ArminShoeibi4",
		"Armin4",
		"Shoeibi4",
	}
	fmt.Printf("%+v", user4)
	fmt.Println()
	fmt.Println(user4)
}
