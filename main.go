package main

import (
	"fmt"
	"sync"
	"time"
)

// Declare Variables
const conferenceTickets int = 50

var conferenceName = "My Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	ticketCount uint
}

var wg = sync.WaitGroup{}

func main() {

	// welcome users
	greetUsers()
	// Get user info
	firstName, lastName, email, userTickets := getUserInput()
	// validate user
	boolName, boolEmail, boolTicketNum := validateUser(firstName, lastName, email, userTickets)

	// if valid - do something
	if boolName && boolEmail && boolTicketNum {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {
		// if not valid - complain why it isn't valid
		if !boolName {
			fmt.Println("bad name")
		}
		if !boolEmail {
			fmt.Println("bad email")
		}
		if !boolTicketNum {
			fmt.Println("bad ticket num")
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here")

}

func getName() []string {
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

	fmt.Println("First Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Email: ")
	fmt.Scan(&email)

	fmt.Println("# of tickets?: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		ticketCount: userTickets,
	}

	bookings = append(bookings, userData)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var tkt = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("ticket:\n %v \nto email address %v\n", tkt, email)
	wg.Done()

}
