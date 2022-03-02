package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets int = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUser()

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	// var bookings = [50]string{}
	//slice is same as array but more dynamic
	// var bookings [50]string
	// var bookings []string

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTicket, firstName, lastName, email)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does'nt contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is invalid try again")
			// continue
		}

		city := "London"

		switch city {
		case "New work":
			//do new work code logic
		case "London", "Berlin":
			// do london and berlin coed logic
		default:
			// code block in case no of cases match
		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstName() []string {

	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket int, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTicket

	// create a map for a user
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thanks you %v %v for booking %v tickets. You will recvice confrimation email at %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

	fmt.Printf("The first names of booking are: %v \n", getFirstName())

}
