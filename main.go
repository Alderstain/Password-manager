package main

import (
	"fmt"
	"math/rand/v2"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) generatePassword(n int) {
	password := make([]rune, n)
	for i := 0; i < n; i++ {
		password[i] = rand.Int32N(73) + 50 // от 50 символа до 123 включительно
	}
	acc.password = string(password)
}

func main() {
	newLogin := getData("Введите логин: ")
	newUrl := getData("Введите url: ")
	myAccount := account{
		login: newLogin,
		url:   newUrl,
	}
	myAccount.generatePassword(12)
	printMyData(&myAccount)

}

func printMyData(acc *account) {
	fmt.Print(acc.login, " ", acc.password, " ", acc.url)
}

func getData(message string) string {
	fmt.Print(message)
	var data string
	fmt.Scanln(&data)
	return data
}
