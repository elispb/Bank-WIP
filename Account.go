//Package bank contians structs and methods for operation of a bank
package bank

import "errors"

//Account stuff
type Account struct {
	Transactions                            []Transaction
	CURRENT, SAVINGS, BUSINESS, AccountType int
}

//NewAccount Creates and populates an account
func NewAccount(AccType int) *Account {
	a := new(Account)
	a.AccountType = AccType
	return a
}

//Deposit adds money to an account
func Deposit(account Account, amount int) (bool, error) {
	if amount > 0 {
		t := NewTransaction("Deposit", amount)
		account.Transactions = append(account.Transactions, *t)
		return true, nil
	} else {
		return false, errors.New("Deposit must be positive")
	}
}

//Withdrawal adds money to an account
func Withdrawal(account Account, amount int) (bool, error) {
	if amount > 0 {
		t := NewTransaction("Withdrawal", -amount)
		account.Transactions = append(account.Transactions, *t)
		return true, nil
	} else {
		return false, errors.New("Withdrawal must be positive")
	}
}
