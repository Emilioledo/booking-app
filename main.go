package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking app \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
