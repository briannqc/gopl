package bank_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	bank "github.com/briannqc/gopl/ch9/sec9.1/ex9.1"
)

func TestDeposit(t *testing.T) {
	assert.Equal(t, 0, bank.Balance(), "Initial balance should be 0")

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bank.Deposit(10)
		}()
	}

	wg.Wait()
	assert.Equal(t, 100, bank.Balance(), "Balance after deposits should be 100")

	assert.True(t, bank.Withdraw(10), "Withdrawing less than balance should succeed")
	assert.Equal(t, 90, bank.Balance(), "Balance after withdraw should be 90")

	assert.False(t, bank.Withdraw(100), "Withdrawing more than balance should fail")
	assert.Equal(t, 90, bank.Balance(), "Balance after withdraw fail should be unchanged")
}
