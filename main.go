package main

import (
	"bookingapp/helper"
	"fmt"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, firstName, lastName, email)

			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {

				fmt.Println("Our Conversation is booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			//fmt.Println("Your input data is invalid, try again")
			//fmt.Printf("We have only %v tickets remaining. so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}

func greetUser() {

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names := string.Feilds(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
