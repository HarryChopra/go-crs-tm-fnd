package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Balance int
	Mutex   *sync.Mutex
}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += amount
	a.Mutex.Unlock()
	wg.Done()
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= amount
	a.Mutex.Unlock()
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	a := Account{
		Balance: 100,
		Mutex:   &sync.Mutex{},
	}

	wg.Add(2)
	go func() { a.Deposit(50, wg) }()
	go func() { a.Withdraw(20, wg) }()
	wg.Wait()
	fmt.Println("A's Account Balance:", a.Balance)
}
