package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
)

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("BLANK_LOGIN")
	}
	urlString = "https://" + urlString
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		Login:    login,
		Password: password,
		Url:      urlString,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)
	for i := 0; i < n; i++ {
		password[i] = rand.Int32N(73) + 50 // from 50 to 123 symbols of utf-8
	}
	acc.Password = string(password)
}
