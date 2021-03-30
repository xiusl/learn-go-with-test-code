package _6_pointer

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

/*NOTE
	Version 1
	测试失败，wallet_test.go:14: got 0 want 10
	修改了 wallet 的 balance 的值，但是没有生效？
	这里是因为 func (w Wallet) Deposit(amount int) {} 方法传入的 Wallet 是值，
	会将测试 wallet 拷贝一份传入方法中，修改的是拷贝后的 balance，这个可以通过打印对象地址查看
*/