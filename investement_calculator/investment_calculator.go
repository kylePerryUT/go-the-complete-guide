package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const acctBalanceFileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err1 := os.ReadFile(acctBalanceFileName)
	balanceTxt := string(data)
	balance, err2 := strconv.ParseFloat(balanceTxt, 64)
	if err1 != nil || err2 != nil {
		return 0.0, errors.New("there was an error getting the account balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	blanaceTxt := fmt.Sprint(balance)
	os.WriteFile(acctBalanceFileName, []byte(blanaceTxt), 0644)
}

func main() {
	var acctBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Welome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is ", acctBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmt float64
			fmt.Scan(&depositAmt)

			if depositAmt < 0 {
				fmt.Println("Invalid amount. Please enter an amount greater than 0")
				continue
			}

			acctBalance += depositAmt
			fmt.Println("Balance updated! New amount: ", acctBalance)
			writeBalanceToFile(acctBalance)
		case 3:
			fmt.Print("Amount to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount < 0 {
				fmt.Println("Invalid amount. Please enter an amount greater than 0")
				continue
			}

			if withdrawalAmount > acctBalance {
				fmt.Print("Invalid withdrawal amount. The withdrawal amount can't be greater than you current account balance.")
				continue
			}

			acctBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", acctBalance)
			writeBalanceToFile(acctBalance)
		default:
			fmt.Println("Goodbye")
			fmt.Println("Thanks for choosing Go Bank!")
			return
		}

	}
}
