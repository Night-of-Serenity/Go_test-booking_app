package main

import (
	"booking-app/helper"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	fmt.Println("Welcome to our conference booking applications")
	fmt.Println("Get your tickets here to attend")

	for {
		firstName, lastName, email, userTickets := getUserInput()

		greetingUsers()

		// validate input
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := printFirstName()

			fmt.Printf("The first names of booking users are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			fmt.Println("Your input invalid, try again")
		}
	}
}

func greetingUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets :")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v  for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
