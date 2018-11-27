package Bank

import "math/big"

type Bank struct {
	BankName string
	Capital, RentalIncome, LastAccNo int64
	Deposits map[string][]Account
}

func NewAccountArray(Deposit int64, Cost int64,AccNo int64, Interest big.Rat) []Account {
	acc := Account{Deposit, Cost, AccNo, Interest}
	var custAccounts []Account
	custAccounts = append(custAccounts, acc)
	cust := custAccounts
	return cust
}

func NewAccount(Deposit int64, Cost int64,AccNo int64, Interest big.Rat) Account {
	acc := Account{Deposit, Cost, AccNo, Interest}
	return acc
}