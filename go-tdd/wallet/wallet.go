// author: ashing
// time: 2020/4/16 11:09 下午
// mail: axingfly@gmail.com
// Less is more.

package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")


func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// btc wallet
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	//w.balance -= amount	 // 不可以直接这么写
	if w.balance-amount >= 0 {
		w.balance -= amount
	} else {
		return InsufficientFundsError
	}
	return nil
}
