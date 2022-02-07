package accounts

import (
	"errors"
	"fmt"
)

// 지역변수
var errNoMoney = errors.New("u dont have enough money")

// Account struct
type Account struct {
	owner   string
	balance int
}

// * 와 &의 사용
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Struct method
func (a *Account) Deposit(balance int) {
	a.balance += balance
}

// *가 없으면 복사, 있으면 기존 메모리
func (a Account) ShowDeposit() int {
	return a.balance
}

// error exceprion
// 여기서는 지역번수로 선언한 errNoMoney 를 return 하거나
// 정상인 경우 return 값을 줘야 하기 때문에 nil 사용
func (a *Account) WithDraw(balance int) error {
	if a.balance < balance {
		return errNoMoney
	}
	a.balance -= balance
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), " has ", a.ShowDeposit(), " won")
}
