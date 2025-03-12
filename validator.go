package main

import "strings"

func validateUser(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	//Assume valid is at least 2 characters per name
	boolName := len(firstName) >= 2 && len(lastName) >= 2
	boolEmail := strings.Contains(email, "@") // Leave for now
	boolTicket := userTickets > 0 && userTickets <= remainingTickets
	return boolName, boolEmail, boolTicket
}
