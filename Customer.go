//Package bank contians structs and methods for operation of a bank
package bank

//Customer contains a name and a list of active accounts
type Customer struct {
	Name     string
	Accounts map[int]Account
}
