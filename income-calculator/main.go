package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankBalance int

	fmt.Printf("Initial amount : $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	var wg sync.WaitGroup
	var mt sync.Mutex
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for range 52 {
				mt.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				mt.Unlock()
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final Amount: $%d.00", bankBalance)
	fmt.Println()
}
