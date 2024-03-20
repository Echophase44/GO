package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"time"
)

var firstName string
var lastName string
var userEmail string
var userTickets uint

var bookings = make([]SUserData, 0)

type SUserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	helper.PrintThisLine()

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// bookingss := []string{}

	for {

		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email name:")
		fmt.Scan(&userEmail)

		fmt.Println("Please enter the number of tickets you wish to purchase:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			var userData = SUserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           userEmail,
				numberOfTickets: userTickets,
			}

			// Adding to the end of the Slice:
			bookings = append(bookings, userData)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			go sendTicket(userTickets, firstName, lastName, userEmail)

			var allFirstNames = printFirstNames()

			fmt.Printf("The first names of booking are: %v\n", allFirstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email you entered is missing the @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers(user string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our conference %v!\n", user)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, user := range bookings {
		firstNames = append(firstNames, user.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#############")
}
