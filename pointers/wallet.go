package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a bitcoin value
type Bitcoin int

// ErrInsufficientFunds will be thrown when withdraw amount exceeds balance
var ErrInsufficientFunds = errors.New("insufficient funds")

// Wallet stores a balance and allows deposits, withdrawls, and similaer wallet functionality
type Wallet struct {
	balance Bitcoin
}

// Deposit updates the balance with a deposit amount
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the current balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw reduces the balance by the amount
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
