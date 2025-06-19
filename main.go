package main

import (
	"fmt"
)

func main(){

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainigTickets = conferenceTickets

	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets out of %v tickets.\n", remainigTickets, conferenceTickets)
	fmt.Println("Get your tickets here to attend")
}
