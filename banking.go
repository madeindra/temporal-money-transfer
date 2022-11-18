package app

import (
	"errors"
	"math/rand"
)

// account includes AccountNumber and Balance
type account struct {
	AccountNumber string
	Balance       int64
}

// bank includes list of Account
type bank struct {
	Accounts []account
}

// finding account in database
func (b bank) findAccount(accountNumber string) (account, error) {

	for _, v := range b.Accounts {
		if v.AccountNumber == accountNumber {
			return v, nil
		}
	}

	return account{}, errors.New("account not found")
}

// InsufficientFundsError is an error type when the account does not have sufficent balance
type InsufficientFundsError struct{}

func (m *InsufficientFundsError) Error() string {
	return "Insufficient Funds"
}

// InvalidAccountError is an error type when the account number is not valid
type InvalidAccountError struct{}

func (m *InvalidAccountError) Error() string {
	return "Account number supplied is invalid"
}

// create mock bank
var mockBank = &bank{
	Accounts: []account{
		{AccountNumber: "85-150", Balance: 2000},
		{AccountNumber: "43-812", Balance: 0},
	},
}

// BankingService can be created using a hostmane
type BankingService struct {
	Hostname string
}

// withdraw money from the account
func (client BankingService) Withdraw(accountNumber string, amount int, referenceID string) (string, error) {
	acct, err := mockBank.findAccount(accountNumber)

	if err != nil {
		return "", &InvalidAccountError{}
	}

	if amount > int(acct.Balance) {
		return "", &InsufficientFundsError{}
	}

	return generateTransactionID("W", 10), nil
}

// deposit money to the account
func (client BankingService) Deposit(accountNumber string, amount int, referenceID string) (string, error) {

	_, err := mockBank.findAccount(accountNumber)
	if err != nil {
		return "", &InvalidAccountError{}
	}

	return generateTransactionID("D", 10), nil
}

// deposit money but return error
func (client BankingService) DepositThatFails(accountNumber string, amount int, referenceID string) (string, error) {
	return "", errors.New("this deposit has failed")
}

// transaction id generator
func generateTransactionID(prefix string, length int) string {
	randChars := make([]byte, length)
	for i := range randChars {
		allowedChars := "0123456789"
		randChars[i] = allowedChars[rand.Intn(len(allowedChars))]
	}
	return prefix + string(randChars)
}
