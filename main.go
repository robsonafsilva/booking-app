package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var raminingTickets uint = 50

//create a list of map the 0 indicate the inicial size of
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserinput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, raminingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// call function get firts names store the value in a var and show
		// the output in a printf statement
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if raminingTickets == 0 {
			// end program
			fmt.Println("Our conferecnce is booked out. Come back next year.")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("First or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered is invalid!")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of ticket is invalid")
		}
	}
	wg.Wait()

}

func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availiable.\n", conferenceTickets, raminingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames() []string {
	// []type{} declare a slice - slices are a type of array with dynamic size
	firstNames := []string{}
	for _, booking := range bookings { // _ (underline is used in GO to set a variable, that the content is not important )
		// in GO variables need to be used if you don't need of the content in the variable use _
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserinput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	raminingTickets = raminingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// the %v is used to format text output there is a lot of %keys to format text outputs
	fmt.Printf("Thank you %v %v for booking %v. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", raminingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	//Sprintf store a formated string in a variable
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	println("################")
	wg.Done()
}
