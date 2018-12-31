package errors

import "fmt"

// Dollar type
type Dollar int

// Wallet struct
type Wallet struct {
	balance Dollar
}

// Deposit method
func (w *Wallet) Deposit(amount Dollar) {
	w.balance += amount
}

// Balance method
func (w *Wallet) Balance() Dollar {
	return w.balance
}

// Withdraw method
func (w *Wallet) Withdraw(amount Dollar) {
	w.balance -= amount
}

// Stringer interface
type Stringer interface {
	String() string
}

func (d Dollar) String() string {
	return fmt.Sprintf("$%d", d)
}
