package accounts

import "errors"

var errNoMoney = errors.New("u dont have enough money")

// Account struct
type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(balance int) {
	a.balance += balance
}

func (a Account) ShowDeposit() int {
	return a.balance
}

func (a *Account) WithDraw(balance int) error {
	if a.balance < balance {
		return errNoMoney
	}
	a.balance -= balance
	return nil
}
