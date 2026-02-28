package main

import "fmt"

type Account struct {
	ID      string
	Owner   string
	Balance float64
}

// TODO: 实现以下方法

// Deposit 存款（指针接收者）
func (a *Account) Deposit(amount float64) {
	a.Balance = a.Balance + amount

}

// Withdraw 取款（指针接收者），返回成功或失败
func (a *Account) Withdraw(amount float64) {
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

// Transfer 转账：从当前账户转账到目标账户
func (a *Account) Transfer(target *Account, amount float64) bool {
	if amount <= 0 {
		return false
	}

	if a.Balance < amount {
		return false
	}

	if a.Balance > amount {
		a.Balance = a.Balance - amount
		target.Balance = target.Balance + amount
		return true
	}

	return true

}

// GetInfo 获取账户信息
func (a Account) GetInfo() string {
	return fmt.Sprintf("id:%s,Owner:%s,Balance:%.2f", a.ID, a.Owner, a.Balance)
}

func main() {
	acc1 := &Account{ID: "001", Owner: "Alice", Balance: 1000}
	acc2 := &Account{ID: "002", Owner: "Bob", Balance: 500}

	fmt.Println(acc1.GetInfo())

	acc1.Deposit(200)
	fmt.Println("存款后:", acc1.Balance) // 1200

	acc1.Withdraw(300)
	fmt.Println(acc1.Balance)

	acc1.Transfer(acc2, 200)
	fmt.Println(acc1.GetInfo())
	fmt.Println(acc2.GetInfo())

}
