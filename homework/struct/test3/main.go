package main

import "fmt"

func main() {

	acc := Account{
		Owner:   "heidi",
		Balance: 20000,
	}
	fmt.Println(acc.CheckBalance())
	acc.Deposit(1000)
	fmt.Println(acc.CheckBalance())
	acc.WithDraw(2000)
	fmt.Println(acc.CheckBalance())

}

type Account struct {
	Owner   string
	Balance float64
}

// 返回余额 值接收器
func (a Account) CheckBalance() string {
	return fmt.Sprintf("余额:%.2f", a.Balance)
}

// 存款
func (a *Account) Deposit(amount float64) {
	if amount < 0 {
		fmt.Println("存款必须大于0")
		return
	}
	a.Balance = a.Balance + amount
}

// 取款
func (a *Account) WithDraw(amount float64) {
	if amount < 0 {
		fmt.Println("取款金额必须大于0")
		return
	}
	if amount > a.Balance {
		fmt.Println("余额不足，请冲新输入")
		return
	} else {
		a.Balance = a.Balance - amount
	}

}
