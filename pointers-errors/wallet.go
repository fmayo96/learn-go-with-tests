package pointerserrors

import (
	"errors"
	"fmt"
)

var withdrawError = errors.New("withdraw amount exceeds balance")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin { // By convention all method receivers should be
	return w.balance // the same, pointers or not.
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return withdrawError
	}
	w.balance -= amount
	return nil
}
