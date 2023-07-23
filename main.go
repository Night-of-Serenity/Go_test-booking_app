package main

import "fmt"

func main() {
	fmt.Println("Welcome to our conference booking applications")
	fmt.Println("Get your tickets here to attend")

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n",conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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
	fmt.Scan(&userTickets )


	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v  for booking %v tickets. You will receive a confirmation email at %v \n",firstName,lastName,userTickets,email)
	
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)

}	