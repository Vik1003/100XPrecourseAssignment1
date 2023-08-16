package main

import "fmt"

func main() {
	fmt.Println("HELLO WORLD ! +8")
	var conferenceName = "Go Conference"
	var conferenceTickets = 50
	var remainingTickets = 50

	//fmt.Println("Welcome to",conferenceName , "booking application")
	fmt.Printf("WELCOME TO '%v' BOOKING APPLICATION \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string //="Tom"
	// ask user for their name

	userName = "Tom"
	userTickets := 2
	fmt.Printf("User %v booked %v tickets for the conference", userName, userTickets)

}
