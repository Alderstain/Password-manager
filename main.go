package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
)

func main() {
	fmt.Println("Password manager")
exit:
	for {
		choice := getUserChoice()
		switch choice {
		case "1":
			createNewAccount()
		case "2":
			findAccount()
		case "3":
			deleteAccount()
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

func createNewAccount() {
	newLogin := getData("Enter login: ")
	newPassword := getData("Enter password (or leave it blank to generate): ")
	newUrl := getData("Enter url: ")
	myAccount, err := account.NewAccount(newLogin, newPassword, newUrl)
	if err != nil {
		panic(err)
	}
	accountDB := account.NewAccountsDB()
	accountDB.AddAccount(*myAccount)
	file, err := accountDB.ToBytes()
	if err != nil {
		panic(err)
	}
	files.WriteToFile(file, "data.json")
}

func findAccount() {
}

func deleteAccount() {

}

func getData(message string) string {
	fmt.Print(message)
	var data string
	fmt.Scanln(&data)
	return data
}
