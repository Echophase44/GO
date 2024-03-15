package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// var bookings = [50] string {"Nick", "Isla", "Beck"}
	// The Array version:
	// var bookings [50]string

	//The Slice version:
	var bookings []string

	for {
		var firstName string
		var lastName string
		var userEmail string
		var userTickets uint

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
			// Adding to specific postion in the Array: bookings[0] = firstName + " " + lastName
			// Adding to the end of the Slice:
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			// the _ is a blank identifier, for a variable that won't be used, in this case, "index"
			for _, user := range bookings {
				var names = strings.Fields(user)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Bookings made by: %v\n", firstNames)

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
