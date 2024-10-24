package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
	"strings"
)

func main() {
	accountDB := account.NewAccountsDB()
	fmt.Println("Password manager 2: Return of the J-son.")
exit:
	for {
		choice := getUserChoice()
		switch choice {
		case "1":
			createNewAccount(accountDB)
		case "2":
			findAccount(accountDB)
		case "3":
			deleteAccount(accountDB)
		case "4":
			break exit
		}
	}
}

func getUserChoice() string {
	var choice string
	fmt.Println("1. Create new account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Scanln(&choice)
	return choice
}

func createNewAccount(accountDB *account.AccountsDB) {
	newLogin := getData("Enter login: ")
	newPassword := getData("Enter password (or leave it blank to generate): ")
	newUrl := getData("Enter url: ")
	myAccount, err := account.NewAccount(newLogin, newPassword, newUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	accountDB.AddAccount(*myAccount)
}

func findAccount(accountDB *account.AccountsDB) {
	if len(accountDB.Accounts) == 0 {
		fmt.Println("Nothing to show.")
		fmt.Println("")
		return
	}
	var url string
	fmt.Print("Enter account`s url: ")
	fmt.Scanln(&url)
	for _, acc := range accountDB.Accounts {
		if strings.Contains(acc.Url, url) {
			fmt.Println(acc)
		}
	}
	fmt.Println("")
}

func deleteAccount(accountDB *account.AccountsDB) {
	var url string
	var choice string
	if len(accountDB.Accounts) == 0 {
		fmt.Println("Nothing to delete.")
		fmt.Println("")
		return
	}
	fmt.Print("Enter account`s url: ")
	fmt.Scanln(&url)
	for index, acc := range accountDB.Accounts {
		if strings.Contains(acc.Url, url) {
			fmt.Println("Found an account with login: \n", acc.Login)
			fmt.Print("Want to delete this account? (Y/N)")
			fmt.Scanln(&choice)
			if choice == "Y" || choice == "y" {
				accountDB.Accounts = append(accountDB.Accounts[:index], accountDB.Accounts[index+1:]...)
			}
		}
	}
	file, err := accountDB.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	files.WriteToFile(file, "data.json")
	fmt.Println("")
}

func getData(message string) string {
	fmt.Print(message)
	var data string
	fmt.Scanln(&data)
	return data
}
