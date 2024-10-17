package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {
	dandan := generatePassword(6)
	account := account{
		login:    "bakhmnmn",
		password: dandan,
		url:      "google.com",
	}
	printMyData(account)

}

func printMyData(acc account) {
	fmt.Print(acc.login, " ", acc.password, " ", acc.url)
}

func generatePassword(n int) string {
	password := make([]rune, n)
	for i := 0; i < n; i++ {
		password[i] = rand.Int32N(73) + 50
	}
	return string(password)
}

func testGeneratePassword(t *testing.T) {
	a := 6
	result := generatePassword(a)
	expected := len(result)

	if expected != 6 {
		t.Errorf("generatePassword(6): expected string of len = %d", a)
	}
}

// от 50 символа до 123 включительно
