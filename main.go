package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("BLANK_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
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
	newPassword := getData("Введите пароль: ")
	newUrl := getData("Введите url: ")
	myAccount, err := newAccount(newLogin, newPassword, newUrl)
	if err != nil {
		panic(err)
	}
	printMyData(myAccount)

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
