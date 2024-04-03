package wallet

import (
    "errors"
    "fmt"
    "sync"
)

type Bitcoin float64

type Wallet struct {
    balance Bitcoin
    mutex   sync.Mutex
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
    w.mutex.Lock()
    defer w.mutex.Unlock()

    w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
    w.mutex.Lock()
    defer w.mutex.Unlock()

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%.2f BTC", b)
}
