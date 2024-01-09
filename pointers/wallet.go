package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
  balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
  w.balance += amount  // (*w).balance += amount convention
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance // (*w).balance convention
}

var ErrorInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.balance {
    return ErrorInsufficientFunds 
  }

  w.balance -= amount
  return nil
}

type Stinger interface {
  String() string
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

