//Package bank contians structs and methods for operation of a bank
package bank

//Account stuff
type Account struct {
	Transactions              Transaction
	Cost, AccountNo, Interest int64
}

//NewAccount Creates and populates an account
func NewAccount(AccNoIn int64, AccType string) *Account {
	a := new(Account)
	a.AccountNo = AccNoIn
	switch AccType {
	case "Basic":
		a.Interest = 1
		a.Cost = 0
	}
	return a
}
