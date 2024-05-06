// main_test.go
package main

import "testing"

func TestWithdraw(t *testing.T) {
    // Test case 1: Sufficient funds
    initialBalance := 1000
    withdrawalAmount := 200
    expectedBalance := 800

    result := Withdraw(initialBalance, withdrawalAmount)
    if result != expectedBalance {
        t.Errorf("Withdraw(%d, %d) expected %d, but got %d", initialBalance, withdrawalAmount, expectedBalance, result)
    }

    // Test case 2: Insufficient funds
    initialBalance = 100
    withdrawalAmount = 200
    expectedBalance = 100

    result = Withdraw(initialBalance, withdrawalAmount)
    if result != expectedBalance {
        t.Errorf("Withdraw(%d, %d) expected %d, but got %d", initialBalance, withdrawalAmount, expectedBalance, result)
    }
}