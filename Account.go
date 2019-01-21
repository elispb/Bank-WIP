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
	a.CURRENT = 0
	a.SAVINGS = 1
	a.BUSINESS = 2
	return a
}

//Deposit adds money to an account
func (a *Account) Deposit(amount int) (bool, error) {
	if amount > 0 {
		t := NewTransaction("Deposit", amount)
		a.Transactions = append(a.Transactions, *t)
		return true, nil
	}
	return false, errors.New("Deposit must be positive")

}

//Withdrawal adds money to an account
func (a *Account) Withdrawal(amount int) (bool, error) {
	if amount > 0 {
		t := NewTransaction("Withdrawal", -amount)
		a.Transactions = append(a.Transactions, *t)
		return true, nil
	}
	return false, errors.New("Withdrawal must be positive")
}

//CalculateValue determins the value of the account
func (a *Account) CalculateValue() int {
	total := 0
	for _, t := range a.Transactions {
		total += t.amount
	}
	return total
}
