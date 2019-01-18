//Package bank contians structs and methods for operation of a bank
package bank

import (
	"strconv"
	"strings"
)

//Customer contains a name and a list of active accounts
type Customer struct {
	Name     string
	Accounts []Account
}

//NewCustomer creates and returns a new customer
func NewCustomer(nameIn string) *Customer {
	c := new(Customer)
	c.Name = nameIn
	c.Accounts = []Account{}
	return c
}

//OpenAccount opens a new account
func (c *Customer) OpenAccount(account Account) bool {
	start := len(c.Accounts)
	c.Accounts = append(c.Accounts, account)
	if start < len(c.Accounts) {
		return true
	}
	return false
}

//Statement returns the value of each account and a list of the last 5 transactions
func (c *Customer) Statement() string {
	var statement strings.Builder
	for _, acc := range c.Accounts {
		//Convert acc type to string
		switch acc.AccountType {
		case acc.CURRENT:
			statement.WriteString("Current Account\n")
		case acc.SAVINGS:
			statement.WriteString("Savings Account\n")
		case acc.BUSINESS:
			statement.WriteString("Business Account\n")
		}
		//Write last 5 Transactions
		for i := 0; i < len(acc.Transactions); i++ {
			if i < 5 {
				a := strconv.Itoa(acc.Transactions[i].amount)
				statement.WriteString(acc.Transactions[i].note + " " + a + "\n")
			}
		}
		//Write value
		statement.WriteString("Balance: " + strconv.Itoa(acc.CalculateValue()) + "\n")
	}
	return statement.String()
}
