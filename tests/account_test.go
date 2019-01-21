package tests

import (
	"testing"

	"../bank"
)

//TestDeposit tests the deposit function
func TestDeposit(t *testing.T) {
	acc := *bank.NewAccount(0)
	ran, _ := acc.Deposit(100)
	if !ran {
		t.Error("Method didn't return success")
	}
	if acc.CalculateValue() != 100 {
		t.Errorf("Acc value is %d, expected %d", acc.CalculateValue(), 100)
	}
}

//TestWithdraw Test withdrawal with valid value
func TestWithdraw(t *testing.T) {
	acc := *bank.NewAccount(0)
	ran, _ := acc.Withdrawal(100)
	if !ran {
		t.Error("Method didn't return success")
	}
	if acc.CalculateValue() != -100 {
		t.Errorf("Acc value is %d, expected %d", acc.CalculateValue(), -100)
	}
}
