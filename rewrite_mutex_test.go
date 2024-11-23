package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


type BankAccount struct {
	RWmutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWmutex.Lock()
	account.Balance += amount
	account.RWmutex.Unlock()
}

func (account *BankAccount) GetBalance() int  {
	account.RWmutex.RLock()
	balance := account.Balance
	account.RWmutex.RUnlock()
	return balance
}


func TestRewriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Balance int
	Name string
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int)  {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(- amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock() 
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Balance: 10000,
		Name: "Eko",
	}

	user2 := UserBalance{
		Balance: 10000,
		Name: "Budi",
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 1000)

	time.Sleep(2 * time.Second)
	fmt.Println("User 1 = ", user1.Balance)
	fmt.Println("User 2 = ", user2.Balance)
}