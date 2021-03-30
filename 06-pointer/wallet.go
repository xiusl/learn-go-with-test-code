package _6_pointer

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance = amount
}

func (w *Wallet) Balance() int {
	return w.balance
}


/*Version 1
func (w Wallet) Deposit(amount int) {
	w.balance = amount
}

func (w Wallet) Balance() int {
	return w.balance
}
*/