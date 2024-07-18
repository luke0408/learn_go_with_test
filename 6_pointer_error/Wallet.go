package main

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// (*w).balance라고 쓰지 않아도 자동 역참조가 가능함
	return w.balance
}
