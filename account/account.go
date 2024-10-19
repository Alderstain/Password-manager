package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func NewAccount(login, password, urlString string) (*account, error) {
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
