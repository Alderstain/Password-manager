package account

import (
	"encoding/json"
	"fmt"
	"passwordManager/files"
)

type AccountsDB struct {
	Accounts []Account `json:"accounts"`
}

func NewAccountsDB() *AccountsDB {
	file, err := files.ReadFromFile("data.json")
	if err != nil {
		return &AccountsDB{
			Accounts: []Account{},
		}
	}
	var accDB AccountsDB
	err = json.Unmarshal(file, &accDB)
	if err != nil {
		fmt.Print(err)
	}
	return &accDB
}

func (accDB *AccountsDB) AddAccount(acc Account) {
	accDB.Accounts = append(accDB.Accounts, acc)
	file, err := accDB.ToBytes()
	if err != nil {
		panic(err)
	}
	files.WriteToFile(file, "data.json")
}

func (accDB *AccountsDB) ToBytes() ([]byte, error) {
	file, err := json.Marshal(accDB)
	if err != nil {
		return nil, err
	}
	return file, nil
}
