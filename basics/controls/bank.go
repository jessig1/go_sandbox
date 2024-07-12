package main

import (
	//imports from fileops package we created
	"examplecom/bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	//uses GetFloat function from fileops file
	var accountBalance, err = fileops.GetFloat(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----")
		// stops execution with error message
		//panic("Can't continue")
	}
	fmt.Println("Welcome to the bank")
	fmt.Println("reach us at ", randomdata.PhoneNumber())
	//prevents script from ending after choice made
	for {
		presentOptions()
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
			fileops.WriteFloat(accountBalance, accountBalanceFile)
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
			fileops.WriteFloat(accountBalance, accountBalanceFile)
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

//go get installs dependent packages, can also specify path to library for new package
