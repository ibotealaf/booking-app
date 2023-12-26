package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func stats(customers []string) string {
	namesOfCustomers := ""
	customersFirstNames := []string{}

	for _, value := range customers {
		firstname := strings.Fields(value)[0]
		present := slices.Contains(customersFirstNames, firstname)

		if !present {
			customersFirstNames = append(customersFirstNames, firstname)
		}
	}

	filteredFirstNames := len(customersFirstNames)
	if filteredFirstNames >= 1 {
		namesOfCustomers += customersFirstNames[0]

		for i := 1; i < filteredFirstNames; i++ {
			if i == filteredFirstNames-1 {
				namesOfCustomers += " & "
			} else {
				namesOfCustomers += ", "
			}
			namesOfCustomers += customersFirstNames[i]
		}
	} else {
		return "No customers yet"
	}

	statement := "We had a total of " + strconv.Itoa(filteredFirstNames) + " customer(s) with the following names: " + namesOfCustomers
	return statement
}

func main() {
	const bookingName = "2023 EOY Party"
	const totalTickets = 40
	const ticketCost float32 = 99.99
	var remainingTickets uint8 = 40
	customers := []string{}

	fmt.Println("#############################################")
	fmt.Printf("Welcome to the %v booking page.\n", bookingName)
	fmt.Printf("We have total of %v up for grabs and %v remaining tickets\n", totalTickets, remainingTickets)
	fmt.Println("#############################################")

	fmt.Println("Proceed to buy ticket")

	for remainingTickets > 0 {
		var firstName string
		var lastName string
		var email string
		var qtyPurchase uint8
		var confirm string

		fmt.Print("\nEnter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter email address: ")
		fmt.Scan(&email)
		fmt.Print("How many tickets do you need: ")
		fmt.Scan(&qtyPurchase)

		if qtyPurchase > remainingTickets {
			fmt.Printf("Sorry we have only %v tickets left", remainingTickets)
			continue
		} else if qtyPurchase == remainingTickets {
			fmt.Println("LUCKY YOU :) - You're getting our last set of tickets")
		}
		qtyPurchaseCost := ticketCost * float32(qtyPurchase)

		fmt.Printf("Hey %v %v the cost for %v ticket is $%v\n", firstName, lastName, qtyPurchase, qtyPurchaseCost)
		fmt.Println("Proceed to make payment...")
		fmt.Println("Press yes/y to confirm and no/n to cancel")
		fmt.Scan(&confirm)

		confirm = strings.ToLower(confirm)
		if confirm == "y" || confirm == "yes" {
			fmt.Printf("Thanks for your purchase %v. Sent invoice to %v\n", firstName, email)
			remainingTickets -= qtyPurchase
			customers = append(customers, firstName+" "+lastName)
		} else {
			fmt.Println("Operation terminated!")
		}

		fmt.Printf("%v tickets already sold and we have %v tickets left. \n", (totalTickets - remainingTickets), remainingTickets)
	}
	fmt.Println("###########")
	fmt.Println("SOLD OUT :)")
	fmt.Println("###########")
	fmt.Println(stats(customers))

}
