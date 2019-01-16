//Package bank contians structs and methods for operation of a bank
package bank

import (
	"time"
)

//Transaction note of transaction, amount, and the time it occured at.
type Transaction struct {
	note      string
	timeStamp time.Time
	amount    int
}

//NewTransaction creates a new transaction
func NewTransaction(noteIn string, amountIn int) *Transaction {
	t := new(Transaction)
	t.timeStamp = time.Now()
	t.note = noteIn
	t.amount = amountIn
	return t
}
