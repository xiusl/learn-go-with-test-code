package _6_pointer

import "fmt"

type Bitcoin int

func (b Bitcoin)String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

/*Version 1
func (w Wallet) Deposit(amount int) {
	w.balance = amount
}

func (w Wallet) Balance() int {
	return w.balance
}
*/

/*NOTE
type Bitcoin int ==> go 语言支持根据现有类型创建新类型（取别名），可以使类型根据有描述性
*/