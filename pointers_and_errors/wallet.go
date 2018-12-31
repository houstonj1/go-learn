package errors

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit method
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance method
func (w *Wallet) Balance() int {
	return w.balance
}
