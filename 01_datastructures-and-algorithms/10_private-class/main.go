package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id string
	accountType string
}

type Account struct {
	details *AccountDetails
	CustomerName string
}

func (account *Account) setDetails(id, accountType string) {
	account.details = &AccountDetails{
		id:          id,
		accountType: accountType,
	}
}

func (account *Account) getId() string {
	return account.details.id
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}

func main() {
	var account *Account = &Account{
		CustomerName: "John Smith",
	}
	account.setDetails("123", "current")
	jsonAccount, _ := json.MarshalIndent(account, "", "  ")
	fmt.Printf("Private class\n%s", string(jsonAccount))
	fmt.Println("account id:", account.getId())
	fmt.Println("account type:", account.getAccountType())
}
