//Package bank contians structs and methods for operation of a bank
package bank

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
func OpenAccount(account Account, cust Customer) bool {
	start := len(cust.Accounts)
	cust.Accounts = append(cust.Accounts, account)
	if start < len(cust.Accounts) {
		return true
	}
	return false
}
