// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan WithdrawRequest)

type WithdrawRequest struct {
	amount int
	result chan bool
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	result := make(chan bool)
	withdraws <- WithdrawRequest{
		amount: amount,
		result: result,
	}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			continue
		case req := <-withdraws:
			if req.amount > balance {
				req.result <- false
			} else {
				balance -= req.amount
				req.result <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
