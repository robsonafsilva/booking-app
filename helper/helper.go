package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, raminingTickets uint) (bool, bool, bool) {
	//validate the user input
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= raminingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
