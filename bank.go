package main

import (
	"fmt"
	"os"
	"strconv"
)

type Transaction struct {
	Type   string
	Amount float32
}

func saveBalance(balance float32) {
	balanceText := fmt.Sprintf("%.2f", balance)
	err := os.WriteFile("balance.txt",
		[]byte(balanceText),
		0644)
	if err != nil {
		fmt.Println("Could not save balance:", err)
	}
}

func readBalance() float32 {
	balanceData, err := os.ReadFile("balance.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return 0
		}
		fmt.Println("Could not read balance:", err)
		return 0
	}

	balance, err := strconv.ParseFloat(string(balanceData), 32)
	if err != nil {
		fmt.Println("Could not parse balance:", err)
		return 0
	}

	return float32(balance)
}

func main() {
	accountBalance := readBalance()
	transactions := []Transaction{}

	fmt.Println("Welcome to GO Bank!")

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println("5. Transaction History")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance: %.2f\n", accountBalance)

		case 2:
			fmt.Println("How much would you like to deposit?")
			var depositAmount float32
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("You cant deposit null/negative amount!")
			} else {
				accountBalance += depositAmount
				saveBalance(accountBalance)
				transactions = append(transactions, Transaction{
					Type:   "Deposit",
					Amount: depositAmount,
				})
				fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
			}

		case 3:
			fmt.Println("How much would you like to withdraw?")
			var withdrawAmount float32
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Please enter a positive amount.")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance. Please deposit money.")
			} else {
				accountBalance -= withdrawAmount
				saveBalance(accountBalance)
				transactions = append(transactions, Transaction{
					Type:   "Withdrawal",
					Amount: withdrawAmount,
				})
				fmt.Printf("Balance updated! New amount: %.2f\n", accountBalance)
			}

		case 4:
			fmt.Println("Thank you for using our services!")
			return

		case 5:
			if len(transactions) == 0 {
				fmt.Println("No transactions yet.")
			} else {
				fmt.Println("Transaction History:")
				for _, transaction := range transactions {
					fmt.Printf("%s: %.2f\n", transaction.Type, transaction.Amount)
				}
			}

		default:
			fmt.Println("Invalid choice. Please choose 1, 2, 3, 4, or 5.")
			continue
		}

		var continueChoice string
		fmt.Print("Would you like to make another transaction? (y/n): ")
		fmt.Scan(&continueChoice)
		if continueChoice != "y" && continueChoice != "Y" {
			fmt.Println("Thank you for using our services!")
			return
		}
	}
}
