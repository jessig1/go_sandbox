package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalance() float64 {
	// _ placeholder for a variable
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	// converts string to float64
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	// writes the balance amount to a text file, has a permission of 0644
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
func main() {
	var accountBalance = getBalance()
	fmt.Println("Welcome to the bank")
	//prevents script from ending after choice made
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice = == 1

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			// nested condition
			if depositAmount <= 0 {
				fmt.Println("Invalid amount must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New balance: ", accountBalance)
			writeBalance(accountBalance)
		case 3:
			fmt.Print("Withdrawl amount: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Must be greater than 0")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("cannot withdrawal more than account balance")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Println("New Balance: ", accountBalance)
			writeBalance(accountBalance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("thank you for using the bank")
			return
		}
		// if statement moved to switch statement
		// if choice == 1 {
		// 	fmt.Println("Your balance is: ", accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Deposit amount: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)
		// 	// nested condition
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid amount must be greater than 0")
		// 		//stops program
		// 		//return
		// 		// restarts loop
		// 		continue
		// 	}
		// 	// shorthand of accountBalance = accountBalance + depositAmount
		// 	accountBalance += depositAmount
		// 	fmt.Println("New balance: ", accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdrawl amount: ")
		// 	var withdrawlAmount float64
		// 	fmt.Scan(&withdrawlAmount)

		// 	if withdrawlAmount <= 0 {
		// 		fmt.Println("Must be greater than 0")
		// 		continue
		// 	}

		// 	if withdrawlAmount > accountBalance {
		// 		fmt.Println("cannot withdrawal more than account balance")
		// 		continue
		// 	}
		// 	accountBalance -= withdrawlAmount
		// 	fmt.Println("New Balance: ", accountBalance)
		// } else {
		// 	fmt.Println("Thank you!")
		// 	// breaks out of for loop
		// 	break
		// }

		// fmt.Println("Choose", choice)
	}
	//break in if statement
	//fmt.Println("Thanks for choosing the bank")
}
