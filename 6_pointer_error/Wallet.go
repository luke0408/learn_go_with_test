package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	// (*w).balance라고 쓰지 않아도 자동 역참조가 가능함
	return w.balance
}
