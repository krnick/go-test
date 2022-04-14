// go mod init booking-app before you init the project

package main

import (
	"fmt"
)

func main() {

	var conference_name = "nick conf"
	const constant_ticket = 50
	conference_name = "123"
	fmt.Println("hello", conference_name, " world ")
	fmt.Println("hello", constant_ticket, " world ")

	fmt.Printf("hello world %v \n", conference_name)

	// variable

	var email string
	var ticket uint
	email = "sung@"
	// ticket = -1  this will overflow since the uint, it could not be the native number
	fmt.Print(email)
	fmt.Print(ticket)

	var user_name = ""
	fmt.Scan(&user_name)
	fmt.Println(user_name)

}
