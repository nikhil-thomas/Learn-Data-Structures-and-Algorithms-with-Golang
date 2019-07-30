package main

import (
	"fmt"
)

type Account struct {
	id string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by id")
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("delete account by id")
}

type Customer struct {
	name string
	id int
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

type Transaction struct {
	id string
	amount float32
	srcAccountID string
	destAccountID string
}

func (tr *Transaction) create(src, dest string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	tr.srcAccountID = src
	tr.destAccountID = dest
	tr.amount = amount
	return tr
}

type BranchManagerFacade struct {
	account *Account
	customer *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{
		account: &Account{},
		customer: &Customer{},
		transaction: &Transaction{},
	}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	customer, account = facade.createCustomerAccount("John smith", "savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("1234", "5678", 1000000.0)
	fmt.Println(transaction.amount)
}
