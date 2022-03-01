package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	// var bookings = [50]string{}

	//slice is same as array but more dynamic

	// var bookings [50]string
	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTicket int
	// ask the user for their name

	// accept user input
	fmt.Println("Please enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name :")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email :")
	fmt.Scan(&email)

	fmt.Println("Please enter number of ticket :")
	fmt.Scan(&userTicket)

	remainingTickets = remainingTickets - userTicket
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	// fmt.Printf("The whole array: %v \n", bookings)
	// fmt.Printf("The first value: %v \n", bookings[0])
	// fmt.Printf("Array type: %T \n", bookings)
	// fmt.Printf("Array length: %v \n", len(bookings))

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("slice type: %T \n", bookings)
	fmt.Printf("slice length: %v \n", len(bookings))

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will recvice confrimation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}
