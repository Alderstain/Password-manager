package main

import (
	"fmt"
	"passwordManager/account"
)

func main() {
	newLogin := getData("Введите логин: ")
	newPassword := getData("Введите пароль: ")
	newUrl := getData("Введите url: ")
	myAccount, err := account.NewAccount(newLogin, newPassword, newUrl)
	if err != nil {
		panic(err)
	}
	fmt.Print(myAccount)

}

func getData(message string) string {
	fmt.Print(message)
	var data string
	fmt.Scanln(&data)
	return data
}
