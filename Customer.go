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
	return c
}
