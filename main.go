package main

import (
	"fmt"
)

func main(){

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainigTickets uint = conferenceTickets
	var bookings [] string

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets out of %v tickets.\n", remainigTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	for remainigTickets > 0 {

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

		if userTickets > remainigTickets {
			fmt.Printf("\nWe only have %v tickets remaining, booking terminated.\n", remainigTickets)
			continue
		}

		remainigTickets = remainigTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName + " - " + fmt.Sprint(userTickets))

		fmt.Printf("\nThank You %v %v,\n%v tickets will be sent to your email: %v shortly.\n", firstName, lastName, userTickets, email)
		fmt.Printf("\nRemainig Tickets: %v\n", remainigTickets)

	}

	fmt.Println("No tickets availbale, exiting the program.")
	fmt.Printf("Bookings: %v\n", bookings)

}
