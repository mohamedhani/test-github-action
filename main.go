// main.go
package main

import "fmt"

func Withdraw(currentBalance, amount int) int {
    if amount > currentBalance {
        
fmt.Println("Insufficient funds.")
               return currentBalance
    }

    newBalance := currentBalance - amount
    fmt.Printf("Withdrawal of %d successful. New balance: %d\n", amount, newBalance)
    return newBalance
}

func main(){

    currentBalance := 1000
    withdrawalAmount := 200

    newBalance := Withdraw(currentBalance, withdrawalAmount)
    fmt.Printf("Remaining balance: %d\n", newBalance)
}