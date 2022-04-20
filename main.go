// go mod init booking-app before you init the project

package main

import (
	"fmt"
	"go-test/helper"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	var ticket uint

	fmt.Println("Hello world, welcome to ticket order system")
	fmt.Println("Please type the number of the ticket you want to book")

	fmt.Scan(&ticket)

	fmt.Printf("your ticket number is %v \n", ticket)

	helper.Say_hello()

	//send_mail()
	wg.Add(1)
	go send_mail() // but the main thread won't wait for it, so it need to use three function in waitgroud
	wg.Wait()

}

func send_mail() {
	// which take lots of time here
	// we need to do it concurrency, it can simply put 'go' keyword before call this function
	fmt.Println("Sending your email right now")
	time.Sleep(10 * time.Second)
	fmt.Println("Done for sending mail")
	wg.Done() // it will decrease 1 the number of wg.add(1)
}
