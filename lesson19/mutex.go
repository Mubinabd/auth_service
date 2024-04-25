package main

import (
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common/math"
)

type initial_amount struct {
	amount int
	sync.Mutex
}

func (ia *initial_amount) withdraw_amount(withdraw_val int) {
	ia.Lock()
	ia.amount = ia.amount - withdraw_val
	ia.Unlock()
}
func (ia *initial_amount) deposit_amount(deposit_val int) {
	ia.Lock()
	ia.amount = ia.amount + deposit_val
	ia.Unlock()
}
// func (ia *initial_amount) get_balance() int {
// 	ia.Lock()
// 	val := ia.amount
// 	defer ia.Unlock()
// 	return val
// }
func main() {
	ia := initial_amount{amount: 2000}
	for itr := 1; itr <= 10; itr++ {
		go ia.deposit_amount(500)
		go ia.withdraw_amount(300)
	}
	final_amount := 0
	fmt.Println("Final balance is:", final_amount, ia.amount)
	
}
