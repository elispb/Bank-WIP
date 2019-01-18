//Package bank contians structs and methods for operation of a bank
package bank

//Bank contains the bank details and a list of customers
type Bank struct {
	BankName  string
	LastAccNo int64
	Customers map[string]Customer
}

//NewBank creates a new empty Bank
func NewBank(nameIn string) *Bank {
	b := new(Bank)
	b.BankName = nameIn
	return b
}

//NewBankFull creates a new completed Bank
func NewBankFull(nameIn string, lastAccNoIn int) *Bank {
	b := new(Bank)
	b.BankName = nameIn
	return b
}
