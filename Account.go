//Package bank contians structs and methods for operation of a bank
package bank

//Account stuff
type Account struct {
	Transactions                            Transaction
	CURRENT, SAVINGS, BUSINESS, AccountType int
}

//NewAccount Creates and populates an account
func NewAccount(AccType int) *Account {
	a := new(Account)
	a.AccountType = AccType
	return a
}
