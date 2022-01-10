package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	const conferenceTickets int = 50

	fmt.Printf("Welcome to %v booking app \n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend")

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email adress:")
	fmt.Scan(&email)

	fmt.Println("Enter numbers of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
