package main

import "fmt"

func displayCinema(seats [][]byte) {
	fmt.Println("\nCinema:")
	fmt.Print("  ")
	// display the columns
	for i := range seats[0] {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	// display cinema
	for i := range seats {
		fmt.Printf("%d ", i+1)
		for j := range seats[i] {
			fmt.Printf("%c ", seats[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func getNumber(prompt string) int {
	var n int
	fmt.Print(prompt)
	fmt.Scan(&n)
	return n
}

func calculateTotalIncome(rows, cols int) int {
	var profit, firstHalfRows, secondHalfRows int
	numberOfSeats := rows * cols
	if numberOfSeats <= 60 {
		profit = numberOfSeats * 10
	} else {
		firstHalfRows = rows / 2
		secondHalfRows = rows - firstHalfRows
		profit = firstHalfRows*cols*10 + secondHalfRows*cols*8
	}
	return profit
}

func initializeCinema(seats [][]byte) {
	for i := range seats {
		for j := range seats[i] {
			seats[i][j] = 'S'
		}
	}
}

func isValidCoordinate(coordinate int) bool {
	return coordinate >= 1 && coordinate <= 9
}

func calculateTicketPrice(row, col, rows, cols int) int {
	var ticketPrice int
	// check if both coordinates are valid
	if isValidCoordinate(row) && isValidCoordinate(col) {
		// check which row the seat is in
		totalSeats := rows * cols
		if totalSeats <= 60 {
			ticketPrice = 10
		} else {
			firstHalfRows := rows / 2
			if row <= firstHalfRows {
				ticketPrice = 10
			} else {
				ticketPrice = 8
			}
		}
	}
	return ticketPrice
}

func bookSeat(row, col int, seats [][]byte) {
	seats[row-1][col-1] = 'B'
}

func numberOfPurchasedTickets(seats [][]byte) int {
	var purchasedTickets int
	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == 'B' {
				purchasedTickets++
			}
		}
	}
	return purchasedTickets
}

func hasBeenPurchased(row, col int, seats [][]byte) bool {
	return seats[row][col] == 'B'
}

func processCommands() {
	// Get the rows and number of seats in each row of the cinema
	rows := getNumber("Enter the number of rows: ")
	cols := getNumber("Enter the number of seats in each row: ")
	// make a slice of runes
	var seats [][]byte
	// allocate slice
	seats = make([][]byte, rows)
	for i := range seats {
		seats[i] = make([]byte, cols)
	}
	// initialize cinema
	initializeCinema(seats)
	// to store currentIncome
	var currentIncome int
	// show menu
	choice := -1
	for {
		fmt.Println("\n1. Show the seats\n2. Buy a ticket\n3. Statistics\n0. Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			// display initial cinema
			displayCinema(seats)
		case 2:
			fmt.Println()
			// get the coordinates to calculate the ticket price for
			row := getNumber("Enter a row number: ")
			col := getNumber("Enter a seat number in that row: ")

			// validate coordinates
			for !isValidCoordinate(row) || !isValidCoordinate(col) {
				fmt.Println("Wrong input!")
				row = getNumber("Enter a row number: ")
				col = getNumber("Enter a seat number in that row: ")
			}

			for hasBeenPurchased(row-1, col-1, seats) {
				fmt.Print("\nThat ticket has already been purchased!\n\n")
				row = getNumber("Enter a row number: ")
				col = getNumber("Enter a seat number in that row: ")
			}
			// calculate ticket price
			ticketPrice := calculateTicketPrice(row, col, rows, cols)
			// display ticket price
			fmt.Printf("\nTicket price: $%d\n", ticketPrice)
			// book the seat to make it unavailable
			bookSeat(row, col, seats)
			// add ticketPrice to currentIncome
			currentIncome += ticketPrice
		case 3:
			purchasedTickets := numberOfPurchasedTickets(seats)
			fmt.Printf("\nNumber of purchased tickets: %d\n", purchasedTickets)
			percentage := (float64(purchasedTickets) / float64(rows*cols)) * 100.0
			fmt.Printf("Percentage: %.2f%%\n", percentage)
			fmt.Printf("Current income: $%d\n", currentIncome)
			fmt.Printf("Total income: $%d\n", calculateTotalIncome(rows, cols))
		case 0:
			// exit loop and function
			return
		}
	}
}

func main() {
	processCommands()
}
