package main

import "fmt"

func main() {
	// var tajukWeb string = "Projek Sendiri"
	tajukWeb := "Projek Sendiri"

	const projectTickets int = 50 // value can't be change

	var remainingTickets int = 20
	// var negativeValue uint = -20 // cannot be signed negative value

	fmt. Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", projectTickets, remainingTickets, tajukWeb)

	fmt.Println("Welcome to ", tajukWeb, " booking application")
	fmt. Printf ("We have total of %v tickets and %v are still available. \n", projectTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	var userName string
	var userEmail string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your username")
	fmt.Scan(&userName)

	fmt.Println("Enter your email")
	fmt.Scan(&userEmail)

	fmt.Println("Enter your number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets booked. You will receive confirmation email at%v.\n", userName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, tajukWeb)
}