package main

import (
	"fmt"
)

func main(){

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainigTickets uint = conferenceTickets

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets out of %v tickets.\n", remainigTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("\nEnter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("\nEnter your email:")
	fmt.Scan(&email)

	fmt.Println("\nEnter the number of tickets you want:")
	fmt.Scan(&userTickets)

	fmt.Printf("\nThank You %v %v,\n%v tickets will be sent to your email: %v shortly.\n", firstName, lastName, userTickets, email)

}
