package Bank

import "math/big"

type Account struct {
	Value, Cost, AccountNo int64
	Interest big.Rat
}
