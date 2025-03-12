package main

import (
	"sync"
)

// Declare Variables
const conferenceTickets int = 50

var conference = "My Conference"
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

	// Get user info

	// validate user

	// if valid - do something

	// if not valid - complain why it isn't valid

}

func greetUsers() {

}

func getUserInput() {

}

func bookTicket() {

}

func sendTicket() {

}
